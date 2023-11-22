package repositories

import (
	"fmt"
	"testing"
	"time"

	"github.com/ELITE-Kinoticketsystem/Backend-KTS/src/mocks"
	"github.com/ELITE-Kinoticketsystem/Backend-KTS/src/models/schemas"
	"github.com/google/uuid"
	"go.uber.org/mock/gomock"
)

// type Event struct {
// 	Id           *uuid.UUID `json:"id"`
// 	Title        string     `json:"title"`
// 	Start        time.Time  `json:"start"`
// 	End          time.Time  `json:"end"`
// 	Price        int        `json:"price"` // requires conversion
// 	EventTypeId  *uuid.UUID `json:"eventTypeId"`
// 	CinemaHallId *uuid.UUID `json:"cinemaHallId"`
// }

// type EventType struct {
// 	Id       *uuid.UUID `json:"id"`
// 	TypeName string     `json:"typeName"`
// }

func TestCreateEvent(t *testing.T) {
	// GIVEN

	// create mock controller
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	dbManager := mocks.NewMockDatabaseManagerI(mockCtrl)

	time := time.Now()
	id := &uuid.UUID{}
	cinemaId := &uuid.UUID{}
	eventTypeId := &uuid.UUID{}

	dbManager.EXPECT().ExecuteStatement("INSERT INTO events (id, title, start_time, end_time, price, cinema_hall_id) VALUES (?, ?, ?, ?, ?, ?);", id, "Test Event", time, time, 100, cinemaId, eventTypeId).Return(nil, nil)

	EventRepository := EventRepository{DatabaseMgr: dbManager}

	// Create mock event
	event := &schemas.Event{
		Id:           id,
		Title:        "Test Event",
		Start:        time,
		End:          time,
		Price:        100,
		EventTypeId:  cinemaId,
		CinemaHallId: eventTypeId,
	}

	// WHEN
	err := EventRepository.CreateEvent(event)

	fmt.Println(err)

	// THEN
	// Write an expect with the err, if the test fails, the error should be printed
	if err != nil {
		// t.Fail()
	}
}
