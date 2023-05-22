package main

import (
	"context"
	"fmt"
	"os"

	//spanner "cloud.google.com/go/spanner/apiv1"

	"github.com/pratikdhanavesearce/mongodb-adapter/cloudspanner"
	"github.com/pratikdhanavesearce/mongodb-adapter/mongodb"
)

func main() {
	str := os.Args
	//fmt.Println(str)
	mongo_client, err := mongodb.NewConnection(str[1])
	if err != nil {
		fmt.Println("Error: ", err)
	}
	db := mongo_client.Database(str[2])
	collections, err := mongodb.ListCollection(db)
	if err != nil {
		fmt.Println("Error: ", err)
	}
	for _, coll := range collections {
		fmt.Println(coll)
	}
	err = mongodb.Head(mongo_client.Database(str[2]).Collection(collections[0]))
	if err != nil {
		fmt.Println("Error: ", err)
	}
	people, err := mongodb.Retrive(mongo_client.Database(str[2]).Collection(collections[0]))
	if err != nil {
		fmt.Println("Error: ", err)
	}
	for _, value := range people {
		fmt.Println(value)
	}
	//fmt.Println(mongo_client)
	//x := mongo_client.Database(str[2]).Collection(collections[0])
	mongo_client.Disconnect(context.TODO())
	//ctx := context.TODO()
	connection_string := str[3]
	// err = cloudspanner.CreateTable("person", connection_string)
	// if err != nil {
	// 	fmt.Println("Error :", err)
	// }
	// if err = cloudspanner.Insert(connection_string, people); err != nil {
	// 	fmt.Println("Insertion Error:", err)
	// }
	err = cloudspanner.Read(connection_string)
	if err != nil {
		fmt.Println("Error :", err)
	}
}
