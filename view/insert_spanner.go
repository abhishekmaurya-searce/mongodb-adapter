package view

import (
"fmt"
	
"cloud.google.com/go/spanner"
)
	
type Collection interface {
InsertSpanner(table string, client *spanner.Client) (*spanner.Mutation, error)
}

	func (data *Shipwrecks) InsertSpanner(table string, client *spanner.Client) (*spanner.Mutation, error) {
		mut, err := spanner.InsertStruct(table, data)
		if err != nil {
			fmt.Println("Error in inserting Struct: ")
			return nil, err
		}
		return mut, nil
	}
func Insert(table Collection, collection string, client *spanner.Client) (*spanner.Mutation, error) {
	return table.InsertSpanner(collection, client)
}