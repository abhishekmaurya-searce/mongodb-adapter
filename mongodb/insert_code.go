package mongodb

import (
	"fmt"
	"strings"
)

func CodeInsert(table []string) string {
	return fmt.Sprintf(`
	func (data *%s) InsertSpanner(table string, client *spanner.Client) (*spanner.Mutation, error) {
		mut, err := spanner.InsertStruct(table, data)
		if err != nil {
			fmt.Println("Error in inserting Struct: ")
			return nil, err
		}
		return mut, nil
	}`, strings.Join(table, "_"))
}
func InsertCode(tables []string) string {
	var code string
	for i, val := range tables {
		tables[i] = strings.ToUpper(string(val[0])) + val[1:]
	}
	table := strings.Join(tables, "_")
	if len(tables) > 1 {
		code += fmt.Sprintf(`
		arr_%s, err := mongodb.Retrive_%s(db.Collection(table),arr_%s)
			fmt.Println("Inserting into ", "%s")
			if err != nil {
				fmt.Println("Error in retriving data", err)
			}
			mut = mut[:0]
			for _, value := range arr_%s {
				temp, err := view.Insert(&value, "%s", spanner_client)
				if err != nil {
					fmt.Println(err)
				}
				mut = append(mut, temp)
			}
			_, err = spanner_client.Apply(context.TODO(), mut)
			if err != nil {
				fmt.Println("Error: ", err)
			}`, table, table, strings.Join(tables[0:len(tables)-1], "_"), table, table, strings.ToLower(table))
	}
	return code
}
