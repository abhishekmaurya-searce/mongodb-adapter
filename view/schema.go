package view

import "go.mongodb.org/mongo-driver/bson/primitive"

type Author struct {
	Id           string `bson:"_id"`
	Firstname    string `bson:"Firstname"`
	Lastname     string `bson:"Lastname"`
	Deletestatus bool   `bson:"Deletestatus"`
	Age          int64  `bson:"Age"`
}

type Books struct {
	Title           string `bson:"Title"`
	AuthorID        int64  `bson:"AuthorID"`
	ISBN            string `bson:"ISBN"`
	TotalCopies     int64  `bson:"TotalCopies"`
	AvailableCopies int64  `bson:"AvailableCopies"`
	DeleteStatus    bool   `bson:"DeleteStatus"`
	Id              string `bson:"_id"`
	ID              int64  `bson:"ID"`
}

type Borrower struct {
	Lastname     string `bson:"Lastname"`
	Email        string `bson:"Email"`
	Deletestatus bool   `bson:"Deletestatus"`
	Id           string `bson:"_id"`
	ID           int64  `bson:"ID"`
	Firstname    string `bson:"Firstname"`
}

type Test struct {
	Id         string      `bson:"_id"`
	Codetest   string      `bson:"codetest"`
	Regtest    string      `bson:"regtest"`
	Binarytest string      `bson:"binarytest"`
	Array      primitive.A `bson:"array"`
}

type Test_Objectarray struct {
	Id      string
	Test_id string
	Name    string `bson:"name"`
	Age     int64  `bson:"age"`
}

type Test_Object struct {
	Id      string
	Test_id string
	Age     int64  `bson:"age"`
	Name    string `bson:"name"`
}

type Test_Object_Fullname struct {
	Id        string
	Object_id string
	First     string `bson:"first"`
	Last      string `bson:"last"`
}

type Borrow struct {
	BorrowerID   int64  `bson:"BorrowerID"`
	BookID       int64  `bson:"BookID"`
	BorrowDate   string `bson:"BorrowDate"`
	ReturnDate   string `bson:"ReturnDate"`
	Deletestatus bool   `bson:"Deletestatus"`
	Id           string `bson:"_id"`
	ID           int64  `bson:"ID"`
}
