package main

type User struct {
	ID        int     `json:"id"`
	FirstName string  `json:"firstName"`
	LastName  string  `json:"LastName"`
	Email     string  `json:"email"`
	Password  string  `json:"password"` // TODO: hash or what?
	Address   Address `json:"address"`
}

type Customer struct {
	WatchedMovies   *[]Movie `json:"watchedMovies"`
	SuggestedMovies *[]Movie `json:"suggestedMovies"`
	SavedMovies     *[]Movie `json:"savedMovies"`
	Orders          *[]Order `json:"orders"`
	Cart            *[]Cart  `json:"card"`
	User
}

type Employee struct {
	Customer
}

type Admin struct {
	Employee
}

type Cashier struct {
	Employee
}