// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

type Dog struct {
	ID        string `json:"_id" bson:"_id"`
	Name      string `json:"name"`
	IsGoodBoy bool   `json:"isGoodBoy"`
}

type NewDog struct {
	Name      string `json:"name"`
	IsGoodBoy bool   `json:"isGoodBoy"`
}

type NewUser struct {
	Name string    `json:"name"`
	Dogs []*NewDog `json:"dogs"`
}

type User struct {
	ID   string `json:"_id" bson:"_id"`
	Name string `json:"name"`
	Dogs []*Dog `json:"dogs"`
}
