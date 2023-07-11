package mongodb

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/pratikdhanavesearce/mongodb-adapter/cloudspanner"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"gopkg.in/mgo.v2/bson"
)

func CollectionToStruct(result_original, result bson.M, val []string, original []string) (string, string, string, string, string) {
	var code, sql_code, retrive_code, nesting_code, nesting_code_sql, nesting_code_retrive, inset_code, nesting_insert_code, insert_retrive, nesting_insert_retrive string
	table := val[len(val)-1]
	val[len(val)-1] = strings.ToUpper(string(table[0])) + table[1:]
	name := strings.Join(val, "_")
	code += fmt.Sprintf(`
type %s struct{`, name)
	if len(val) > 1 {
		code += fmt.Sprintf(`
	Id string %s
	%s_Id string %s
	`, "`bson:\"Id\"`", strings.Join(val[0:len(val)-1], "_"), "`bson:\""+strings.Join(val[0:len(val)-1], "_")+"_Id\"`")
	}
	for key, value := range result {
		value_type := reflect.TypeOf(value).String()
		if value_type == "bson.M" {
			str, sql, retr, inst, inst_retr := CollectionToStruct(result_original, value.(bson.M), append(val, key), append(original, key))
			nesting_insert_code += inst
			nesting_code += str
			nesting_code_sql += sql
			nesting_code_retrive += retr
			nesting_insert_retrive += inst_retr
		} else if value_type == "primitive.A" && reflect.TypeOf(value.(primitive.A)[0]).String() == "bson.M" {
			str, sql, retr, inst, inst_retr := CollectionToStruct(result_original, value.(primitive.A)[0].(bson.M), append(val, key), append(original, key))
			nesting_insert_code += inst
			nesting_code += str
			nesting_code_sql += sql
			nesting_code_retrive += retr
			nesting_insert_retrive += inst_retr
		} else {
			code += structcode(key, value)
		}
	}
	code += `
}
` + nesting_code
	sql_code += cloudspanner.SqlScripts(val, result)
	retrive_code += RetriveCollection(result_original, original, val)
	insert_retrive += InsertCode(val)
	insert_retrive += nesting_insert_retrive
	retrive_code += nesting_code_retrive
	sql_code += nesting_code_sql
	inset_code += CodeInsert(val)
	inset_code += nesting_insert_code
	return code, sql_code, retrive_code, inset_code, insert_retrive
}
