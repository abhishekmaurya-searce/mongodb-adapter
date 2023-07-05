
CREATE TABLE author(
	firstname STRING(MAX),
	lastname STRING(MAX),
	deletestatus BOOL,
	age INT64,
	id STRING(24),
	id INT64,
	) PRIMARY KEY (id);
CREATE TABLE books(
	title STRING(MAX),
	authorid INT64,
	isbn STRING(MAX),
	totalcopies INT64,
	availablecopies INT64,
	deletestatus BOOL,
	id STRING(24),
	id INT64,
	) PRIMARY KEY (id);
CREATE TABLE borrower(
	id STRING(24),
	id INT64,
	firstname STRING(MAX),
	lastname STRING(MAX),
	email STRING(MAX),
	deletestatus BOOL,
	) PRIMARY KEY (id);
CREATE TABLE test(
	objectarray ARRAY<STRING(MAX)>,
	id STRING(24),
	codetest STRING(MAX),
	regtest STRING(MAX),
	binarytest STRING(MAX),
	array ARRAY<INT64>,
	object STRING(MAX),
	) PRIMARY KEY (id);
CREATE TABLE test_objectarray(
	id STRING(24),
	name STRING(MAX),
	age INT64,
	test_id STRING(24),
	FOREIGN KEY (test_id) REFERENCES test (id) 
	) PRIMARY KEY (id);
CREATE TABLE test_object(
	id STRING(24),
	age INT64,
	name STRING(MAX),
	fullname STRING(MAX),
	test_id STRING(24),
	FOREIGN KEY (test_id) REFERENCES test (id) 
	) PRIMARY KEY (id);
CREATE TABLE test_object_fullname(
	id STRING(24),
	first STRING(MAX),
	last STRING(MAX),
	test_object_id STRING(24),
	FOREIGN KEY (test_object_id) REFERENCES test_object (id) 
	) PRIMARY KEY (id);
CREATE TABLE borrow(
	deletestatus BOOL,
	id STRING(24),
	id INT64,
	borrowerid INT64,
	bookid INT64,
	borrowdate STRING(MAX),
	returndate STRING(MAX),
	) PRIMARY KEY (id);