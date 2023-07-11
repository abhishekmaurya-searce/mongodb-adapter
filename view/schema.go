package view

type Shipwrecks struct {
	Gp_quality    string    `bson:"gp_quality"`
	Id            string    `bson:"_id"`
	Latdec        float64   `bson:"latdec"`
	Sounding_type string    `bson:"sounding_type"`
	History       string    `bson:"history"`
	Recrd         string    `bson:"recrd"`
	Vesslterms    string    `bson:"vesslterms"`
	Coordinates   []float64 `bson:"coordinates"`
	Depth         int64     `bson:"depth"`
	Quasou        string    `bson:"quasou"`
	Londec        float64   `bson:"londec"`
	Watlev        string    `bson:"watlev"`
	Feature_type  string    `bson:"feature_type"`
	Chart         string    `bson:"chart"`
}
