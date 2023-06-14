
CREATE TABLE author(
	mongo_id STRING(24),
	id INT64,
	firstname STRING(MAX),
	lastname STRING(MAX),
	deletestatus BOOL,
	age INT64,
	) PRIMARY KEY (mongo_id);
CREATE TABLE books(
	id INT64,
	title STRING(MAX),
	authorid INT64,
	isbn STRING(MAX),
	totalcopies INT64,
	availablecopies INT64,
	deletestatus BOOL,
	mongo_id STRING(24),
	) PRIMARY KEY (mongo_id);
CREATE TABLE borrower(
	firstname STRING(MAX),
	lastname STRING(MAX),
	email STRING(MAX),
	deletestatus BOOL,
	mongo_id STRING(24),
	id INT64,
	) PRIMARY KEY (mongo_id);
CREATE TABLE borrow(
	borrowerid INT64,
	bookid INT64,
	borrowdate STRING(MAX),
	returndate STRING(MAX),
	deletestatus BOOL,
	mongo_id STRING(24),
	id INT64,
	) PRIMARY KEY (mongo_id);