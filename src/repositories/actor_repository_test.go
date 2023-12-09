package repositories

import (
	"database/sql"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"

	"github.com/ELITE-Kinoticketsystem/Backend-KTS/src/.gen/KinoTicketSystem/model"
	kts_errors "github.com/ELITE-Kinoticketsystem/Backend-KTS/src/errors"
	"github.com/ELITE-Kinoticketsystem/Backend-KTS/src/managers"
	"github.com/ELITE-Kinoticketsystem/Backend-KTS/src/models"
	"github.com/ELITE-Kinoticketsystem/Backend-KTS/src/utils"
)

func TestGetActorById(t *testing.T) {

	query := "\nSELECT actors.id AS \"actors.id\",\n     actors.name AS \"actors.name\",\n     actors.birthdate AS \"actors.birthdate\",\n     actors.description AS \"actors.description\",\n     actors.pic_url AS \"actors.pic_url\",\n     actor_pictures.id AS \"actor_pictures.id\",\n     actor_pictures.actor_id AS \"actor_pictures.actor_id\",\n     actor_pictures.pic_url AS \"actor_pictures.pic_url\",\n     movies.id AS \"movies.id\",\n     movies.title AS \"movies.title\",\n     movies.description AS \"movies.description\",\n     movies.banner_pic_url AS \"movies.banner_pic_url\",\n     movies.cover_pic_url AS \"movies.cover_pic_url\",\n     movies.trailer_url AS \"movies.trailer_url\",\n     movies.rating AS \"movies.rating\",\n     movies.release_date AS \"movies.release_date\",\n     movies.time_in_min AS \"movies.time_in_min\",\n     movies.fsk AS \"movies.fsk\"\nFROM `KinoTicketSystem`.actors\n     LEFT JOIN `KinoTicketSystem`.actor_pictures ON (actor_pictures.actor_id = actors.id)\n     LEFT JOIN `KinoTicketSystem`.movie_actors ON (movie_actors.actor_id = actors.id)\n     LEFT JOIN `KinoTicketSystem`.movies ON (movies.id = movie_actors.movie_id)\nWHERE actors.id = ?;\n"

	actor := *GetActor()

	id := uuid.MustParse(actor.ID.String())

	testCases := []struct {
		name            string
		setExpectations func(mock sqlmock.Sqlmock)
		expectedActor   *models.ActorDTO
		expectedError   *models.KTSError
	}{
		{
			name: "Select actor by id",
			setExpectations: func(mock sqlmock.Sqlmock) {
				mock.ExpectQuery(query).WithArgs(utils.EqUUID(&id)).
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
				mock.ExpectQuery(query).WithArgs(utils.EqUUID(&id)).
					WillReturnRows(sqlmock.NewRows([]string{"actors.id", "actors.name", "actors.birthdate", "actors.description", "actor_pictures.id", "actor_pictures.actor_id", "actor_pictures.pic_url", "movies.id", "movies.title", "movies.description", "movies.banner_pic_url", "movies.cover_pic_url", "movies.trailer_url", "movies.rating", "movies.release_date", "movies.time_in_min", "movies.fsk"}))
			},
			expectedActor: nil,
			expectedError: kts_errors.KTS_NOT_FOUND,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
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
	actors := *GetActors()

	query := "\nSELECT actors.id AS \"actors.id\",\n     actors.name AS \"actors.name\",\n     actors.birthdate AS \"actors.birthdate\",\n     actors.description AS \"actors.description\",\n     actors.pic_url AS \"actors.pic_url\",\n     actor_pictures.id AS \"actor_pictures.id\",\n     actor_pictures.actor_id AS \"actor_pictures.actor_id\",\n     actor_pictures.pic_url AS \"actor_pictures.pic_url\"\nFROM `KinoTicketSystem`.actors\n     LEFT JOIN `KinoTicketSystem`.actor_pictures ON (actor_pictures.actor_id = actors.id);\n"

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

			db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
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

func GetActor() *models.ActorDTO {

	actorId := uuid.New()

	url1 := "https://de.wikipedia.org/wiki/Brad_Pitt#/media/Datei:SevenYearsInTibeta.jpg"
	url2 := "https://en.wikipedia.org/wiki/Brad_Pitt_filmography#/media/File:Brad_Pitt_Fury_2014.jpg"

	picId1 := uuid.New()
	picId2 := uuid.New()

	picture1 := model.ActorPictures{
		ID:      &picId1,
		ActorID: &actorId,
		PicURL:  &url1,
	}

	picture2 := model.ActorPictures{
		ID:      &picId2,
		ActorID: &actorId,
		PicURL:  &url2,
	}

	releaseDate1 := time.Date(1972, 3, 24, 0, 0, 0, 0, time.UTC)
	releaseDate2 := time.Date(1999, 10, 15, 0, 0, 0, 0, time.UTC)

	rating := 0.

	movieId1 := uuid.New()
	movieId2 := uuid.New()

	movie1 := model.Movies{
		ID:           &movieId1,
		Title:        "The Godfather",
		Description:  "The aging patriarch of an organized crime dynasty transfers control of his clandestine empire to his reluctant son.",
		BannerPicURL: nil,
		CoverPicURL:  nil,
		TrailerURL:   nil,
		Rating:       &rating,
		ReleaseDate:  releaseDate1,
		TimeInMin:    0,
		Fsk:          0,
	}

	movie2 := model.Movies{
		ID:           &movieId2,
		Title:        "Fight Club",
		Description:  "An insomniac office worker and a devil-may-care soapmaker form an underground fight club that evolves into something much, much more.",
		BannerPicURL: nil,
		CoverPicURL:  nil,
		TrailerURL:   nil,
		Rating:       &rating,
		ReleaseDate:  releaseDate2,
		TimeInMin:    0,
		Fsk:          0,
	}

	actor := models.ActorDTO{
		Actors: model.Actors{
			ID:          &actorId,
			Name:        "Brad Pitt",
			Description: "Brad Pitt is an actor.",
			Birthdate:   time.Date(1963, 12, 18, 0, 0, 0, 0, time.UTC),
		},
		Pictures: []model.ActorPictures{
			picture1,
			picture2,
		},
		Movies: []model.Movies{
			movie1,
			movie2,
		},
	}

	return &actor
}

func GetActors() *[]models.GetActorsDTO {

	actor1Id := uuid.New()

	url := "BradPitt.jpg"

	pic1Id := uuid.New()

	actor1Picture := model.ActorPictures{
		ID:      &pic1Id,
		ActorID: &actor1Id,
		PicURL:  &url,
	}

	actor1 := models.GetActorsDTO{
		Actors: model.Actors{
			ID:          &actor1Id,
			Name:        "Brad Pitt",
			Description: "Brad Pitt is an actor.",
			Birthdate:   time.Date(1963, 12, 18, 0, 0, 0, 0, time.UTC),
		},
		Pictures: []model.ActorPictures{
			actor1Picture,
		},
	}

	actor2Id := uuid.New()

	url2 := "EdwardNorton.jpg"

	pic2Id := uuid.New()

	actor2Picture := model.ActorPictures{
		ID:      &pic2Id,
		ActorID: &actor2Id,
		PicURL:  &url2,
	}

	actor2 := models.GetActorsDTO{
		Actors: model.Actors{
			ID:          &actor2Id,
			Name:        "Edward Norton",
			Description: "Edward Norton is an actor.",
			Birthdate:   time.Date(1969, 8, 18, 0, 0, 0, 0, time.UTC),
		},
		Pictures: []model.ActorPictures{
			actor2Picture,
		},
	}

	return &[]models.GetActorsDTO{
		actor1,
		actor2,
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
				mock.ExpectExec("INSERT INTO `KinoTicketSystem`.actors .*").WillReturnResult(sqlmock.NewResult(1, 1))
			},
			expectActorId: true,
			expectedError: nil,
		},
		{
			name: "Create actor sql error",
			setExpectations: func(mock sqlmock.Sqlmock) {
				mock.ExpectExec("INSERT INTO `KinoTicketSystem`.actors .*").WillReturnError(sql.ErrConnDone)
			},
			expectActorId: false,
			expectedError: kts_errors.KTS_INTERNAL_ERROR,
		},
		{
			name: "Create actor no rows affected",
			setExpectations: func(mock sqlmock.Sqlmock) {
				mock.ExpectExec("INSERT INTO `KinoTicketSystem`.actors .*").WillReturnResult(sqlmock.NewResult(1, 0))
			},
			expectActorId: false,
			expectedError: kts_errors.KTS_INTERNAL_ERROR,
		},
	}

	for _, tc := range teststCases {
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
		ActorID: utils.NewUUID(),
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
				mock.ExpectExec("INSERT INTO `KinoTicketSystem`.actor_pictures .*").WillReturnResult(sqlmock.NewResult(1, 1))
			},
			expectActorPictureId: true,
			expectedError:        nil,
		},
		{
			name: "Create actor picture sql error",
			setExpectations: func(mock sqlmock.Sqlmock) {
				mock.ExpectExec("INSERT INTO `KinoTicketSystem`.actor_pictures .*").WillReturnError(sql.ErrConnDone)
			},
			expectActorPictureId: false,
			expectedError:        kts_errors.KTS_INTERNAL_ERROR,
		},
		{
			name: "Create actor picture no rows affected",
			setExpectations: func(mock sqlmock.Sqlmock) {
				mock.ExpectExec("INSERT INTO `KinoTicketSystem`.actor_pictures .*").WillReturnResult(sqlmock.NewResult(1, 0))
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
