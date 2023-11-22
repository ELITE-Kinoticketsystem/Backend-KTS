package schemas

import "github.com/google/uuid"

type User struct {
	Id         *uuid.UUID `json:"id"`
	Username   string     `json:"username"`
	Email      string     `json:"email"`
	Password   string     `json:"password"`
	FirstName  string     `json:"firstName"`
	LastName   string     `json:"lastName"`
	AddressId  *uuid.UUID `json:"addressId"`
	UserTypeId *uuid.UUID `json:"userTypeId"`
}

type UserMovies struct {
	UserId   *uuid.UUID `json:"userId"`
	MovieId  *uuid.UUID `json:"movieId"`
	ListType string     `json:"listType"`
}

type Admin struct {
	Id       *uuid.UUID `json:"id"`
	Email    string
	Password string
	Theatre  *uuid.UUID `json:"theatreId"`
}
