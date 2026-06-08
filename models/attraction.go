package models

type AttractionFeature struct {
	Type       string             `json:"type"`
	ID         string             `json:"id"`
	Properties AttractionProperty `json:"properties"`
	Geometry   AttractionGeometry `json:"geometry"`
}

type AttractionProperty struct {
	XID      string  `json:"xid"`
	Name     string  `json:"name"`
	Dist     float64 `json:"dist"`
	Rate     int     `json:"rate"`
	Wikidata string  `json:"wikidata"`
	Kinds    string  `json:"kinds"`
}

type AttractionGeometry struct {
	Type        string    `json:"type"`
	Coordinates []float64 `json:"coordinates"`
}

type AttractionResponse struct {
	Type     string              `json:"type"`
	Features []AttractionFeature `json:"features"`
}

type AttractionDTO struct {
	Name  string `json:"name"`
	Kinds string `json:"kinds"`
	XID   string `json:"xid"`
}
