package repositories

import (
	"database/sql"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"

	kts_errors "github.com/ELITE-Kinoticketsystem/Backend-KTS/src/errors"
	"github.com/ELITE-Kinoticketsystem/Backend-KTS/src/gen/KinoTicketSystem/model"
	"github.com/ELITE-Kinoticketsystem/Backend-KTS/src/managers"
	"github.com/ELITE-Kinoticketsystem/Backend-KTS/src/models"
	"github.com/ELITE-Kinoticketsystem/Backend-KTS/src/myid"
	"github.com/ELITE-Kinoticketsystem/Backend-KTS/src/samples"
	"github.com/ELITE-Kinoticketsystem/Backend-KTS/src/utils"
)

func TestGetActorById(t *testing.T) {

	query := "\nSELECT .* FROM `KinoTicketSystem`.actors .*"

	actor := *samples.GetSampleActor()

	id := myid.MustParse(actor.ID.String())

	testCases := []struct {
		name            string
		setExpectations func(mock sqlmock.Sqlmock)
		expectedActor   *models.ActorDTO
		expectedError   *models.KTSError
	}{
		{
			name: "Select actor by id",
			setExpectations: func(mock sqlmock.Sqlmock) {
				mock.ExpectQuery(query).WithArgs(utils.EqUUID(id)).
					WillReturnRows(sqlmock.NewRows([]string{"actors.id", "actors.name", "actors.birthdate", "actors.description", "actor_pictures.id", "actor_pictures.actor_id", "actor_pictures.pic_url", "movies.id", "movies.title", "movies.description", "movies.banner_pic_url", "movies.cover_pic_url", "movies.trailer_url", "movies.rating", "movies.release_date", "movies.time_in_min", "movies.fsk"}).
						AddRow(actor.ID, actor.Name, actor.Birthdate, actor.Description, actor.Pictures[0].ID, actor.Pictures[0].ActorID, actor.Pictures[0].PicURL, actor.Movies[0].ID, actor.Movies[0].Title, actor.Movies[0].Description, actor.Movies[0].BannerPicURL, actor.Movies[0].CoverPicURL, actor.Movies[0].TrailerURL, actor.Movies[0].Rating, actor.Movies[0].ReleaseDate, actor.Movies[0].TimeInMin, actor.Movies[0].Fsk).
						AddRow(actor.ID, actor.Name, actor.Birthdate, actor.Description, actor.Pictures[1].ID, actor.Pictures[1].ActorID, actor.Pictures[1].PicURL, actor.Movies[0].ID, actor.Movies[0].Title, actor.Movies[0].Description, actor.Movies[0].BannerPicURL, actor.Movies[0].CoverPicURL, actor.Movies[0].TrailerURL, actor.Movies[0].Rating, actor.Movies[0].ReleaseDate, actor.Movies[0].TimeInMin, actor.Movies[0].Fsk).
						AddRow(actor.ID, actor.Name, actor.Birthdate, actor.Description, actor.Pictures[0].ID, actor.Pictures[0].ActorID, actor.Pictures[0].PicURL, actor.Movies[1].ID, actor.Movies[1].Title, actor.Movies[1].Description, actor.Movies[1].BannerPicURL, actor.Movies[1].CoverPicURL, actor.Movies[1].TrailerURL, actor.Movies[1].Rating, actor.Movies[1].ReleaseDate, actor.Movies[1].TimeInMin, actor.Movies[1].Fsk).
						AddRow(actor.ID, actor.Name, actor.Birthdate, actor.Description, actor.Pictures[1].ID, actor.Pictures[1].ActorID, actor.Pictures[1].PicURL, actor.Movies[1].ID, actor.Movies[1].Title, actor.Movies[1].Description, actor.Movies[1].BannerPicURL, actor.Movies[1].CoverPicURL, actor.Movies[1].TrailerURL, actor.Movies[1].Rating, actor.Movies[1].ReleaseDate, actor.Movies[1].TimeInMin, actor.Movies[1].Fsk),
					)

			},
			expectedActor: &actor,
			expectedError: nil,
		},
		{
			name: "Select actor by id - no actor found",
			setExpectations: func(mock sqlmock.Sqlmock) {
				mock.ExpectQuery(query).WithArgs(utils.EqUUID(id)).
					WillReturnRows(sqlmock.NewRows([]string{"actors.id", "actors.name", "actors.birthdate", "actors.description", "actor_pictures.id", "actor_pictures.actor_id", "actor_pictures.pic_url", "movies.id", "movies.title", "movies.description", "movies.banner_pic_url", "movies.cover_pic_url", "movies.trailer_url", "movies.rating", "movies.release_date", "movies.time_in_min", "movies.fsk"}))
			},
			expectedActor: nil,
			expectedError: kts_errors.KTS_NOT_FOUND,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			db, mock, err := sqlmock.New()
			if err != nil {
				t.Fatalf("Failed to create mock database connection: %v", err)
			}
			defer db.Close()

			// Create a new instance of the MovieRepository with the mock database connection
			actorRepo := &ActorRepository{
				DatabaseManager: &managers.DatabaseManager{
					Connection: db,
				},
			}

			tc.setExpectations(mock)

			// Call the method on the repository instance
			actor, kts_err := actorRepo.GetActorById(&id)

			// Verify that all expectations were met

			if err := mock.ExpectationsWereMet(); err != nil {
				t.Errorf("There were unfulfilled expectations: %s", err)
			}

			if kts_err != tc.expectedError {
				t.Errorf("Unexpected error: %v", kts_err)
			}

			assert.Equal(t, tc.expectedActor, actor)

		})
	}
}

