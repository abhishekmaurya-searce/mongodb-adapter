package util

import (
	"context"
	"fmt"
	"os"
	"strings"

	"github.com/pratikdhanavesearce/mongodb-adapter/mongodb"
	"go.mongodb.org/mongo-driver/mongo"
	"gopkg.in/mgo.v2/bson"
)

func FileCreate(ctx context.Context, db *mongo.Database, col []string) error {
	code := `package view
import (
	"time"
)
`
	file, err := os.Create("./view/schema.go")
	if err != nil {
		fmt.Println("Failed to create file:", err)
		return err
	}
	sqlfile, err := os.Create("./migration/up.sql")
	if err != nil {
		fmt.Println("Failed to Create file:", err)
		return err
	}
	retrive, err := os.Create("./mongodb/retrive_mongodb.go")
	if err != nil {
		fmt.Println(err)
	}

	insertfile, err := os.Create("./view/insert_spanner.go")
	if err != nil {
		fmt.Println("Failed to Create file:", err)
		return err
	}
	file_part2, err := os.Create("./cmd/library/part2.go")
	if err != nil {
		fmt.Println("Error creating part2  file:", err)
	}
	var sql_code string
	insert_code := `package view

import (
"fmt"
	
"cloud.google.com/go/spanner"
)
	
type Collection interface {
InsertSpanner(table string, client *spanner.Client) (*spanner.Mutation, error)
}
`
	retrive_code := `
package mongodb
import (
	"context"
	"encoding/json"
	"reflect"
	"github.com/pratikdhanavesearce/mongodb-adapter/view"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"gopkg.in/mgo.v2/bson"
)
`
	part2_code := `package main

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
	for _, val := range col {
		var result bson.M
		_ = db.Collection(val).FindOne(context.TODO(), bson.M{}).Decode(&result)
		var tables []string
		tables = append(tables, val)
		table := strings.ToUpper(string(val[0])) + val[1:]
		part2_code += fmt.Sprintf(`
		case "%s":arr_%s, err := mongodb.Retrive_%s(db.Collection(table))
		fmt.Println("Inserting into ", table)
		if err != nil {
			fmt.Println("Error in retriving data", err)
		}
		var mut []*spanner.Mutation
		for _, value := range arr_%s {
			temp, err := view.Insert(&value, table, spanner_client)
			if err != nil {
				fmt.Println(err)
			}
			mut = append(mut, temp)
		}
		_, err = spanner_client.Apply(context.TODO(), mut)
		if err != nil {
			fmt.Println("Error: ", err)
		}`, val, table, table, table)
		code1, code2, code3, code4, code5 := mongodb.CollectionToStruct(result, result, tables, tables)
		part2_code += code5
		code += code1
		sql_code += code2
		retrive_code += code3
		insert_code += code4
	}
	part2_code += `
		}
	}
	spanner_client.Close()
	mongo_client.Disconnect(context.TODO())
	}
`
	insert_code += `
func Insert(table Collection, collection string, client *spanner.Client) (*spanner.Mutation, error) {
	return table.InsertSpanner(collection, client)
}`
	_, err = file.WriteString(code)
	if err != nil {
		return err
	}
	defer file.Close()
	_, err = retrive.WriteString(retrive_code)
	if err != nil {
		fmt.Println(err)
	}
	defer retrive.Close()
	_, err = sqlfile.WriteString(sql_code)
	if err != nil {
		return err
	}
	defer sqlfile.Close()
	_, err = insertfile.WriteString(insert_code)
	if err != nil {
		return err
	}
	defer sqlfile.Close()
	_, err = file_part2.WriteString(part2_code)
	if err != nil {
		fmt.Println("Error: Writing file:", err)
	}
	file_part2.Close()
	return nil
}
