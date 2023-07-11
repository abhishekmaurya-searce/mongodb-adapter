package main

import (
	"context"
	"fmt"

	"cloud.google.com/go/spanner"
	"github.com/pratikdhanavesearce/mongodb-adapter/cloudspanner"
	"github.com/pratikdhanavesearce/mongodb-adapter/mongodb"
	"github.com/pratikdhanavesearce/mongodb-adapter/view"
)

func RetriveAndInsertPart(str []string) {
	mongo_client, err := mongodb.NewConnection(str[1])
	if err != nil {
		fmt.Println("Error: ", err)
	}
	db := mongo_client.Database(str[2])
	spanner_client, err := cloudspanner.NewConnection(str[3])
	if err != nil {
		fmt.Println("Error Spanner client", err)
	}
	collections, err := mongodb.ListCollection(db)
	if err != nil {
		fmt.Println("Error: ", err)
	}
	for _, table := range collections {
		switch table {

		case "shipwrecks":arr_Shipwrecks, err := mongodb.Retrive_Shipwrecks(db.Collection(table))
		fmt.Println("Inserting into ", table)
		if err != nil {
			fmt.Println("Error in retriving data", err)
		}
		var mut []*spanner.Mutation
		for _, value := range arr_Shipwrecks {
			temp, err := view.Insert(&value, table, spanner_client)
			if err != nil {
				fmt.Println(err)
			}
			mut = append(mut, temp)
		}
		_, err = spanner_client.Apply(context.TODO(), mut)
		if err != nil {
			fmt.Println("Error: ", err)
		}
		}
	}
	spanner_client.Close()
	mongo_client.Disconnect(context.TODO())
	}
