package models

type Flat struct {
	FlatId  int    `json:"id"`
	HouseId int    `json:"house_id"`
	Price   int    `json:"price"`
	Rooms   int    `json:"rooms"`
	Status  string `json:"status"`
}

type FlatInputObject struct {
	HouseId int `json:"house_id"`
	Price   int `json:"price"`
	Rooms   int `json:"rooms"`
}

type FlatUpdateObject struct {
	HouseId int `json:"house_id"`
	FlatId  int `json:"id"`
	Price   int `json:"price"`
	Rooms   int `json:"rooms"`
}
