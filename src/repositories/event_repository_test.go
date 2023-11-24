package repositories

import (
	"errors"
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

	dbManager.EXPECT().ExecuteStatement("INSERT INTO events (id, title, start, end, event_type_id, cinema_hall_id) VALUES (?, ?, ?, ?, ?, ?)", id, title, time, time, id, id).Return(nil, nil)

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

	dbManager.EXPECT().ExecuteStatement("INSERT INTO events (id, title, start, end, event_type_id, cinema_hall_id) VALUES (?, ?, ?, ?, ?, ?)", id, title, time, time, id, id).Return(nil, errors.New("Error"))

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

	dbManager.EXPECT().ExecuteStatement("UPDATE events SET title=?, start=?, end=?, event_type_id=?, cinema_hall_id=? WHERE id=?", title, time, time, id, id, id).Return(result, nil)

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

	dbManager.EXPECT().ExecuteStatement("UPDATE events SET title=?, start=?, end=?, event_type_id=?, cinema_hall_id=? WHERE id=?", title, time, time, id, id, id).Return(result, errors.New("Error"))

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

	dbManager.EXPECT().ExecuteStatement("UPDATE events SET title=?, start=?, end=?, event_type_id=?, cinema_hall_id=? WHERE id=?", title, time, time, id, id, id).Return(result, nil)

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

	dbManager.EXPECT().ExecuteStatement("DELETE FROM events WHERE id=?", id).Return(result, nil)

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

	dbManager.EXPECT().ExecuteStatement("DELETE FROM events WHERE id=?", id).Return(result, errors.New("Error"))

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

	dbManager.EXPECT().ExecuteStatement("DELETE FROM events WHERE id=?", id).Return(result, nil)

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

func TestGetEventsForMovieId(t *testing.T) {
	//GIVEN
	db, mock, _ := sqlmock.New()

	id := &uuid.UUID{}
	// id2 := &uuid.UUID{}

	rows := sqlmock.NewRows([]string{"id", "title", "start", "end", "event_type_id", "cinema_hall_id"}).AddRow(id, "Test Event", time.Now(), time.Now(), id, id).AddRow(id, "Test Event", time.Now(), time.Now(), id, id)

	if rows == nil {
		t.Fail()
	}

	mock.ExpectQuery("SELECT(.*)").WithArgs(id).WillReturnRows(rows)
	dbManager := &managers.DatabaseManager{Connection: db}
	eventRepository := EventRepository{DatabaseMgr: dbManager}

	//WHEN
	events, err := eventRepository.GetEventsForMovieId(id)

	//THEN
	if err != nil {
		t.Fail()
	}
	if events == nil {
		t.Fail()
	}
	if len(events) != 2 {
		t.Fail()
	}
}

func TestGetEventsForMovieIdWithWrongId(t *testing.T) {
	//GIVEN
	db, mock, _ := sqlmock.New()

	id := &uuid.UUID{}

	mock.ExpectQuery("SELECT(.*)").WithArgs(id).WillReturnError(errors.New("Error"))
	dbManager := &managers.DatabaseManager{Connection: db}
	eventRepository := EventRepository{DatabaseMgr: dbManager}

	//WHEN
	events, err := eventRepository.GetEventsForMovieId(id)

	//THEN
	if err == nil {
		t.Fail()
	}
	if events != nil {
		t.Fail()
	}
}

func TestGetEventsDateTimeIsBetween(t *testing.T) {
	//GIVEN
	db, mock, _ := sqlmock.New()

	start := time.Now()
	end := time.Now()

	id := &uuid.UUID{}

	rows := sqlmock.NewRows([]string{"id", "title", "start", "end", "event_type_id", "cinema_hall_id"}).AddRow(id, "Test Event", start, end, id, id).AddRow(id, "Test Event", start, end, id, id)

	if rows == nil {
		t.Fail()
	}

	mock.ExpectQuery("SELECT(.*)").WithArgs(start, end, start, end).WillReturnRows(rows)
	dbManager := &managers.DatabaseManager{Connection: db}
	eventRepository := EventRepository{DatabaseMgr: dbManager}

	//WHEN
	events, err := eventRepository.GetEventsDateTimeIsBetween(start, end)

	//THEN
	if err != nil {
		t.Fail()
	}
	if events == nil {
		t.Fail()
	}
	if len(events) != 2 {
		t.Fail()
	}
}


func TestGetEventsDateTimeIsBetweenWithWrongDates(t *testing.T) {
	//GIVEN
	db, mock, _ := sqlmock.New()

	start := time.Now()
	end := time.Now()

	mock.ExpectQuery("SELECT(.*)").WithArgs(start, end, start, end).WillReturnError(errors.New("Error"))
	dbManager := &managers.DatabaseManager{Connection: db}
	eventRepository := EventRepository{DatabaseMgr: dbManager}

	//WHEN
	events, err := eventRepository.GetEventsDateTimeIsBetween(start, end)

	//THEN
	if err == nil {
		t.Fail()
	}
	if events != nil {
		t.Fail()
	}
}

func TestGetEventsForCinemaHallId(t *testing.T) {
	//GIVEN
	db, mock, _ := sqlmock.New()

	id := &uuid.UUID{}

	rows := sqlmock.NewRows([]string{"id", "title", "start", "end", "event_type_id", "cinema_hall_id"}).AddRow(id, "Test Event", time.Now(), time.Now(), id, id).AddRow(id, "Test Event", time.Now(), time.Now(), id, id)

	if rows == nil {
		t.Fail()
	}

	mock.ExpectQuery("SELECT(.*)").WithArgs(id).WillReturnRows(rows)
	dbManager := &managers.DatabaseManager{Connection: db}
	eventRepository := EventRepository{DatabaseMgr: dbManager}

	//WHEN
	events, err := eventRepository.GetEventsForCinemaHallId(id)

	//THEN
	if err != nil {
		t.Fail()
	}
	if events == nil {
		t.Fail()
	}
	if len(events) != 2 {
		t.Fail()
	}
}

func TestGetEventsForCinemaHallIdWithWrongId(t *testing.T) {
	//GIVEN
	db, mock, _ := sqlmock.New()

	id := &uuid.UUID{}

	mock.ExpectQuery("SELECT(.*)").WithArgs(id).WillReturnError(errors.New("Error"))
	dbManager := &managers.DatabaseManager{Connection: db}
	eventRepository := EventRepository{DatabaseMgr: dbManager}

	//WHEN
	events, err := eventRepository.GetEventsForCinemaHallId(id)

	//THEN
	if err == nil {
		t.Fail()
	}
	if events != nil {
		t.Fail()
	}
}

