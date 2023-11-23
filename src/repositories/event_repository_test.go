package repositories

import (
	"encoding/json"
	"errors"
	"fmt"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/ELITE-Kinoticketsystem/Backend-KTS/src/managers"
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
	title := "Test Event"

	dbManager.EXPECT().ExecuteStatement("INSERT INTO events (id, title, start, end, event_type_id, cinema_hall_id) VALUES ($1, $2, $3, $4, $5, $6)", id, title, time, time, id, id).Return(nil, nil)

	EventRepository := EventRepository{DatabaseMgr: dbManager}

	// Create mock event
	event := &schemas.Event{
		Id:           id,
		Title:        title,
		Start:        time,
		End:          time,
		EventTypeId:  cinemaId,
		CinemaHallId: eventTypeId,
	}

	// WHEN
	err := EventRepository.CreateEvent(event)

	// THEN
	if err != nil {
		t.Fail()
	}
}

func TestCreateEventWithErrorInInsert(t *testing.T) {
	// GIVEN
	// create mock controller
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	dbManager := mocks.NewMockDatabaseManagerI(mockCtrl)

	time := time.Now()
	id := &uuid.UUID{}
	cinemaId := &uuid.UUID{}
	eventTypeId := &uuid.UUID{}
	title := "Test Event"

	dbManager.EXPECT().ExecuteStatement("INSERT INTO events (id, title, start, end, event_type_id, cinema_hall_id) VALUES ($1, $2, $3, $4, $5, $6)", id, title, time, time, id, id).Return(nil, errors.New("Error"))

	EventRepository := EventRepository{DatabaseMgr: dbManager}

	// Create mock event
	event := &schemas.Event{
		Id:           id,
		Title:        title,
		Start:        time,
		End:          time,
		EventTypeId:  cinemaId,
		CinemaHallId: eventTypeId,
	}

	// WHEN
	err := EventRepository.CreateEvent(event)

	// THEN
	if err == nil {
		t.Fail()
	}
}

func TestUpdateEvent(t *testing.T) {
	// GIVEN
	// create mock controller
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	dbManager := mocks.NewMockDatabaseManagerI(mockCtrl)

	time := time.Now()
	id := &uuid.UUID{}
	cinemaId := &uuid.UUID{}
	eventTypeId := &uuid.UUID{}
	title := "Test Event"

	result := mocks.NewMockResult(mockCtrl)
	result.EXPECT().RowsAffected().Return(int64(1), nil)

	dbManager.EXPECT().ExecuteStatement("UPDATE events SET title=$1, start=$2, end=$3, event_type_id=$4, cinema_hall_id=$5 WHERE id=$6", title, time, time, id, id, id).Return(result, nil)

	EventRepository := EventRepository{DatabaseMgr: dbManager}

	// Create mock event
	event := &schemas.Event{
		Id:           id,
		Title:        title,
		Start:        time,
		End:          time,
		EventTypeId:  cinemaId,
		CinemaHallId: eventTypeId,
	}

	// WHEN
	err := EventRepository.UpdateEvent(event)

	// THEN
	if err != nil {
		t.Fail()
	}
}

func TestUpdateEventWithErrorInUpdate(t *testing.T) {
	// GIVEN
	// create mock controller
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	dbManager := mocks.NewMockDatabaseManagerI(mockCtrl)

	time := time.Now()
	id := &uuid.UUID{}
	cinemaId := &uuid.UUID{}
	eventTypeId := &uuid.UUID{}
	title := "Test Event"

	result := mocks.NewMockResult(mockCtrl)

	dbManager.EXPECT().ExecuteStatement("UPDATE events SET title=$1, start=$2, end=$3, event_type_id=$4, cinema_hall_id=$5 WHERE id=$6", title, time, time, id, id, id).Return(result, errors.New("Error"))

	EventRepository := EventRepository{DatabaseMgr: dbManager}

	// Create mock event
	event := &schemas.Event{
		Id:           id,
		Title:        title,
		Start:        time,
		End:          time,
		EventTypeId:  cinemaId,
		CinemaHallId: eventTypeId,
	}

	// WHEN
	err := EventRepository.UpdateEvent(event)

	// THEN
	if err == nil {
		t.Fail()
	}
}

