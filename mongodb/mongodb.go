package mongodb

import (
	"context"
	"fmt"
	"log"
	"reflect"
	"strings"

	"github.com/pratikdhanavesearce/mongodb-adapter/cloudspanner"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func NewConnection(str string) (*mongo.Client, error) {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(str)) //Connection to mongodb client which returns a *mongo.Client
	if err != nil {
		return nil, err
	}
	err = client.Ping(context.TODO(), nil) //Double checking if the connection is made or not
	if err != nil {
		return nil, err
	}
	//client.Disconnect(context.TODO())
	return client, nil
}

// Function for Geting all the Collection in the database
func ListCollection(db *mongo.Database) ([]string, error) {

	collections, err := db.ListCollectionNames(context.TODO(), bson.D{}) //Calling the function to get the names of Collection in the database
	if err != nil {
		log.Fatal(err)
		return collections, err
	}
	return collections, nil
}
func Head(db *mongo.Collection) error {
	cursor, err := db.Find(context.TODO(), bson.D{})
	if err != nil {
		return err
	}
	var result []bson.D
	if err = cursor.All(context.TODO(), &result); err != nil {
		return err
	}
	i := 0
	for _, value := range result {
		if i == 5 {
			break
		}
		fmt.Printf("%T", value)
		fmt.Println(value)
		i++
	}
	return nil

}

// func Retrive(db *mongo.Collection) ([]model.Person, error) {
// 	cursor, err := db.Find(context.TODO(), bson.D{})
// 	if err != nil {
// 		return []model.Person{}, err
// 	}
// 	var data []model.Person
// 	err = cursor.All(context.TODO(), &data)
// 	if err != nil {
// 		return data, err
// 	}
// 	return data, nil
// }

func CollectionToStruct(collection mongo.Collection, ctx context.Context, val string) (string, string) {
	var code string
	var sql_code string
	var result bson.M
	_ = collection.FindOne(ctx, bson.M{}).Decode(&result)
	code += fmt.Sprintf(`type %s struct{
	`, (strings.ToUpper(string(val[0])) + val[1:]))
	for key, value := range result {
		key = strings.ToUpper(string(key[0])) + key[1:]
		if reflect.TypeOf(value).String() == "primitive.ObjectID" {
			code += fmt.Sprintf(`%s %s %s:"%s`+`"`+"`"+`
	`, "Mongo_id", "string", "`bson", key)
		} else if reflect.TypeOf(value).String() == "int32" {
			code += fmt.Sprintf(`%s %s %s:"%s`+`"`+"`"+`
			`, key, "int64", "`bson", key)
		} else {
			code += fmt.Sprintf(`%s %s %s:"%s`+`"`+"`"+`
	`, key, reflect.TypeOf(value), "`bson", key)
		}
	}
	code += `
}
`
	sql_code += cloudspanner.SqlScripts(val, result)
	return code, sql_code
}

func RetriveCollection(collection []string) string {
	retrive_code := `package mongodb
	import (
		"github.com/pratikdhanavesearce/mongodb-adapter/view"
		"context"
	
		"go.mongodb.org/mongo-driver/bson"
		"go.mongodb.org/mongo-driver/mongo"
	)`
	for _, val := range collection {
		val = strings.ToUpper(string(val[0])) + val[1:]
		retrive_code += fmt.Sprintf(`
func Retrive_%s(db *mongo.Collection) ([]view.%s, error) {
	cursor, err := db.Find(context.TODO(), bson.D{})
	var data []view.%s
	if err != nil {
		return data, err
	}
	err = cursor.All(context.TODO(), &data)
	if err != nil {
		return data, err
	}
	return data, nil
}`, val, val, val)
	}

	return retrive_code
}
