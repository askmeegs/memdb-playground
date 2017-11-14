package main // Wrapper on bolt

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/asdine/storm"
	"github.com/asdine/storm/q"
	"github.com/sigma-dev/sigma/model"
)

type Vacation struct {
	ID            int
	Name          string
	Cost          float64
	Kind          string
	International bool
	PackingList   []string
}

func main() {
	sigmaQuery()
	os.Remove("my.db")
}

func sigmaQuery() {
	db, err := storm.Open("my.db")
	if err != nil {
		log.Fatalf("Could not open: %v", err)
	}
	defer db.Close()

	name := "stormy"
	version := "0.1"

	p := model.Plugin{
		MetaImpl:    metaNameID(name),
		Summary:     "summary for " + name,
		Description: "description for" + name,
		Maintainer:  "sigmadev",
		Version:     version,
		Vars: map[string]model.PluginVar{
			"var1": model.PluginVar{
				Type:        "string",
				Default:     "amazing",
				Description: "placeholder var1 description",
			},
		},
		Requirements: []model.PluginReq{
			model.PluginReq{
				Name:        "my-req",
				Description: "req1 Description",
				Version:     "0.1",
				Min:         1,
			},
		},
		Container: fmt.Sprintf("%s-container:%s", name, version),
		Types: map[string]model.Type{
			name + "-custom-type": model.Type{
				Name:    name + "-custom-type",
				Base:    "string",
				Default: "magic",
			},
		},
	}

	id := genID(p.GetName())
	p.SetID(id)
	err = db.Save(&p)
	if err != nil {
		log.Fatal(err)
	}

	var result []model.Plugin

	matchers := []q.Matcher{q.Eq("Version", "0.1")}
	prettymatch, _ := json.MarshalIndent(matchers, "", "  ")
	fmt.Printf("got matchers: \n %s\n", string(prettymatch))

	err = db.Select(matchers...).Find(&result)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(result)

}

func complexQuery() {
	db, err := storm.Open("my.db")
	if err != nil {
		log.Fatalf("Could not open: %v", err)
	}
	defer db.Close()
	v := Vacation{
		ID:            123,
		Name:          "EuropeAdventure",
		Cost:          314.15,
		Kind:          "backpacking",
		International: true,
		PackingList:   []string{"backpack", "boots", "compass", "map", "train pass"},
	}

	// INSERT
	err = db.Save(&v)
	if err != nil {
		log.Fatalf("Could not save EuropeAdventure: %v", err)
	}

	v = Vacation{
		ID:            124,
		Name:          "USAAdventure",
		Cost:          314.15,
		Kind:          "backpacking",
		International: false,
		PackingList:   []string{"backpack", "boots", "compass", "map", "train pass"},
	}

	// INSERT
	err = db.Save(&v)
	if err != nil {
		log.Fatalf("Could not save EuropeAdventure: %v", err)
	}

	var result []Vacation

	matchers := []q.Matcher{q.Eq("Kind", "backpacking"), q.Eq("International", true)}
	prettymatch, _ := json.MarshalIndent(matchers, "", "  ")
	fmt.Printf("got matchers: \n %s\n", string(prettymatch))

	err = db.Select(matchers...).Find(&result)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(result)
}

func easyStuff() {
	db, err := storm.Open("my.db")
	if err != nil {
		log.Fatalf("Could not open: %v", err)
	}
	defer db.Close()
	v := Vacation{
		ID:            123,
		Name:          "EuropeAdventure",
		Cost:          314.15,
		Kind:          "backpacking",
		International: true,
		PackingList:   []string{"backpack", "boots", "compass", "map", "train pass"},
	}

	// INSERT
	err = db.Save(&v)
	if err != nil {
		log.Fatalf("Could not save EuropeAdventure: %v", err)
	}

	v.ID++
	v.Name = "AsiaAdventure"
	v.Cost = 1342.50
	err = db.Save(&v)
	if err != nil {
		log.Fatalf("Could not save AsiaAdventure: %v", err)
	}

	// GET
	var gotV Vacation
	err = db.One("Name", "EuropeAdventure", &gotV)

	if err != nil {
		log.Fatalf("Could not get one: %v", err)
	}
	log.Printf("AFTER GET ONE: \n %#v \n", gotV)

	// GET MANY
	var vacs []Vacation
	err = db.Find("Kind", "backpacking", &vacs)
	if err != nil {
		log.Fatalf("Could not get many: %v", err)
	}
	log.Printf("AFTER GET MANY: \n %#v \n", vacs)

}

// ported helpers
func genID(name string) model.MetaID {
	return metaNameID(name).GetID()
}
func metaNameID(name string) model.MetaImpl {
	return model.NewMeta(nameID(name), name)
}
func nameID(name string) model.MetaID {
	return model.MetaID(fmt.Sprintf("%s-id", name))
}
func metaFromID(id string, t model.MetaType) model.MetaImpl {
	met := model.NewMeta(model.MetaID(id), strings.TrimSuffix(id, "-id"))
	subMeta := make(map[string]interface{})
	subMeta["Type"] = t
	met.Metadata = subMeta
	return met
}

func nameRef(name string) model.ModelRef {
	return model.ModelRef(nameID(name).String())
}
