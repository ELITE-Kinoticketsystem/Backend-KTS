// package repositories

// import (
// 	"log"

// 	"github.com/ELITE-Kinoticketsystem/Backend-KTS/src/managers"
// 	"github.com/ELITE-Kinoticketsystem/Backend-KTS/src/models/schemas"
// 	"github.com/google/uuid"
// )

// // Structs representing the database tables
// // type Event struct {
// // 	Id           *uuid.UUID `json:"id"`
// // 	Title        string     `json:"title"`
// // 	Start        time.Time  `json:"start"`
// // 	End          time.Time  `json:"end"`
// // 	Price        int        `json:"price"` // requires conversion
// // 	EventTypeId  *uuid.UUID `json:"eventTypeId"`
// // 	CinemaHallId *uuid.UUID `json:"cinemaHallId"`
// // }

// // type EventType struct {
// // 	Id       *uuid.UUID `json:"id"`
// // 	TypeName string     `json:"typeName"`
// // }

// // type EventMovies struct {
// // 	EventId *uuid.UUID `json:"eventId"`
// // 	MovieId *uuid.UUID `json:"movieId"`
// // }

// type EventByIdResponse struct {
// 	Id 		 *uuid.UUID `json:"id"`
// 	Title 	 string     `json:"title"`
// 	Start 	 time.Time  `json:"start"`
// 	End 	 time.Time  `json:"end"`
// 	Price 	 int        `json:"price"`

// type EventRepository interface {
// 	GetShowingById(id *uuid.UUID) (*schemas.Event, error)
// }

// type EventRepositoryImpl struct {
// 	DatabaseMgr managers.DatabaseManagerI
// }

// func (er *EventRepositoryImpl) GetShowingById(id *uuid.UUID) (*schemas.Event, error) {

// 	query := "SELECT * FROM events WHERE id = ?"
// 	row := er.DatabaseMgr.ExecuteQueryRow(query, id)

// 	var event schemas.Event
// 	err := row.Scan(&event.Id, &event.Title, &event.Start, &event.End, &event.Price, &event.CinemaHallId)
// 	if err != nil {
// 		log.Println(err)
// 		return nil, err
// 	}

// 	return &event, nil
// }
