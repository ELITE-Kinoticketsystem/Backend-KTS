package schemas

import "github.com/google/uuid"

type Producer struct {
	Id   *uuid.UUID `json:"id"`
	Name string     `json:"name"`
	Age  int        `json:"age"`
}

type Actor struct {
	Id   *uuid.UUID `json:"id"`
	Name string     `json:"name"`
	Age  int        `json:"age"`
}
