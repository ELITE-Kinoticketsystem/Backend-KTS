package repositories

import (
	"database/sql"
	"log"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/ELITE-Kinoticketsystem/Backend-KTS/src/.gen/KinoTicketSystem/model"
	"github.com/ELITE-Kinoticketsystem/Backend-KTS/src/managers"
	"github.com/ELITE-Kinoticketsystem/Backend-KTS/src/models"
	"github.com/google/uuid"
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
)

func uuidPtr(id uuid.UUID) *uuid.UUID {
	return &id
}

func GetSampleMovies() *[]model.Movies {
	modelMovies := []model.Movies{}
	uuid1 := uuid.New()
	uuid2 := uuid.New()
	banner := ""
	cover := ""
	trailer := ""
	rating := 5.0

	modelMovies = append(modelMovies, model.Movies{
		ID:           uuid1,
		Title:        "Test Movie 1",
		Description:  "Test Description 1",
		BannerPicURL: &banner,
		CoverPicURL:  &cover,
		TrailerURL:   &trailer,
		Rating:       &rating,
		ReleaseDate:  time.Now(),
		TimeInMin:    120,
		Fsk:          18,
	})

	modelMovies = append(modelMovies, model.Movies{
		ID:           uuid2,
		Title:        "Test Movie 2",
		Description:  "Test Description 2",
		BannerPicURL: &banner,
		CoverPicURL:  &cover,
		TrailerURL:   &trailer,
		Rating:       &rating,
		ReleaseDate:  time.Now(),
		TimeInMin:    120,
		Fsk:          18,
	})

	return &modelMovies
}
// func TestGetMovies2(t *testing.T) {
// 	// (*sampleMovies) := utils.get(*sampleMovies)()
// 	sampleMovies := GetSampleMovies()

// 	testCases := []struct {
// 		name            string
// 		query           string
// 		setExpectations func(mock sqlmock.Sqlmock)
// 		expectedMovies  *[]model.Movies
// 		expectedError   *models.KTSError
// 	}{
// 		// {
// 		// 	name: "Empty result",
// 		// 	setExpectations: func(mock sqlmock.Sqlmock) {
// 		// 		mock.ExpectQuery(
// 		// 			`SELECT movies.id AS "movies.id",
// 		// 			movies.title AS "movies.title",
// 		// 			movies.description AS "movies.description",
// 		// 			movies.banner_pic_url AS "movies.banner_pic_url",
// 		// 			movies.cover_pic_url AS "movies.cover_pic_url",
// 		// 			movies.trailer_url AS "movies.trailer_url",
// 		// 			movies.rating AS "movies.rating",
// 		// 			movies.release_date AS "movies.release_date",
// 		// 			movies.time_in_min AS "movies.time_in_min",
// 		// 			movies.fsk AS "movies.fsk_"
// 		// 	   FROM "KinoTicketSystem".movies;"`,
// 		// 		).WillReturnRows(
// 		// 			sqlmock.NewRows([]string{"id", "title", "description", "release_date", "time_in_min", "fsk"}),
// 		// 		)
// 		// 	},
// 		// 	expectedMovies: nil,
// 		// 	expectedError:  kts_errors.KTS_INTERNAL_ERROR,
// 		// },
// 		{
// 			name: "Multiple movies",
// 			setExpectations: func(mock sqlmock.Sqlmock) {
// 				mock.ExpectQuery("\nSELECT movies.id AS \"movies.id\",\n     movies.title AS \"movies.title\",\n     movies.description AS \"movies.description\",\n     movies.banner_pic_url AS \"movies.banner_pic_url\",\n     movies.cover_pic_url AS \"movies.cover_pic_url\",\n     movies.trailer_url AS \"movies.trailer_url\",\n     movies.rating AS \"movies.rating\",\n     movies.release_date AS \"movies.release_date\",\n     movies.time_in_min AS \"movies.time_in_min\",\n     movies.fsk AS \"movies.fsk\"\nFROM `KinoTicketSystem`.movies;\n"). // "SELECT movies.id AS \"movies.id\"," +
// 					// "movies.title AS \"movies.title\"," +
// 					// "movies.description AS \"movies.description\"," +
// 					// "movies.banner_pic_url AS \"movies.banner_pic_url\"," +
// 					// "movies.cover_pic_url AS \"movies.cover_pic_url\"," +
// 					// "movies.trailer_url AS \"movies.trailer_url\"," +
// 					// "movies.rating AS \"movies.rating\"," +
// 					// "movies.release_date AS \"movies.release_date\"," +
// 					// "movies.time_in_min AS \"movies.time_in_min\"," +
// 					// "movies.fsk AS \"movies.fsk_\"" +
// 					// "FROM `KinoTicketSystem`.movies;",
// 					WillReturnRows(
// 						sqlmock.NewRows(
// 							[]string{"id", "title", "description", "banner_pic_url", "cover_pic_url", "trailer_url", "rating", "release_date", "time_in_min", "fsk"},
// 						).AddRow(
// 							(*sampleMovies)[0].ID, (*sampleMovies)[0].Title, (*sampleMovies)[0].Description, (*sampleMovies)[0].BannerPicURL, (*sampleMovies)[0].CoverPicURL, (*sampleMovies)[0].TrailerURL, (*sampleMovies)[0].Rating, (*sampleMovies)[0].ReleaseDate, (*sampleMovies)[0].TimeInMin, (*sampleMovies)[0].Fsk,
// 						).AddRow(
// 							(*sampleMovies)[1].ID, (*sampleMovies)[1].Title, (*sampleMovies)[1].Description, (*sampleMovies)[1].BannerPicURL, (*sampleMovies)[1].CoverPicURL, (*sampleMovies)[1].TrailerURL, (*sampleMovies)[1].Rating, (*sampleMovies)[1].ReleaseDate, (*sampleMovies)[1].TimeInMin, (*sampleMovies)[1].Fsk,
// 						),
// 					)
// 			},
// 			expectedMovies: &(*sampleMovies),
// 			expectedError:  nil,
// 		},
// 		// {
// 		// 	name:           "Error while querying movies",
// 		// 	query:          "SELECT * FROM movies;",
// 		// 	expectedMovies: nil,
// 		// 	expectedError:  kts_errors.KTS_NOT_FOUND,
// 		// },
// 		// {
// 		// 	name:           "Error while scanning movie",
// 		// 	query:          "SELECT * FROM movies;",
// 		// 	expectedMovies: nil,
// 		// 	expectedError:  kts_errors.KTS_INTERNAL_ERROR,
// 		// },
// 	}

