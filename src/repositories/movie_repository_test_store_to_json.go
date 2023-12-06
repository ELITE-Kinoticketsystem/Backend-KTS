package repositories

import (
	"database/sql"
	"encoding/json"
	"log"
	"os"
	"testing"
	"time"

	"github.com/ELITE-Kinoticketsystem/Backend-KTS/src/.gen/KinoTicketSystem/model"
	"github.com/ELITE-Kinoticketsystem/Backend-KTS/src/managers"

	"github.com/google/uuid"
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
)

func JsonSave(path string, v interface{}) {
	jsonText, _ := json.MarshalIndent(v, "", "\t")

	// err := ioutil.WriteFile(path, jsonText, 0644)
	err := os.WriteFile(path, jsonText, 0644)

	if err != nil {
		panic(err)
	}
}

func PrintJson(v interface{}) {
	jsonText, _ := json.MarshalIndent(v, "", "\t")
	println(string(jsonText))
}

func GetDatabaseConnection() (*MovieRepository, *sql.DB) {
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

func TestGetMoviesExecutable(t *testing.T) {
	t.Run("", func(t *testing.T) {
		movieRepo, dbConnection := GetDatabaseConnection()
		defer dbConnection.Close()

		// Call the method under test
		movies, kts_err := movieRepo.GetMovies()

		// For Testing
		JsonSave("./Movies.json", movies)

		assert.Nil(t, kts_err)

	})
}

func TestGetMovieByIdExecutable(t *testing.T) {
	t.Run("", func(t *testing.T) {
		movieRepo, dbConnection := GetDatabaseConnection()
		defer dbConnection.Close()

		id := uuid.MustParse("6ba7b826-9dad-11d1-80b4-00c04fd430c0")

		movies, kts_err := movieRepo.GetMovieById(&id)

		// For Testing
		JsonSave("./MovieDetail.json", movies)

		assert.Nil(t, kts_err)

	})
}

func TestCreateMovieExecutable(t *testing.T) {
	var movie model.Movies

	var banner *string = new(string)
	var cover *string = new(string)
	var trailer *string = new(string)
	var rating *float64 = new(float64)

	*banner = "MyOwnBannerPicURL"
	*cover = "MyOwnCoverPicURL"
	*trailer = "MyOwnTrailerURL"
	*rating = 5

	movie.Title = "MyOwnMovie"
	movie.Description = "MyOwnDescription"
	movie.BannerPicURL = banner
	movie.CoverPicURL = cover
	movie.TrailerURL = trailer
	movie.Rating = rating
	movie.ReleaseDate = time.Now()
	movie.TimeInMin = 120
	movie.Fsk = 18

	t.Run("", func(t *testing.T) {
		movieRepo, dbConnection := GetDatabaseConnection()
		defer dbConnection.Close()

		// Call the method under test
		kts_err := movieRepo.CreateMovie(movie)

		assert.Nil(t, kts_err)

	})
}

func TestUpdateMovieExecutable(t *testing.T) {
	var movie model.Movies

	id := uuid.MustParse("6ba7b826-9dad-11d1-80b4-00c04fd430c0")
	banner := "UpdatedBannerPicURL"
	cover := "UpdatedCoverPicURL"
	trailer := "UpdatedTrailerURL"
	rating := 4.5

	movie.ID = &id
	movie.Title = "UpdatedMovie"
	movie.Description = "UpdatedDescription"
	movie.BannerPicURL = &banner
	movie.CoverPicURL = &cover
	movie.TrailerURL = &trailer
	movie.Rating = &rating
	movie.ReleaseDate = time.Now()
	movie.TimeInMin = 150
	movie.Fsk = 16

	t.Run("", func(t *testing.T) {
		movieRepo, dbConnection := GetDatabaseConnection()
		defer dbConnection.Close()

		// Call the method under test
		kts_err := movieRepo.UpdateMovie(movie)

		assert.Nil(t, kts_err)

	})
}

func TestDeleteMovieExecutable(t *testing.T) {
	t.Run("", func(t *testing.T) {
		movieRepo, dbConnection := GetDatabaseConnection()
		defer dbConnection.Close()

		id := uuid.MustParse("6ba7b826-9dad-11d1-80b4-00c04fd430c0")

		// Call the method under test
		kts_err := movieRepo.DeleteMovie(&id)

		assert.Nil(t, kts_err)
	})
}

// Genre
func TestGetGenresExecutable(t *testing.T) {
	t.Run("", func(t *testing.T) {
		movieRepo, dbConnection := GetDatabaseConnection()
		defer dbConnection.Close()

		// Call the method under test
		genres, kts_err := movieRepo.GetGenres()

		// For Testing
		JsonSave("./Genres.json", genres)

		assert.Nil(t, kts_err)

	})
}

func TestGetGenreByNameExecutable(t *testing.T) {
	t.Run("", func(t *testing.T) {
		movieRepo, dbConnection := GetDatabaseConnection()
		defer dbConnection.Close()

		genreName := "Action"

		// Call the method under test
		genres, kts_err := movieRepo.GetGenreByName(genreName)

		// For Testing
		JsonSave("./GenreDetails.json", genres)

		assert.Nil(t, kts_err)

	})
}

func TestCreateGenreExecutable(t *testing.T) {
	t.Run("", func(t *testing.T) {
		movieRepo, dbConnection := GetDatabaseConnection()
		defer dbConnection.Close()

		genre := "MyOwnGenre2"

		// Call the method under test
		kts_err := movieRepo.CreateGenre(genre)

		assert.Nil(t, kts_err)

	})
}

func TestAddMovieGenreExecutable(t *testing.T) {
	t.Run("", func(t *testing.T) {
		movieRepo, dbConnection := GetDatabaseConnection()
		defer dbConnection.Close()

		movieId := uuid.MustParse("11ee9147-e9e0-f488-bf69-0242ac120003")
		genreId := uuid.MustParse("11ee913a-2c00-05f1-bf69-0242ac120003")

		// Call the method under test
		kts_err := movieRepo.AddMovieGenre(&movieId, &genreId)

		assert.Nil(t, kts_err)
	})
}

// Combine Movie and Genre
func TestGetMovieByIdWithGenreExecutable(t *testing.T) {
	t.Run("", func(t *testing.T) {
		movieRepo, dbConnection := GetDatabaseConnection()
		defer dbConnection.Close()

		movieId := uuid.MustParse("6ba7b828-9dad-11d1-80b4-00c04fd430c2")

		// Call the method under test
		movieWithGenres, kts_err := movieRepo.GetMovieByIdWithGenre(&movieId)

		// For Testing
		JsonSave("./MovieByIdWithGenres.json", movieWithGenres)

		assert.Nil(t, kts_err)

	})
}

// others
func TestGetGenreWithMovies(t *testing.T) {
	t.Run("", func(t *testing.T) {
		movieRepo, dbConnection := GetDatabaseConnection()
		defer dbConnection.Close()

		genre := "Action"

		// Call the method under test
		moviesByGenre, kts_err := movieRepo.GetGenreByNameWithMovies(genre)

		// For Testing
		JsonSave("./GenreWithMovies.json", moviesByGenre)

		assert.Nil(t, kts_err)

	})
}

func TestGetGenresWithMoviesExecutable(t *testing.T) {
	t.Run("", func(t *testing.T) {
		movieRepo, dbConnection := GetDatabaseConnection()
		defer dbConnection.Close()

		// Call the method under test
		genresWithMovies, kts_err := movieRepo.GetGenresWithMovies()

		// For Testing
		JsonSave("./GenresWithMovies.json", genresWithMovies)

		assert.Nil(t, kts_err)

	})
}

func TestGetMoviesWithGenresExecutable(t *testing.T) {
	t.Run("", func(t *testing.T) {
		movieRepo, dbConnection := GetDatabaseConnection()
		defer dbConnection.Close()

		// Call the method under test
		moviesWithGenres, kts_err := movieRepo.GetMoviesWithGenres()

		// For Testing
		JsonSave("./MoviesWithGenres.json", moviesWithGenres)

		assert.Nil(t, kts_err)

	})
}
