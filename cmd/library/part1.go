package main

import (
	"context"
	"fmt"
	"os"
	"strings"

	"github.com/pratikdhanavesearce/mongodb-adapter/mongodb"
	"github.com/pratikdhanavesearce/mongodb-adapter/util"
)

func part2(coll []string) string {
	code := `package main

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
			`
	for _, value := range coll {
		table := strings.ToUpper(string(value[0])) + value[1:]
		code += fmt.Sprintf(`
			case "%s":arr, err := mongodb.Retrive_%s(db.Collection(table))
			fmt.Println("Inserting into ", table)
			if err != nil {
				fmt.Println("Error in retriving data", err)
			}
			var mut []*spanner.Mutation
			for _, value := range arr {
				temp, err := view.Insert(&value, table, spanner_client)
				if err != nil {
					fmt.Println(err)
				}
				mut = append(mut, temp)
			}
			_, err = spanner_client.Apply(context.TODO(), mut)
			if err != nil {
				fmt.Println("Error: ", err)
			}`, value, table)
	}
	code += `
	}
}
spanner_client.Close()
mongo_client.Disconnect(context.TODO())
}

`
	return code
}
func FileCreationsPart(str []string) {
	//fmt.Println(str)
	ctx := context.TODO()
	mongo_client, err := mongodb.NewConnection(str[1])
	if err != nil {
		fmt.Println("Error: ", err)
	}
	db := mongo_client.Database(str[2])
	collections, err := mongodb.ListCollection(db)
	if err != nil {
		fmt.Println("Error: ", err)
	}
	if err = util.FileCreate(ctx, db, collections); err != nil {
		fmt.Println(err)
	}
	file_part2, err := os.Create("./cmd/library/part2.go")
	part2_code := part2(collections)
	if err != nil {
		fmt.Println("Error creating part2  file:", err)
	}
	_, err = file_part2.WriteString(part2_code)
	if err != nil {
		fmt.Println(err)
	}
	file_part2.Close()
	mongo_client.Disconnect(context.TODO())
}