func TestUpdateEventWithNoRowsAffected(t *testing.T) {
	// GIVEN
	// create mock controller
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	dbManager := mocks.NewMockDatabaseManagerI(mockCtrl)

	time := time.Now()
	id := &uuid.UUID{}
	cinemaId := &uuid.UUID{}
	eventTypeId := &uuid.UUID{}
	title := "Test Event"

	result := mocks.NewMockResult(mockCtrl)
	result.EXPECT().RowsAffected().Return(int64(0), errors.New("Error"))

	dbManager.EXPECT().ExecuteStatement("UPDATE events SET title=$1, start=$2, end=$3, event_type_id=$4, cinema_hall_id=$5 WHERE id=$6", title, time, time, id, id, id).Return(result, nil)

	EventRepository := EventRepository{DatabaseMgr: dbManager}

	// Create mock event
	event := &schemas.Event{
		Id:           id,
		Title:        title,
		Start:        time,
		End:          time,
		EventTypeId:  cinemaId,
		CinemaHallId: eventTypeId,
	}

	// WHEN
	err := EventRepository.UpdateEvent(event)

	// THEN
	if err == nil {
		t.Fail()
	}
}

func TestDeleteEvent(t *testing.T) {
	//GIVEN
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	dbManager := mocks.NewMockDatabaseManagerI(mockCtrl)

	id := &uuid.UUID{}

	result := mocks.NewMockResult(mockCtrl)
	result.EXPECT().RowsAffected().Return(int64(1), nil)

	dbManager.EXPECT().ExecuteStatement("DELETE FROM events WHERE id=$1", id).Return(result, nil)

	EventRepository := EventRepository{DatabaseMgr: dbManager}

	//WHEN
	err := EventRepository.DeleteEvent(id)

	//THEN
	if err != nil {
		t.Fail()
	}
}

func TestDeleteEventWithErrorInDelete(t *testing.T) {
	//GIVEN
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	dbManager := mocks.NewMockDatabaseManagerI(mockCtrl)

	id := &uuid.UUID{}

	result := mocks.NewMockResult(mockCtrl)

	dbManager.EXPECT().ExecuteStatement("DELETE FROM events WHERE id=$1", id).Return(result, errors.New("Error"))

	EventRepository := EventRepository{DatabaseMgr: dbManager}

	//WHEN
	err := EventRepository.DeleteEvent(id)

	//THEN
	if err == nil {
		t.Fail()
	}
}

func TestDeleteEventWithNoRowsAffected(t *testing.T) {
	//GIVEN
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	dbManager := mocks.NewMockDatabaseManagerI(mockCtrl)

	id := &uuid.UUID{}

	result := mocks.NewMockResult(mockCtrl)
	result.EXPECT().RowsAffected().Return(int64(0), errors.New("Error"))

	dbManager.EXPECT().ExecuteStatement("DELETE FROM events WHERE id=$1", id).Return(result, nil)

	EventRepository := EventRepository{DatabaseMgr: dbManager}

	//WHEN
	err := EventRepository.DeleteEvent(id)

	//THEN
	if err == nil {
		t.Fail()
	}
}

func TestGetEvent(t *testing.T) {
	//GIVEN

	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	id := &uuid.UUID{}

	mock.ExpectQuery("SELECT(.*)").WithArgs(id).WillReturnRows(sqlmock.NewRows([]string{"id", "title", "start", "end", "event_type_id", "cinema_hall_id"}).AddRow(id, "Test Event", time.Now(), time.Now(), id, id))

	dbManager := &managers.DatabaseManager{Connection: db}

	EventRepository := EventRepository{DatabaseMgr: dbManager}

	//WHEN
	event, err := EventRepository.GetEvent(id)

	//THEN
	if err != nil {
		t.Fail()
	}
	if event == nil {
		t.Fail()
	}
}

func TestGetEventWithWrongId(t *testing.T) {
	//GIVEN

	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	id := &uuid.UUID{}

	mock.ExpectQuery("SELECT(.*)").WithArgs(id).WillReturnError(errors.New("Error"))

	dbManager := &managers.DatabaseManager{Connection: db}

	EventRepository := EventRepository{DatabaseMgr: dbManager}

	//WHEN
	event, err := EventRepository.GetEvent(id)

	//THEN
	if err == nil {
		t.Fail()
	}
	if event != nil {
		t.Fail()
	}
}

func TestMyTest(t *testing.T) {
	event := &schemas.Event{
		Title: "Test Event",
	}

	// create a JSON form this event
	eventJson, err := json.Marshal(event)
	if err != nil {
		t.Fail()
	}
	fmt.Println(eventJson)
}
