package schemas

import "github.com/google/uuid"

type User struct {
	Id        *uuid.UUID `json:"id"`
	FirstName string     `json:"firstName"`
	LastName  string     `json:"lastName"`
	Email     string     `json:"email"`
	Age       int        `json:"age"`
	Password  string     `json:"password"` // TODO: hash or what?
	AddressId *uuid.UUID `json:"addressId"`
}

type UserMovies struct {
	Id       *uuid.UUID `json:"id"` // probably not needed
	UserId   *uuid.UUID `json:"userId"`
	MovieId  *uuid.UUID `json:"movieId"`
	ListType string     `json:"listType"`
}
