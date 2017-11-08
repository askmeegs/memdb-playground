package main

import (
	"fmt"

	"github.com/hashicorp/go-memdb"
	"github.com/sigma-dev/sigma/model"
)

// skeleton how to create memdb for sigma
func main() {

	// Create object schema for data model
	schema := createSchema()

	// Init memdb
	db, err := memdb.NewMemDB(schema)
	if err != nil {
		panic(err)
	}

	// EXAMPLE - INSERT
	txn := db.Txn(true)
	p := &model.Plugin{
		MetaImpl:    model.NewMeta("pluginID", "pluginName"),
		Summary:     "pluginSummary",
		Description: "pluginDescription",
		Maintainer:  "pluginMaintainer",
		Version:     "pluginVersion",
		Vars: map[string]model.PluginVar{
			"var1": model.PluginVar{
				Description: "varDesc",
				Summary:     "varSummary",
				Type:        "string",
				Default:     "helloworld",
			},
		},
		Requirements: []model.PluginReq{
			model.PluginReq{
				Name:        "req1",
				Description: "req1 Description",
				Version:     ">=0.1",
			},
		},
		Container: "testimage:1.0",
		Types:     map[string]model.Type{},
	}
	if err := txn.Insert("plugin", p); err != nil {
		panic(err)
	}
	txn.Commit()

	// EXAMPLE  - GET
	txn = db.Txn(false)
	defer txn.Abort()
	raw, err := txn.First("plugin", "id", "testimage:1.0")
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s\n", raw.(*model.Plugin).Container)
}

// Pain. Have to copy the whole sigma object definition into here
func createSchema() *memdb.DBSchema {
	return &memdb.DBSchema{
		Tables: map[string]*memdb.TableSchema{
			"plugin": &memdb.TableSchema{
				Name: "plugin",
				Indexes: map[string]*memdb.IndexSchema{
					"id": &memdb.IndexSchema{
						Name:    "id",
						Unique:  true,
						Indexer: &memdb.StringFieldIndex{Field: "Container"},
					},
				},
			},
		},
	}
}
