package main // Wrapper on bolt

import (
	"log"
	"os"

	"github.com/asdine/storm"
	"github.com/asdine/storm/q"
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
	complexQuery()
	os.Remove("my.db")
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
	err = db.Select(q.Eq("Kind", "backpacking"), q.Eq("International", true)).Find(&result)
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
