package main

import (
	"context"
	"fmt"

	"github.com/pratikdhanavesearce/mongodb-adapter/mongodb"
	"github.com/pratikdhanavesearce/mongodb-adapter/util"
)

func FileCreationsPart(str []string) {
	//fmt.Println(str)
	ctx := context.TODO()
	mongo_client, err := mongodb.NewConnection(str[1])
	if err != nil {
		fmt.Println("Error: ", err)
	}
	//fmt.Println("Mongo Connection Done")
	db := mongo_client.Database(str[2])
	//fmt.Println(db.Name())
	collections, err := mongodb.ListCollection(db)
	if err != nil {
		fmt.Println("Error: ", err)
	}
	//fmt.Println("Got List of collections")
	if err = util.FileCreate(ctx, db, collections); err != nil {
		fmt.Println("Error: in file Creation: ", err)
	}
	mongo_client.Disconnect(context.TODO())
}
