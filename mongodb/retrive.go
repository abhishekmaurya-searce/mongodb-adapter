package mongodb
import (
	"context"
	"encoding/json"
	"reflect"
	
	"github.com/pratikdhanavesearce/mongodb-adapter/view"
	
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	)
func Retrive_Author(db *mongo.Collection) ([]view.Author, error) {
	cursor, err := db.Find(context.TODO(), bson.M{})
	var data []view.Author
	if err != nil {
		return data, err
	}
	for cursor.Next(context.TODO()) {
		var doc bson.M
		if err = cursor.Decode(&doc); err != nil {
			return data,err
		}
		modified := bson.M{}
		for key, value := range doc {
			type_value := reflect.TypeOf(value).String()
			if type_value == "primitive.Binary" || type_value == "primitive.Regex" || type_value == "primitive.CodeWithScope" {
				jsondata, err := json.Marshal(value)
				if err != nil {
					return data,err
				}
				value = string(jsondata)
			}
			modified[key] = value
		}
		var temp view.Author
		document, err := bson.Marshal(modified)
		if err != nil {
			return data,err
		}
		err = bson.Unmarshal(document, &temp)
		if err != nil {
			return data,err
		}
		data = append(data, temp)
	}
	return data, nil
}
func Retrive_Books(db *mongo.Collection) ([]view.Books, error) {
	cursor, err := db.Find(context.TODO(), bson.M{})
	var data []view.Books
	if err != nil {
		return data, err
	}
	for cursor.Next(context.TODO()) {
		var doc bson.M
		if err = cursor.Decode(&doc); err != nil {
			return data,err
		}
		modified := bson.M{}
		for key, value := range doc {
			type_value := reflect.TypeOf(value).String()
			if type_value == "primitive.Binary" || type_value == "primitive.Regex" || type_value == "primitive.CodeWithScope" {
				jsondata, err := json.Marshal(value)
				if err != nil {
					return data,err
				}
				value = string(jsondata)
			}
			modified[key] = value
		}
		var temp view.Books
		document, err := bson.Marshal(modified)
		if err != nil {
			return data,err
		}
		err = bson.Unmarshal(document, &temp)
		if err != nil {
			return data,err
		}
		data = append(data, temp)
	}
	return data, nil
}
func Retrive_Borrower(db *mongo.Collection) ([]view.Borrower, error) {
	cursor, err := db.Find(context.TODO(), bson.M{})
	var data []view.Borrower
	if err != nil {
		return data, err
	}
	for cursor.Next(context.TODO()) {
		var doc bson.M
		if err = cursor.Decode(&doc); err != nil {
			return data,err
		}
		modified := bson.M{}
		for key, value := range doc {
			type_value := reflect.TypeOf(value).String()
			if type_value == "primitive.Binary" || type_value == "primitive.Regex" || type_value == "primitive.CodeWithScope" {
				jsondata, err := json.Marshal(value)
				if err != nil {
					return data,err
				}
				value = string(jsondata)
			}
			modified[key] = value
		}
		var temp view.Borrower
		document, err := bson.Marshal(modified)
		if err != nil {
			return data,err
		}
		err = bson.Unmarshal(document, &temp)
		if err != nil {
			return data,err
		}
		data = append(data, temp)
	}
	return data, nil
}
func Retrive_Test(db *mongo.Collection) ([]view.Test, error) {
	cursor, err := db.Find(context.TODO(), bson.M{})
	var data []view.Test
	if err != nil {
		return data, err
	}
	for cursor.Next(context.TODO()) {
		var doc bson.M
		if err = cursor.Decode(&doc); err != nil {
			return data,err
		}
		modified := bson.M{}
		for key, value := range doc {
			type_value := reflect.TypeOf(value).String()
			if type_value == "primitive.Binary" || type_value == "primitive.Regex" || type_value == "primitive.CodeWithScope" {
				jsondata, err := json.Marshal(value)
				if err != nil {
					return data,err
				}
				value = string(jsondata)
			}
			modified[key] = value
		}
		var temp view.Test
		document, err := bson.Marshal(modified)
		if err != nil {
			return data,err
		}
		err = bson.Unmarshal(document, &temp)
		if err != nil {
			return data,err
		}
		data = append(data, temp)
	}
	return data, nil
}
func Retrive_Borrow(db *mongo.Collection) ([]view.Borrow, error) {
	cursor, err := db.Find(context.TODO(), bson.M{})
	var data []view.Borrow
	if err != nil {
		return data, err
	}
	for cursor.Next(context.TODO()) {
		var doc bson.M
		if err = cursor.Decode(&doc); err != nil {
			return data,err
		}
		modified := bson.M{}
		for key, value := range doc {
			type_value := reflect.TypeOf(value).String()
			if type_value == "primitive.Binary" || type_value == "primitive.Regex" || type_value == "primitive.CodeWithScope" {
				jsondata, err := json.Marshal(value)
				if err != nil {
					return data,err
				}
				value = string(jsondata)
			}
			modified[key] = value
		}
		var temp view.Borrow
		document, err := bson.Marshal(modified)
		if err != nil {
			return data,err
		}
		err = bson.Unmarshal(document, &temp)
		if err != nil {
			return data,err
		}
		data = append(data, temp)
	}
	return data, nil
}