// 	for _, tc := range testCases {
// 		t.Run(tc.name, func(t *testing.T) {
// 			// Create a new mock database connection
// 			db, mock, err := sqlmock.New()
// 			if err != nil {
// 				t.Fatalf("Failed to create mock database connection: %v", err)
// 			}
// 			defer db.Close()

// 			// Create a new instance of the MovieRepository with the mock database connection
// 			movieRepo := MovieRepository{
// 				DatabaseManager: &managers.DatabaseManager{
// 					Connection: db,
// 				},
// 			}

// 			tc.setExpectations(mock)

// 			// Call the method under test
// 			movies, kts_err := movieRepo.GetMovies()

// 			// Verify the results
// 			assert.Equal(t, tc.expectedMovies, movies)
// 			assert.Equal(t, tc.expectedError, kts_err)

// 			// Verify that all expectations were met
// 			if err = mock.ExpectationsWereMet(); err != nil {
// 				t.Errorf("There were unfulfilled expectations: %s", err)
// 			}

// 		})
// 	}
// }

// func TestGetMovies(t *testing.T) {
// 	t.Run("", func(t *testing.T) {
// 		movieRepo, dbConnection := GetDatabaseConnection2()
// 		defer dbConnection.Close()

// 		// Call the method under test
// 		movies, kts_err := movieRepo.GetMovies()

// 		// For Testing
// 		JsonSave("./Movies.json", movies)

// 		assert.Nil(t, kts_err)

// 	})
// }

// func GetDatabaseConnection2() (*MovieRepository, *sql.DB) {
// 	// Create a new mock database connection
// 	log.Println("Loading environment variables...")
// 	err := godotenv.Load()
// 	if err != nil {
// 		log.Printf("error loading .env file: %v\n", err)
// 	} else {
// 		log.Println("Environment variables loaded successfully")
// 	}

// 	log.Println("Initializing database connection...")
// 	dbConnection, err := managers.InitializeDB()
// 	if err != nil {
// 		panic(err)
// 	}
// 	log.Println("Database initialized successfully")

// 	// Create a new instance of the MovieRepository with the mock database connection
// 	movieRepo := MovieRepository{
// 		DatabaseManager: &managers.DatabaseManager{
// 			Connection: dbConnection,
// 		},
// 	}

