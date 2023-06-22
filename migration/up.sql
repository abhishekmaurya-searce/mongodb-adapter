
CREATE TABLE author(
	mongo_id STRING(24),
	id INT64,
	firstname STRING(MAX),
	lastname STRING(MAX),
	deletestatus BOOL,
	age INT64,
	) PRIMARY KEY (mongo_id);
CREATE TABLE books(
	deletestatus BOOL,
	mongo_id STRING(24),
	id INT64,
	title STRING(MAX),
	authorid INT64,
	isbn STRING(MAX),
	totalcopies INT64,
	availablecopies INT64,
	) PRIMARY KEY (mongo_id);
CREATE TABLE borrower(
	mongo_id STRING(24),
	id INT64,
	firstname STRING(MAX),
	lastname STRING(MAX),
	email STRING(MAX),
	deletestatus BOOL,
	) PRIMARY KEY (mongo_id);
CREATE TABLE test(
	mongo_id STRING(24),
	codetest STRING(MAX),
	regtest STRING(MAX),
	binarytest STRING(MAX),
	) PRIMARY KEY (mongo_id);
CREATE TABLE borrow(
	borrowdate STRING(MAX),
	returndate STRING(MAX),
	deletestatus BOOL,
	mongo_id STRING(24),
	id INT64,
	borrowerid INT64,
	bookid INT64,
	) PRIMARY KEY (mongo_id);