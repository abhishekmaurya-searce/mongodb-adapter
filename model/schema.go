package model

type Person struct {
	Id    int    `bson:"id"`
	First string `bson:"firstname"`
	Last  string `bson:"lastname"`
	Age   int    `bson:"age"`
}
