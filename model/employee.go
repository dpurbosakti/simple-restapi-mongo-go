package model

type Employee struct {
	ID         string `json:"id,omitempty" bson:"id"`
	Name       string `json:"name,omitempty" bson:"name"`
	Department string `json:"department,omitempty" bson:"department"`
}
