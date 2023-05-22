package mongodb

import (
	"context"
	"fmt"
	"log"

	"github.com/pratikdhanavesearce/mongodb-adapter/model"
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

func Retrive(db *mongo.Collection) ([]model.Person, error) {
	cursor, err := db.Find(context.TODO(), bson.D{})
	if err != nil {
		return []model.Person{}, err
	}
	var people []model.Person
	err = cursor.All(context.TODO(), &people)
	if err != nil {
		return people, err
	}
	return people, nil
}