// 	return &movieRepo, dbConnection
// }


func TestGetMovies3(t *testing.T) {
	// (*sampleMovies) := utils.get(*sampleMovies)()
	sampleMovies := GetSampleMovies()

	testCases := []struct {
		name            string
		query           string
		setExpectations func(mock sqlmock.Sqlmock)
		expectedMovies  *[]model.Movies
		expectedError   *models.KTSError
	}{
		{
			name: "Multiple movies",
			setExpectations: func(mock sqlmock.Sqlmock) {
				mock.ExpectQuery("\nSELECT movies.id AS \"movies.id\",\n     movies.title AS \"movies.title\",\n     movies.description AS \"movies.description\",\n     movies.banner_pic_url AS \"movies.banner_pic_url\",\n     movies.cover_pic_url AS \"movies.cover_pic_url\",\n     movies.trailer_url AS \"movies.trailer_url\",\n     movies.rating AS \"movies.rating\",\n     movies.release_date AS \"movies.release_date\",\n     movies.time_in_min AS \"movies.time_in_min\",\n     movies.fsk AS \"movies.fsk\"\nFROM `KinoTicketSystem`.movies;\n"). 
					WillReturnRows(
						sqlmock.NewRows(
							[]string{"id", "title", "description", "banner_pic_url", "cover_pic_url", "trailer_url", "rating", "release_date", "time_in_min", "fsk"},
						).AddRow(
							(*sampleMovies)[0].ID, (*sampleMovies)[0].Title, (*sampleMovies)[0].Description, (*sampleMovies)[0].BannerPicURL, (*sampleMovies)[0].CoverPicURL, (*sampleMovies)[0].TrailerURL, (*sampleMovies)[0].Rating, (*sampleMovies)[0].ReleaseDate, (*sampleMovies)[0].TimeInMin, (*sampleMovies)[0].Fsk,
						).AddRow(
							(*sampleMovies)[1].ID, (*sampleMovies)[1].Title, (*sampleMovies)[1].Description, (*sampleMovies)[1].BannerPicURL, (*sampleMovies)[1].CoverPicURL, (*sampleMovies)[1].TrailerURL, (*sampleMovies)[1].Rating, (*sampleMovies)[1].ReleaseDate, (*sampleMovies)[1].TimeInMin, (*sampleMovies)[1].Fsk,
						),
					)
			},
			expectedMovies: &(*sampleMovies),
			expectedError:  nil,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Create a new mock database connection
			db, mock, err := sqlmock.New()
			if err != nil {
				t.Fatalf("Failed to create mock database connection: %v", err)
			}
			defer db.Close()

			// Create a new instance of the MovieRepository with the mock database connection
			movieRepo := MovieRepository{
				DatabaseManager: &managers.DatabaseManager{
					Connection: db,
				},
			}

			tc.setExpectations(mock)

			// Call the method under test
			movies, kts_err := movieRepo.GetMovies()

			// Verify the results
			assert.Equal(t, tc.expectedMovies, movies)
			assert.Equal(t, tc.expectedError, kts_err)

			// Verify that all expectations were met
			if err = mock.ExpectationsWereMet(); err != nil {
				t.Errorf("There were unfulfilled expectations: %s", err)
			}

		})
	}
}

func TestGetMovies(t *testing.T) {
	t.Run("", func(t *testing.T) {
		movieRepo, dbConnection := GetDatabaseConnection2()
		defer dbConnection.Close()

		// Call the method under test
		movies, kts_err := movieRepo.GetMovies()

		// For Testing
		JsonSave("./Movies.json", movies)

		assert.Nil(t, kts_err)

	})
}

func GetDatabaseConnection2() (*MovieRepository, *sql.DB) {
	// Create a new mock database connection
	log.Println("Loading environment variables...")
	err := godotenv.Load()
	if err != nil {
		log.Printf("error loading .env file: %v\n", err)
	} else {
		log.Println("Environment variables loaded successfully")
	}

	log.Println("Initializing database connection...")
	dbConnection, err := managers.InitializeDB()
	if err != nil {
		panic(err)
	}
	log.Println("Database initialized successfully")

	// Create a new instance of the MovieRepository with the mock database connection
	movieRepo := MovieRepository{
		DatabaseManager: &managers.DatabaseManager{
			Connection: dbConnection,
		},
	}

	return &movieRepo, dbConnection
}