func TestGetActors(t *testing.T) {
	actors := *samples.GetSampleActors()

	query := "SELECT .* FROM `KinoTicketSystem`.actors .*"

	testCases := []struct {
		name            string
		setExpectations func(mock sqlmock.Sqlmock)
		expectedActors  *[]models.GetActorsDTO
		expectedError   *models.KTSError
	}{
		{
			name: "Select all actors",
			setExpectations: func(mock sqlmock.Sqlmock) {
				mock.ExpectQuery(query).
					WillReturnRows(sqlmock.NewRows([]string{"actors.id", "actors.name", "actors.birthdate", "actors.description", "actor_pictures.id", "actor_pictures.actor_id", "actor_pictures.pic_url"}).
						AddRow(actors[0].ID, actors[0].Name, actors[0].Birthdate, actors[0].Description, actors[0].Pictures[0].ID, actors[0].Pictures[0].ActorID, actors[0].Pictures[0].PicURL).
						AddRow(actors[1].ID, actors[1].Name, actors[1].Birthdate, actors[1].Description, actors[1].Pictures[0].ID, actors[1].Pictures[0].ActorID, actors[1].Pictures[0].PicURL),
					)

			},
			expectedActors: &actors,
			expectedError:  nil,
		},
		{
			name: "Select all actors - no actors found",
			setExpectations: func(mock sqlmock.Sqlmock) {
				mock.ExpectQuery(query).
					WillReturnRows(sqlmock.NewRows([]string{"actors.id", "actors.name", "actors.birthdate", "actors.description", "actor_pictures.id", "actor_pictures.actor_id", "actor_pictures.pic_url"}))
			},
			expectedActors: nil,
			expectedError:  kts_errors.KTS_NOT_FOUND,
		},
		{
			name: "Select all actors - internal error",
			setExpectations: func(mock sqlmock.Sqlmock) {
				mock.ExpectQuery(query).
					WillReturnError(sql.ErrConnDone)
			},
			expectedActors: nil,
			expectedError:  kts_errors.KTS_INTERNAL_ERROR,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			db, mock, err := sqlmock.New()
			if err != nil {
				t.Fatalf("Failed to create mock database connection: %v", err)
			}
			defer db.Close()

			// Create a new instance of the MovieRepository with the mock database connection
			actorRepo := &ActorRepository{
				DatabaseManager: &managers.DatabaseManager{
					Connection: db,
				},
			}

			tc.setExpectations(mock)

			// Call the method on the repository instance
			actors, kts_err := actorRepo.GetActors()

			// Verify that all expectations were met

			if err := mock.ExpectationsWereMet(); err != nil {
				t.Errorf("There were unfulfilled expectations: %s", err)
			}

			if kts_err != tc.expectedError {
				t.Errorf("Unexpected error: %v", kts_err)
			}

			assert.Equal(t, tc.expectedActors, actors)

		})
	}

}

