package model

type Sub struct {
	Second    int64  `json:"second"`
	Original  string `json:"original"`
	Translate string `json:"translate"`
}

type SubDB struct {
	Id        int    `json:"id"`
	Original  string `json:"original"`
	Translate string `json:"translate"`
}
