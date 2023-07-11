package mongodb

import (
	"fmt"
	"reflect"
	"strings"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"gopkg.in/mgo.v2/bson"
)

func RetriveCollection(result bson.M, tables, columns []string) string {
	var retrive_code, nested_doc, atr, keys, i, atr_type string
	table := strings.Join(columns, "_")
	if len(tables) > 1 {
		for _, val := range tables[1:] {
			var types string
			if reflect.TypeOf(result[val]).String() == "bson.M" {
				types = "(bson.M)"
				result = result[val].(bson.M)
			} else if reflect.TypeOf(result[val]).String() == "primitive.A" {
				types = "(primitive.A)"
				array := result[val].(primitive.A)
				result = array[0].(bson.M)
			}
			nested_doc += "[\"" + val + "\"]." + types
			atr_type += val
		}
		atr = fmt.Sprintf(`, ref []view.%s`, strings.Join(columns[:len(tables)-1], "_"))
		//ref := strings.Join(tables[0:len(tables)-1], "_")
		keys = fmt.Sprintf(`
	modified["Id"]= primitive.NewObjectID().Hex()
	modified["%s_Id"]=ref[i].Id
	i++`, strings.Join(columns[0:len(columns)-1], "_"))
		i = `
	i:=0`
	}
	var code string
	if len(tables) > 1 && nested_doc[len(nested_doc)-2] == 'A' {
		code = `for _, val := range doc {
			for key, value := range val.(bson.M) {
				type_value := reflect.TypeOf(value).String()
				if type_value == "primitive.Binary" || type_value == "primitive.Regex" || type_value == "primitive.CodeWithScope" || type_value == "primitive.Decimal128" {
					jsondata, err := json.Marshal(value)
					if err != nil {
						return data, err
					}
					value = string(jsondata)
				} else if type_value == "primitive.ObjectID" {
					value = value.(primitive.ObjectID).Hex()
				}else if type_value == "primitive.DateTime"{
					value = value.(primitive.DateTime).Time()
				}
				modified[key] = value
			}
		}`
	} else {
		code = `for key, value := range doc {
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
			}`
	}
	retrive_code += fmt.Sprintf(`
func Retrive_%s(db *mongo.Collection%s) ([]view.%s, error) {
	cursor, err := db.Find(context.TODO(), bson.M{})
	var data []view.%s
	if err != nil {
		return data, err
	}
	%s
	for cursor.Next(context.TODO()) {
		var result bson.M
		if err = cursor.Decode(&result); err != nil {
			return data,err
		}
		doc := result`+nested_doc+`
		modified := bson.M{}
		%s
		%s
		var temp view.%s
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
}`, table, atr, table, table, i, code, keys, table)
	return retrive_code
}
