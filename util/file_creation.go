package util

import (
	"context"
	"fmt"
	"os"

	"github.com/pratikdhanavesearce/mongodb-adapter/cloudspanner"
	"github.com/pratikdhanavesearce/mongodb-adapter/mongodb"
	"go.mongodb.org/mongo-driver/mongo"
)

func FileCreate(ctx context.Context, db *mongo.Database, col []string) error {
	code := `package view
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
	for _, val := range col {
		collection := db.Collection(val)
		code1, code2 := mongodb.CollectionToStruct(*collection, ctx, val)
		code += code1
		sql_code += code2

	}
	_, err = file.WriteString(code)
	if err != nil {
		return err
	}
	defer file.Close()
	retrive_code := mongodb.RetriveCollection(col)
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
