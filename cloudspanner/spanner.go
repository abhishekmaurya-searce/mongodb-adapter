package cloudspanner

import (
	"context"
	"fmt"
	"os"
	"reflect"
	"strings"

	"cloud.google.com/go/spanner"
	"go.mongodb.org/mongo-driver/bson/primitive"

	//"github.com/pratikdhanavesearce/mongodb-adapter/view"
	"gopkg.in/mgo.v2/bson"
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
func SqlScripts(val []string, result bson.M) string {
	table := strings.Join(val, "_")
	s := fmt.Sprintf(`
CREATE TABLE %s(`, strings.ToLower(table))
	if len(val) > 1 {
		s += `
	id STRING(24),`
	}
	for key, value := range result {
		if key == "_id" {
			key = "id"
		}
		value_type := getSpannerDataType(value)
		s += fmt.Sprintf(`
	%s %s,`, strings.ToLower(key), value_type)
	}
	if len(val) > 1 {
		ref := strings.ToLower(strings.Join(val[0:len(val)-1], "_"))
		s += fmt.Sprintf(`
	%s_id STRING(24),
	FOREIGN KEY (%s_id) REFERENCES %s (id) 
	) PRIMARY KEY (id);`, ref, ref, ref)
	} else {
		s += `
	) PRIMARY KEY (id);`
	}
	return s
}

func getSpannerDataType(goDataType interface{}) string {
	value_type := reflect.TypeOf(goDataType).String()
	switch value_type {
	case "primitive.A":
		return "ARRAY<" + getSpannerDataType(goDataType.(primitive.A)[0]) + ">"
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
	case "primitive.DateTime":
		return "TIMESTAMP"
	case "primitive.ObjectID":
		return "STRING(24)"
	case "primitive.Symbol":
		return "STRING(MAX)"
	case "Primitive.Binary":
		return "STRING(MAX)"
	case "primitive.Regex":
		return "STRING(MAX)"
	case "primitve.CodeWithScope":
		return "STRING(MAX)"
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
