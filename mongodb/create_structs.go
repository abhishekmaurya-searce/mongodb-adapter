package mongodb

import (
	"fmt"
	"reflect"
	"strings"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func arraytype(value interface{}) string {
	value_type := reflect.TypeOf(value).String()
	if value_type == "primitive.A" {
		return "[]" + arraytype(value.(primitive.A)[0])
	} else if value_type == "primitive.Timestamp" || value_type == "primitive.DateTime" {
		return "time.Time"
	} else if value_type == "primitive.ObjectID" || value_type == "primitive.Symbol" || value_type == "primitive.CodeWithScope" || value_type == "primitive.Binary" || value_type == "primitive.Regex" {
		return "string"
	} else if value_type == "int32" {
		return "int64"
	} else {
		return value_type
	}
}
func structcode(key string, value interface{}) string {
	structkey := strings.ToUpper(string(key[0])) + key[1:]
	value_type := reflect.TypeOf(value).String()
	if value_type == "primitive.A" {
		return structkey + arraytype(value.(primitive.A)) + "`bson:\"" + key + "\"`" + `
		`
	} else if value_type == "primitive.Timestamp" || value_type == "primitive.DateTime" {
		return fmt.Sprintf(`%s %s %s:"%s`+`"`+"`"+`
		`, structkey, "time.Time", "`bson", key)
	} else if value_type == "primitive.ObjectID" && key == "_id" {
		return fmt.Sprintf(`%s %s %s:"%s`+`"`+"`"+`
		`, "Id", "string", "`bson", key)
	} else if value_type == "primitive.ObjectID" || value_type == "primitive.Symbol" || value_type == "primitive.CodeWithScope" || value_type == "primitive.Binary" || value_type == "primitive.Regex" || value_type == "primitive.Decimal128" {
		return fmt.Sprintf(`%s %s %s:"%s`+`"`+"`"+`
		`, structkey, "string", "`bson", key)
	} else if value_type == "int32" {
		return fmt.Sprintf(`%s %s %s:"%s`+`"`+"`"+`
		`, structkey, "int64", "`bson", key)
	} else {
		return fmt.Sprintf(`%s %s %s:"%s`+`"`+"`"+`
		`, structkey, value_type, "`bson", key)
	}
}
