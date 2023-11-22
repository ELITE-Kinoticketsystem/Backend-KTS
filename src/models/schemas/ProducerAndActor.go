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

type MovieActors struct {
	MovieId *uuid.UUID `json:"movieId"`
	ActorId *uuid.UUID `json:"actorId"`
}

type MovieProducers struct {
	MovieId    *uuid.UUID `json:"movieId"`
	ProducerId *uuid.UUID `json:"producerId"`
}
