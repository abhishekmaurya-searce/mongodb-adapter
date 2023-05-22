package cloudspanner

import (
	"context"
	"fmt"
	"os"

	"cloud.google.com/go/spanner"
	"github.com/pratikdhanavesearce/mongodb-adapter/model"
	"google.golang.org/api/iterator"
)

func NewConnection(uri string) (*spanner.Client, error) {

	if err := os.Setenv("SPANNER_EMULATOR_HOST", "localhost:9010"); err != nil {
		fmt.Println("Error while setting environment variable: ")
		return nil, err
	}
	client, err := spanner.NewClient(context.TODO(), uri)
	if err != nil {
		fmt.Println("Error: While creating client: ")
		return nil, err
	}
	return client, nil
}
func Insert(uri string, data []model.Person) error {
	client, err := NewConnection(uri)
	if err != nil {
		fmt.Println("Error in connecting to spanner: ")
	}
	for _, value := range data {
		_, err = client.Apply(context.TODO(), []*spanner.Mutation{
			spanner.Insert("person",
				[]string{"id", "firstname", "lastname", "age"},
				[]interface{}{int64(value.Id), value.First, value.Last, int64(value.Age)})})
		if err != nil {
			fmt.Println("Error: ", err)
		}
	}
	fmt.Println("Inserted", len(data), "Rows")
	return nil
}
func Read(uri string) error {
	client, err := NewConnection(uri)
	if err != nil {
		fmt.Println("Error: Cant'read:")
	}
	//fmt.Println(client.DatabaseName())
	dml := spanner.NewStatement("SELECT * FROM person;")
	iter := client.Single().Query(context.TODO(), dml)
	for {
		//fmt.Println(iter.Next())
		row, err := iter.Next()
		// var person model.Person
		// if err = row.ToStruct(&person); err != nil {
		// 	fmt.Println("Error in getting rows")
		// 	return err
		// }
		// fmt.Println(person)
		if row != nil {
			fmt.Println(row)
		}
		if err == iterator.Done {
			break
		}
		if err != nil {
			fmt.Println("Error in itterating over rows: ", err)
			return err
		}
		//var i int
	}
	iter.Stop()
	return nil
}
func CreateTable(tablename, uri string) error {
	client, err := NewConnection(uri)
	if err != nil {
		fmt.Println("Error :", err)
	}
	dml := spanner.NewStatement(`
		CREATE TABLE ` + tablename + ` (
			id INT64,
			firstname STRING(MAX),
			lastname STRING(MAX),
			age INT64,
			PRIMARY KEY (id)
		)	`)
	client.Single().Query(context.TODO(), dml)
	fmt.Println("Table Created")
	return nil
}
