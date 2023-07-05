package view

import "go.mongodb.org/mongo-driver/bson/primitive"

type Author struct {
	Firstname    string `bson:"Firstname"`
	Lastname     string `bson:"Lastname"`
	Deletestatus bool   `bson:"Deletestatus"`
	Age          int64  `bson:"Age"`
	Mongo_id     string `bson:"_id"`
	Id           int64  `bson:"Id"`
}
type Books struct {
	Mongo_id        string `bson:"_id"`
	ID              int64  `bson:"ID"`
	Title           string `bson:"Title"`
	AuthorID        int64  `bson:"AuthorID"`
	ISBN            string `bson:"ISBN"`
	TotalCopies     int64  `bson:"TotalCopies"`
	AvailableCopies int64  `bson:"AvailableCopies"`
	DeleteStatus    bool   `bson:"DeleteStatus"`
}
type Borrower struct {
	Deletestatus bool   `bson:"Deletestatus"`
	Mongo_id     string `bson:"_id"`
	ID           int64  `bson:"ID"`
	Firstname    string `bson:"Firstname"`
	Lastname     string `bson:"Lastname"`
	Email        string `bson:"Email"`
}
type Test struct {
	Objectarray primitive.A `bson:"objectarray"`
	Mongo_id    string      `bson:"_id"`
	Codetest    string      `bson:"codetest"`
	Regtest     string      `bson:"regtest"`
	Binarytest  string      `bson:"binarytest"`
	Array       primitive.A `bson:"array"`
	Object      primitive.M `bson:"object"`
}
type Borrow struct {
	BookID       int64  `bson:"BookID"`
	BorrowDate   string `bson:"BorrowDate"`
	ReturnDate   string `bson:"ReturnDate"`
	Deletestatus bool   `bson:"Deletestatus"`
	Mongo_id     string `bson:"_id"`
	ID           int64  `bson:"ID"`
	BorrowerID   int64  `bson:"BorrowerID"`
}
