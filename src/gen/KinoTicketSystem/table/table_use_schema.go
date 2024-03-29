//
// Code generated by go-jet DO NOT EDIT.
//
// WARNING: Changes to this file may cause incorrect behavior
// and will be lost if the code is regenerated
//

package table

// UseSchema sets a new schema name for all generated table SQL builder types. It is recommended to invoke
// this method only once at the beginning of the program.
func UseSchema(schema string) {
	ActorPictures = ActorPictures.FromSchema(schema)
	Actors = Actors.FromSchema(schema)
	Addresses = Addresses.FromSchema(schema)
	CinemaHalls = CinemaHalls.FromSchema(schema)
	EventMovies = EventMovies.FromSchema(schema)
	EventSeatCategories = EventSeatCategories.FromSchema(schema)
	EventSeats = EventSeats.FromSchema(schema)
	Events = Events.FromSchema(schema)
	Genres = Genres.FromSchema(schema)
	MovieActors = MovieActors.FromSchema(schema)
	MovieGenres = MovieGenres.FromSchema(schema)
	Movies = Movies.FromSchema(schema)
	Orders = Orders.FromSchema(schema)
	PaymentMethods = PaymentMethods.FromSchema(schema)
	PriceCategories = PriceCategories.FromSchema(schema)
	Reviews = Reviews.FromSchema(schema)
	SeatCategories = SeatCategories.FromSchema(schema)
	Seats = Seats.FromSchema(schema)
	Theatres = Theatres.FromSchema(schema)
	Tickets = Tickets.FromSchema(schema)
	Users = Users.FromSchema(schema)
}
