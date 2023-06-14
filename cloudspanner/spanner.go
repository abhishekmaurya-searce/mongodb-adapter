package cloudspanner

import (
	"context"
	"fmt"
	"os"
	"reflect"
	"strings"

	"cloud.google.com/go/spanner"
	//"github.com/pratikdhanavesearce/mongodb-adapter/view"
	"go.mongodb.org/mongo-driver/bson"
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

//	func Insert(table string, data []view.Author, client *spanner.Client, ctx context.Context) error {
//		for _, value := range data {
//			mut, err := spanner.InsertStruct(table, value)
//			if err != nil {
//				fmt.Println("Error in inserting Struct: ")
//				return err
//			}
//			_, err = client.Apply(context.TODO(), []*spanner.Mutation{mut})
//			if err != nil {
//				fmt.Println("Error: ", err)
//			}
//		}
//		fmt.Println("Inserted", len(data), "Rows")
//		return nil
//	}
func SqlScripts(table string, result bson.M) string {
	s := fmt.Sprintf(`
CREATE TABLE %s(`, table)
	for key, value := range result {
		if key == "_id" {
			key = "Mongo_id"
		}
		value_type := getSpannerDataType(reflect.TypeOf(value).String())
		s += fmt.Sprintf(`
	%s %s,`, strings.ToLower(key), value_type)
	}
	s += `
	) PRIMARY KEY (mongo_id);`
	return s
}

func getSpannerDataType(goDataType string) string {
	switch goDataType {
	case "string":
		return "STRING(MAX)"
	case "int64":
		return "INT64"
	case "int32":
		return "INT64"
	case "float64":
		return "FLOAT64"
	case "float32":
		return "FLOAT64"
	case "bool":
		return "BOOL"
	case "time.Time":
		return "TIMESTAMP"
	case "[]byte":
		return "BYTES(MAX)"
	case "primitive.Timestamp":
		return "TIMESTAMP"
	case "primitive.ObjectID":
		return "STRING(24)"
	default:
		return "STRING(MAX)"
	}
}

func InsertScripts(tables []string) string {
	code := `package view

import (
	"fmt"

	"cloud.google.com/go/spanner"
)

type Collection interface {
	InsertSpanner(table string, client *spanner.Client) (*spanner.Mutation, error)
}
`
	for _, value := range tables {
		value = strings.ToUpper(string(value[0])) + value[1:]
		code += fmt.Sprintf(`
func (data *%s) InsertSpanner(table string, client *spanner.Client) (*spanner.Mutation, error) {
	mut, err := spanner.InsertStruct(table, data)
	if err != nil {
		fmt.Println("Error in inserting Struct: ")
		return nil, err
	}
	return mut, nil
}
		`, value)
	}
	code += `
func Insert(table Collection, collection string, client *spanner.Client) (*spanner.Mutation, error) {
	return table.InsertSpanner(collection, client)
}`
	return code
}
