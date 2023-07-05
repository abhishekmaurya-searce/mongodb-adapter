package mongodb

import (
	"context"
	"encoding/json"
	"reflect"

	"github.com/google/uuid"
	"github.com/pratikdhanavesearce/mongodb-adapter/view"
	"go.mongodb.org/mongo-driver/mongo"
	"gopkg.in/mgo.v2/bson"
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
			return data, err
		}
		doc = doc["Author"].(bson.M)
		modified := bson.M{}
		for key, value := range doc {
			type_value := reflect.TypeOf(value).String()
			if type_value == "primitive.Binary" || type_value == "primitive.Regex" || type_value == "primitive.CodeWithScope" {
				jsondata, err := json.Marshal(value)
				if err != nil {
					return data, err
				}
				value = string(jsondata)
			}
			modified[key] = value
		}

		var temp view.Author
		document, err := bson.Marshal(modified)
		if err != nil {
			return data, err
		}
		err = bson.Unmarshal(document, &temp)
		if err != nil {
			return data, err
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
			return data, err
		}
		doc = doc["Books"].(bson.M)
		modified := bson.M{}
		for key, value := range doc {
			type_value := reflect.TypeOf(value).String()
			if type_value == "primitive.Binary" || type_value == "primitive.Regex" || type_value == "primitive.CodeWithScope" {
				jsondata, err := json.Marshal(value)
				if err != nil {
					return data, err
				}
				value = string(jsondata)
			}
			modified[key] = value
		}

		var temp view.Books
		document, err := bson.Marshal(modified)
		if err != nil {
			return data, err
		}
		err = bson.Unmarshal(document, &temp)
		if err != nil {
			return data, err
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
			return data, err
		}
		doc = doc["Borrower"].(bson.M)
		modified := bson.M{}
		for key, value := range doc {
			type_value := reflect.TypeOf(value).String()
			if type_value == "primitive.Binary" || type_value == "primitive.Regex" || type_value == "primitive.CodeWithScope" {
				jsondata, err := json.Marshal(value)
				if err != nil {
					return data, err
				}
				value = string(jsondata)
			}
			modified[key] = value
		}

		var temp view.Borrower
		document, err := bson.Marshal(modified)
		if err != nil {
			return data, err
		}
		err = bson.Unmarshal(document, &temp)
		if err != nil {
			return data, err
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
			return data, err
		}
		doc = doc["Test"].(bson.M)
		modified := bson.M{}
		for key, value := range doc {
			type_value := reflect.TypeOf(value).String()
			if type_value == "primitive.Binary" || type_value == "primitive.Regex" || type_value == "primitive.CodeWithScope" {
				jsondata, err := json.Marshal(value)
				if err != nil {
					return data, err
				}
				value = string(jsondata)
			}
			modified[key] = value
		}

		var temp view.Test
		document, err := bson.Marshal(modified)
		if err != nil {
			return data, err
		}
		err = bson.Unmarshal(document, &temp)
		if err != nil {
			return data, err
		}
		data = append(data, temp)
	}
	return data, nil
}
func Retrive_Test_Objectarray(db *mongo.Collection, ref []interface{}) ([]view.Test_Objectarray, error) {
	cursor, err := db.Find(context.TODO(), bson.M{})
	var data []view.Test_Objectarray
	if err != nil {
		return data, err
	}

	i := 0
	for cursor.Next(context.TODO()) {
		var doc bson.M
		if err = cursor.Decode(&doc); err != nil {
			return data, err
		}
		doc = doc["Test"].(bson.M)["Objectarray"].(bson.M)
		modified := bson.M{}
		for key, value := range doc {
			type_value := reflect.TypeOf(value).String()
			if type_value == "primitive.Binary" || type_value == "primitive.Regex" || type_value == "primitive.CodeWithScope" {
				jsondata, err := json.Marshal(value)
				if err != nil {
					return data, err
				}
				value = string(jsondata)
			}
			modified[key] = value
		}

		modified["Id"] = uuid.New().String()
		modified["Test_id"] = ref[i].(bson.M)["Id"]
		i++
		var temp view.Test_Objectarray
		document, err := bson.Marshal(modified)
		if err != nil {
			return data, err
		}
		err = bson.Unmarshal(document, &temp)
		if err != nil {
			return data, err
		}
		data = append(data, temp)
	}
	return data, nil
}
func Retrive_Test_Object(db *mongo.Collection, ref []interface{}) ([]view.Test_Object, error) {
	cursor, err := db.Find(context.TODO(), bson.M{})
	var data []view.Test_Object
	if err != nil {
		return data, err
	}

	i := 0
	for cursor.Next(context.TODO()) {
		var doc bson.M
		if err = cursor.Decode(&doc); err != nil {
			return data, err
		}
		doc = doc["Test"].(bson.M)["Object"].(bson.M)
		modified := bson.M{}
		for key, value := range doc {
			type_value := reflect.TypeOf(value).String()
			if type_value == "primitive.Binary" || type_value == "primitive.Regex" || type_value == "primitive.CodeWithScope" {
				jsondata, err := json.Marshal(value)
				if err != nil {
					return data, err
				}
				value = string(jsondata)
			}
			modified[key] = value
		}

		modified["Id"] = uuid.New().String()
		modified["Test_id"] = ref[i].(bson.M)["Id"]
		i++
		var temp view.Test_Object
		document, err := bson.Marshal(modified)
		if err != nil {
			return data, err
		}
		err = bson.Unmarshal(document, &temp)
		if err != nil {
			return data, err
		}
		data = append(data, temp)
	}
	return data, nil
}
func Retrive_Test_Object_Fullname(db *mongo.Collection, ref []interface{}) ([]view.Test_Object_Fullname, error) {
	cursor, err := db.Find(context.TODO(), bson.M{})
	var data []view.Test_Object_Fullname
	if err != nil {
		return data, err
	}

	i := 0
	for cursor.Next(context.TODO()) {
		var doc bson.M
		if err = cursor.Decode(&doc); err != nil {
			return data, err
		}
		doc = doc["Test"].(bson.M)["Object"].(bson.M)["Fullname"].(bson.M)
		modified := bson.M{}
		for key, value := range doc {
			type_value := reflect.TypeOf(value).String()
			if type_value == "primitive.Binary" || type_value == "primitive.Regex" || type_value == "primitive.CodeWithScope" {
				jsondata, err := json.Marshal(value)
				if err != nil {
					return data, err
				}
				value = string(jsondata)
			}
			modified[key] = value
		}

		modified["Id"] = uuid.New().String()
		modified["Test_Object_id"] = ref[i].(bson.M)["Id"]
		i++
		var temp view.Test_Object_Fullname
		document, err := bson.Marshal(modified)
		if err != nil {
			return data, err
		}
		err = bson.Unmarshal(document, &temp)
		if err != nil {
			return data, err
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
			return data, err
		}
		doc = doc["Borrow"].(bson.M)
		modified := bson.M{}
		for key, value := range doc {
			type_value := reflect.TypeOf(value).String()
			if type_value == "primitive.Binary" || type_value == "primitive.Regex" || type_value == "primitive.CodeWithScope" {
				jsondata, err := json.Marshal(value)
				if err != nil {
					return data, err
				}
				value = string(jsondata)
			}
			modified[key] = value
		}

		var temp view.Borrow
		document, err := bson.Marshal(modified)
		if err != nil {
			return data, err
		}
		err = bson.Unmarshal(document, &temp)
		if err != nil {
			return data, err
		}
		data = append(data, temp)
	}
	return data, nil
}