func TestCreateActor(t *testing.T) {

	actor := &model.Actors{

		Name:        "John Doe",
		Description: "Test actor",
		Birthdate:   time.Now(),
	}

	teststCases := []struct {
		name            string
		setExpectations func(mock sqlmock.Sqlmock)
		expectActorId   bool
		expectedError   *models.KTSError
	}{
		{
			name: "Create actor",
			setExpectations: func(mock sqlmock.Sqlmock) {
				mock.ExpectExec("INSERT INTO `KinoTicketSystem`.actors .*").WithArgs(sqlmock.AnyArg(), actor.Name, actor.Birthdate, actor.Description).WillReturnResult(sqlmock.NewResult(1, 1))
			},
			expectActorId: true,
			expectedError: nil,
		},
		{
			name: "Create actor sql error",
			setExpectations: func(mock sqlmock.Sqlmock) {
				mock.ExpectExec("INSERT INTO `KinoTicketSystem`.actors .*").WithArgs(sqlmock.AnyArg(), actor.Name, actor.Birthdate, actor.Description).WillReturnError(sql.ErrConnDone)
			},
			expectActorId: false,
			expectedError: kts_errors.KTS_INTERNAL_ERROR,
		},
		{
			name: "Create actor no rows affected",
			setExpectations: func(mock sqlmock.Sqlmock) {
				mock.ExpectExec("INSERT INTO `KinoTicketSystem`.actors .*").WithArgs(sqlmock.AnyArg(), actor.Name, actor.Birthdate, actor.Description).WillReturnResult(sqlmock.NewResult(1, 0))
			},
			expectActorId: false,
			expectedError: kts_errors.KTS_INTERNAL_ERROR,
		},
	}

	for _, tc := range teststCases {
		t.Run(tc.name, func(t *testing.T) {

			db, mock, err := sqlmock.New()
			if err != nil {
				t.Fatalf("Failed to create mock database connection: %v", err)
			}
			defer db.Close()

			actorRepo := &ActorRepository{
				DatabaseManager: &managers.DatabaseManager{
					Connection: db,
				},
			}

			tc.setExpectations(mock)

			id, kts_err := actorRepo.CreateActor(actor)

			if kts_err != tc.expectedError {
				t.Errorf("Unexpected error: %v", kts_err)
			}

			if tc.expectActorId && id == nil {
				t.Error("Expected actor ID, got nil")
			}

			if err := mock.ExpectationsWereMet(); err != nil {
				t.Errorf("There were unfulfilled expectations: %s", err)
			}

		})
	}
}

func TestCreateActorPicture(t *testing.T) {

	picUrl := "pic.jpg"

	ActorPicture := &model.ActorPictures{
		ActorID: myid.New(),
		PicURL:  &picUrl,
	}

	testCases := []struct {
		name                 string
		setExpectations      func(mock sqlmock.Sqlmock)
		expectActorPictureId bool
		expectedError        *models.KTSError
	}{
		{
			name: "Create actor picture",
			setExpectations: func(mock sqlmock.Sqlmock) {
				mock.ExpectExec("INSERT INTO `KinoTicketSystem`.actor_pictures .*").WithArgs(sqlmock.AnyArg(), sqlmock.AnyArg(), picUrl).WillReturnResult(sqlmock.NewResult(1, 1))
			},
			expectActorPictureId: true,
			expectedError:        nil,
		},
		{
			name: "Create actor picture sql error",
			setExpectations: func(mock sqlmock.Sqlmock) {
				mock.ExpectExec("INSERT INTO `KinoTicketSystem`.actor_pictures .*").WithArgs(sqlmock.AnyArg(), sqlmock.AnyArg(), picUrl).WillReturnError(sql.ErrConnDone)
			},
			expectActorPictureId: false,
			expectedError:        kts_errors.KTS_INTERNAL_ERROR,
		},
		{
			name: "Create actor picture no rows affected",
			setExpectations: func(mock sqlmock.Sqlmock) {
				mock.ExpectExec("INSERT INTO `KinoTicketSystem`.actor_pictures .*").WithArgs(sqlmock.AnyArg(), sqlmock.AnyArg(), picUrl).WillReturnResult(sqlmock.NewResult(1, 0))
			},
			expectActorPictureId: false,
			expectedError:        kts_errors.KTS_INTERNAL_ERROR,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherRegexp))
			if err != nil {
				t.Fatalf("Failed to create mock database connection: %v", err)
			}
			defer db.Close()

			actorRepo := &ActorRepository{
				DatabaseManager: &managers.DatabaseManager{
					Connection: db,
				},
			}

			tc.setExpectations(mock)

			id, kts_err := actorRepo.CreateActorPicture(ActorPicture)

			if kts_err != tc.expectedError {
				t.Errorf("Unexpected error: %v", kts_err)
			}

			if tc.expectActorPictureId && id == nil {
				t.Error("Expected actor picture ID, got nil")
			}

			if err := mock.ExpectationsWereMet(); err != nil {
				t.Errorf("There were unfulfilled expectations: %s", err)
			}

		})
	}
}
