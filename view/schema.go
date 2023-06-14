package view
type Author struct{
	Firstname string `bson:"Firstname"`
	Lastname string `bson:"Lastname"`
	Deletestatus bool `bson:"Deletestatus"`
	Age int64 `bson:"Age"`
			Mongo_id string `bson:"_id"`
	Id int64 `bson:"Id"`
			
}
type Books struct{
	AvailableCopies int64 `bson:"AvailableCopies"`
			DeleteStatus bool `bson:"DeleteStatus"`
	Mongo_id string `bson:"_id"`
	ID int64 `bson:"ID"`
			Title string `bson:"Title"`
	AuthorID int64 `bson:"AuthorID"`
			ISBN string `bson:"ISBN"`
	TotalCopies int64 `bson:"TotalCopies"`
			
}
type Borrower struct{
	ID int64 `bson:"ID"`
			Firstname string `bson:"Firstname"`
	Lastname string `bson:"Lastname"`
	Email string `bson:"Email"`
	Deletestatus bool `bson:"Deletestatus"`
	Mongo_id string `bson:"_id"`
	
}
type Borrow struct{
	BorrowDate string `bson:"BorrowDate"`
	ReturnDate string `bson:"ReturnDate"`
	Deletestatus bool `bson:"Deletestatus"`
	Mongo_id string `bson:"_id"`
	ID int64 `bson:"ID"`
			BorrowerID int64 `bson:"BorrowerID"`
			BookID int64 `bson:"BookID"`
			
}
