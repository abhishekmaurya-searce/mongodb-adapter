
CREATE TABLE shipwrecks(
	feature_type STRING(MAX),
	chart STRING(MAX),
	londec FLOAT64,
	watlev STRING(MAX),
	id STRING(24),
	latdec FLOAT64,
	gp_quality STRING(MAX),
	recrd STRING(MAX),
	vesslterms STRING(MAX),
	sounding_type STRING(MAX),
	history STRING(MAX),
	depth INT64,
	quasou STRING(MAX),
	coordinates ARRAY<FLOAT64>,
	) PRIMARY KEY (id);