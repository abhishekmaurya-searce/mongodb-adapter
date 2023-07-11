
package mongodb
import (
	"context"
	"encoding/json"
	"reflect"
	"github.com/pratikdhanavesearce/mongodb-adapter/view"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"gopkg.in/mgo.v2/bson"
)

func Retrive_Shipwrecks(db *mongo.Collection) ([]view.Shipwrecks, error) {
	cursor, err := db.Find(context.TODO(), bson.M{})
	var data []view.Shipwrecks
	if err != nil {
		return data, err
	}
	
	for cursor.Next(context.TODO()) {
		var result bson.M
		if err = cursor.Decode(&result); err != nil {
			return data,err
		}
		doc := result
		modified := bson.M{}
		for key, value := range doc {
			type_value := reflect.TypeOf(value).String()
			if type_value == "primitive.Binary" || type_value == "primitive.Regex" || type_value == "primitive.CodeWithScope" || type_value == "primitive.Decimal128" {
				jsondata, err := json.Marshal(value)
				if err != nil {
					return data,err
				}
				value = string(jsondata)
			}else if type_value == "primitive.ObjectID"{
				value = value.(primitive.ObjectID).Hex()
			}else if type_value == "primitive.DateTime"{
				value = value.(primitive.DateTime).Time()
			}
			modified[key] = value
			}
		
		var temp view.Shipwrecks
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