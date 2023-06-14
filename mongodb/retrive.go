package mongodb
	import (
		"github.com/pratikdhanavesearce/mongodb-adapter/view"
		"context"
	
		"go.mongodb.org/mongo-driver/bson"
		"go.mongodb.org/mongo-driver/mongo"
	)
func Retrive_Author(db *mongo.Collection) ([]view.Author, error) {
	cursor, err := db.Find(context.TODO(), bson.D{})
	var data []view.Author
	if err != nil {
		return data, err
	}
	err = cursor.All(context.TODO(), &data)
	if err != nil {
		return data, err
	}
	return data, nil
}
func Retrive_Books(db *mongo.Collection) ([]view.Books, error) {
	cursor, err := db.Find(context.TODO(), bson.D{})
	var data []view.Books
	if err != nil {
		return data, err
	}
	err = cursor.All(context.TODO(), &data)
	if err != nil {
		return data, err
	}
	return data, nil
}
func Retrive_Borrower(db *mongo.Collection) ([]view.Borrower, error) {
	cursor, err := db.Find(context.TODO(), bson.D{})
	var data []view.Borrower
	if err != nil {
		return data, err
	}
	err = cursor.All(context.TODO(), &data)
	if err != nil {
		return data, err
	}
	return data, nil
}
func Retrive_Borrow(db *mongo.Collection) ([]view.Borrow, error) {
	cursor, err := db.Find(context.TODO(), bson.D{})
	var data []view.Borrow
	if err != nil {
		return data, err
	}
	err = cursor.All(context.TODO(), &data)
	if err != nil {
		return data, err
	}
	return data, nil
}