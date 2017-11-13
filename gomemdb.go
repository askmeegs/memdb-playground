package main

//
// import (
// 	"fmt"
//
// 	"github.com/sigma-dev/sigma/model"
// )
//
// var p *model.Plugin
//
// // skeleton how to create memdb for sigma
// func memdb() {
//
// 	// Create object schema for data model
// 	schema := createSchema()
//
// 	// Init memdb
// 	db, err := memdb.NewMemDB(schema)
// 	if err != nil {
// 		panic(err)
// 	}
//
// 	// EXAMPLE - INSERT
// 	txn := db.Txn(true)
// 	p = &model.Plugin{
// 		MetaImpl:    model.NewMeta("coldID", "coldName"),
// 		Summary:     "it's cold",
// 		Description: "more about how it's cold",
// 		Maintainer:  "meegan",
// 		Version:     "0.1",
// 		Vars: map[string]model.PluginVar{
// 			"temperature": model.PluginVar{
// 				Summary: "degrees",
// 				Type:    "int",
// 				Default: 32,
// 			},
// 		},
// 		Requirements: []model.PluginReq{
// 			model.PluginReq{
// 				Name:        "req1",
// 				Description: "req1 Description",
// 				Version:     ">=0.1",
// 			},
// 		},
// 		Container: "testimage:1.0",
// 		Types:     map[string]model.Type{},
// 	}
// 	if err := txn.Insert("plugin", p); err != nil {
// 		panic(err)
// 	}
// 	txn.Commit()
//
// 	// EXAMPLE  - GET
// 	txn = db.Txn(false)
// 	defer txn.Abort()
// 	raw, err := txn.First("plugin", "id", "coldID")
// 	if err != nil {
// 		panic(err)
// 	}
// 	fmt.Printf("%#v\n", raw.(*model.Plugin))
// }
//
// // Have to copy the whole sigma object definition into here #wow
// func createSchema() *memdb.DBSchema {
//
// 	// Extract object subfield (MetaImpl.MetaID) to use as the required id index
// 	idIndex := &memdb.StringFieldIndex{Field: "MetaImpl.MetaID"}
// 	idIndex.FromObject(p)
//
// 	return &memdb.DBSchema{
// 		Tables: map[string]*memdb.TableSchema{
// 			"plugin": &memdb.TableSchema{
// 				Name: "plugin",
// 				Indexes: map[string]*memdb.IndexSchema{
// 					"id": &memdb.IndexSchema{
// 						Name:    "id",
// 						Unique:  true,
// 						Indexer: idIndex,
// 					},
// 					"Container": &memdb.IndexSchema{
// 						Name:    "Container",
// 						Unique:  true,
// 						Indexer: &memdb.StringFieldIndex{Field: "Container"},
// 					},
// 				},
// 			},
// 		},
// 	}
// }
