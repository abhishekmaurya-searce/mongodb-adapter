package util

import (
	"context"
	"fmt"
	"os"

	"github.com/pratikdhanavesearce/mongodb-adapter/cloudspanner"
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
	retrive, err := os.Create("./mongodb/retrive.go")
	if err != nil {
		fmt.Println(err)
	}

	insertfile, err := os.Create("./view/insert.go")
	if err != nil {
		fmt.Println("Failed to Create file:", err)
		return err
	}
	var sql_code string
	insert_code := cloudspanner.InsertScripts(col)
	retrive_code := `
package mongodb
`
	for _, val := range col {
		var result bson.M
		_ = db.Collection(val).FindOne(context.TODO(), bson.M{}).Decode(&result)
		var tables []string
		tables = append(tables, val)
		code1, code2, code3 := mongodb.CollectionToStruct(result, tables)
		code += code1
		sql_code += code2
		retrive_code += code3
	}
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
	return nil
}
