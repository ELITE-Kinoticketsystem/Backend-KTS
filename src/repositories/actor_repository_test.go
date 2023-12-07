package repositories

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"

	"github.com/ELITE-Kinoticketsystem/Backend-KTS/src/.gen/KinoTicketSystem/model"
	"github.com/ELITE-Kinoticketsystem/Backend-KTS/src/managers"
	"github.com/ELITE-Kinoticketsystem/Backend-KTS/src/models"
	"github.com/ELITE-Kinoticketsystem/Backend-KTS/src/utils"
)

// actors.id,actors.name,actors.age,actors.description,movie_actors.movie_id,movies.id,movies.title,movies.description,movies.banner_pic_url,movies.cover_pic_url,movies.trailer_url,movies.rating,movies.release_date,movies.time_in_min,movies.fsk,actor_pictures.id,actor_pictures.actor_id,actor_pictures.pic_url
// 6BA7B8429DAD11D180B400C04FD430C2,Brad Pitt,57,Brad Pitt is an actor.,6BA7B8279DAD11D180B400C04FD430C1,6BA7B8279DAD11D180B400C04FD430C1,The Godfather,The aging patriarch of an organized crime dynasty transfers control of his clandestine empire to his reluctant son.,,,,,1972-03-24,175,16,11EE93695EC9113BA05C0242AC120003,6BA7B8429DAD11D180B400C04FD430C2,https://de.wikipedia.org/wiki/Brad_Pitt#/media/Datei:SevenYearsInTibeta.jpg
// 6BA7B8429DAD11D180B400C04FD430C2,Brad Pitt,57,Brad Pitt is an actor.,6BA7B82A9DAD11D180B400C04FD430C4,6BA7B82A9DAD11D180B400C04FD430C4,Fight Club,"An insomniac office worker and a devil-may-care soapmaker form an underground fight club that evolves into something much, much more.",,,,,1999-10-15,139,18,11EE93695EC9113BA05C0242AC120003,6BA7B8429DAD11D180B400C04FD430C2,https://de.wikipedia.org/wiki/Brad_Pitt#/media/Datei:SevenYearsInTibeta.jpg
// 6BA7B8429DAD11D180B400C04FD430C2,Brad Pitt,57,Brad Pitt is an actor.,6BA7B82D9DAD11D180B400C04FD430C7,6BA7B82D9DAD11D180B400C04FD430C7,Goodfellas,"The story of Henry Hill and his life in the mob, covering his relationship with his wife Karen Hill and his mob partners Jimmy Conway and Tommy DeVito in the Italian-American crime syndicate.",,,,,1990-09-19,146,16,11EE93695EC9113BA05C0242AC120003,6BA7B8429DAD11D180B400C04FD430C2,https://de.wikipedia.org/wiki/Brad_Pitt#/media/Datei:SevenYearsInTibeta.jpg
// 6BA7B8429DAD11D180B400C04FD430C2,Brad Pitt,57,Brad Pitt is an actor.,6BA7B8279DAD11D180B400C04FD430C1,6BA7B8279DAD11D180B400C04FD430C1,The Godfather,The aging patriarch of an organized crime dynasty transfers control of his clandestine empire to his reluctant son.,,,,,1972-03-24,175,16,11EE9368DC48C460A05C0242AC120003,6BA7B8429DAD11D180B400C04FD430C2,https://en.wikipedia.org/wiki/Brad_Pitt_filmography#/media/File:Brad_Pitt_Fury_2014.jpg
// 6BA7B8429DAD11D180B400C04FD430C2,Brad Pitt,57,Brad Pitt is an actor.,6BA7B82A9DAD11D180B400C04FD430C4,6BA7B82A9DAD11D180B400C04FD430C4,Fight Club,"An insomniac office worker and a devil-may-care soapmaker form an underground fight club that evolves into something much, much more.",,,,,1999-10-15,139,18,11EE9368DC48C460A05C0242AC120003,6BA7B8429DAD11D180B400C04FD430C2,https://en.wikipedia.org/wiki/Brad_Pitt_filmography#/media/File:Brad_Pitt_Fury_2014.jpg
// 6BA7B8429DAD11D180B400C04FD430C2,Brad Pitt,57,Brad Pitt is an actor.,6BA7B82D9DAD11D180B400C04FD430C7,6BA7B82D9DAD11D180B400C04FD430C7,Goodfellas,"The story of Henry Hill and his life in the mob, covering his relationship with his wife Karen Hill and his mob partners Jimmy Conway and Tommy DeVito in the Italian-American crime syndicate.",,,,,1990-09-19,146,16,11EE9368DC48C460A05C0242AC120003,6BA7B8429DAD11D180B400C04FD430C2,https://en.wikipedia.org/wiki/Brad_Pitt_filmography#/media/File:Brad_Pitt_Fury_2014.jpg

func GetActors() *[]models.ActorDTO {

	actor1 := models.ActorDTO{
		Actors: model.Actors{
			ID:          utils.UuidMustParse("6BA7B8429DAD11D180B400C04FD430C2"),
			Name:        "Brad Pitt",
			Description: "Brad Pitt is an actor.",
			Age:         57,
		},
	}

	url1 := "https://de.wikipedia.org/wiki/Brad_Pitt#/media/Datei:SevenYearsInTibeta.jpg"
	url2 := "https://en.wikipedia.org/wiki/Brad_Pitt_filmography#/media/File:Brad_Pitt_Fury_2014.jpg"

	picture1 := model.ActorPictures{
		ID:      utils.UuidMustParse("6BA7B8429DAD11D180B400C04FD430C2"),
		ActorID: utils.UuidMustParse("6BA7B8429DAD11D180B400C04FD430C2"),
		PicURL:  &url1,
	}

	picture2 := model.ActorPictures{
		ID:      utils.UuidMustParse("6BA7B8429DAD11D180B400C04FD430C2"),
		ActorID: utils.UuidMustParse("6BA7B8429DAD11D180B400C04FD430C2"),
		PicURL:  &url2,
	}

	actor1.Pictures = append(actor1.Pictures, struct {
		model.ActorPictures
	}{
		picture1,
	})

	actor1.Pictures = append(actor1.Pictures, struct {
		model.ActorPictures
	}{
		picture2,
	})

	releaseDate1 := time.Date(1972, 3, 24, 0, 0, 0, 0, time.UTC)
	releaseDate2 := time.Date(1999, 10, 15, 0, 0, 0, 0, time.UTC)

	rating := 0.

	movie1 := model.Movies{
		ID:           utils.UuidMustParse("6BA7B8279DAD11D180B400C04FD430C1"),
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
		ID:           utils.UuidMustParse("6BA7B82A9DAD11D180B400C04FD430C4"),
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

	actor1.Movies = append(actor1.Movies, struct {
		model.Movies
	}{
		movie1,
	})

	actor1.Movies = append(actor1.Movies, struct {
		model.Movies
	}{
		movie2,
	})

	return &[]models.ActorDTO{
		actor1,
	}
}

// mysql.SELECT(
// 	table.Actors.AllColumns,
// 	table.ActorPictures.AllColumns,
// 	table.Movies.AllColumns,
// ).
// 	FROM(
// 		table.Actors.
// 			LEFT_JOIN(table.ActorPictures, table.ActorPictures.ActorID.EQ(table.Actors.ID)).
// 			LEFT_JOIN(table.MovieActors, table.MovieActors.ActorID.EQ(table.Actors.ID)).
// 			LEFT_JOIN(table.Movies, table.Movies.ID.EQ(table.MovieActors.MovieID)),
// 	).
// 	WHERE(
// 		table.Actors.ID.EQ(mysql.String(string(binary_id))),
// 	)

func TestGetActorById(t *testing.T) {

	query := "SELECT actors.id AS \"actors.id\",\n     actors.name AS \"actors.name\",\n     actors.age AS \"actors.age\",\n     actors.description AS \"actors.description\",\n     actor_pictures.id AS \"actor_pictures.id\",\n     actor_pictures.actor_id AS \"actor_pictures.actor_id\",\n     actor_pictures.pic_url AS \"actor_pictures.pic_url\",\n     movies.id AS \"movies.id\",\n     movies.title AS \"movies.title\",\n     movies.description AS \"movies.description\",\n     movies.banner_pic_url AS \"movies.banner_pic_url\",\n     movies.cover_pic_url AS \"movies.cover_pic_url\",\n     movies.trailer_url AS \"movies.trailer_url\",\n     movies.rating AS \"movies.rating\",\n     movies.release_date AS \"movies.release_date\",\n     movies.time_in_min AS \"movies.time_in_min\",\n     movies.fsk AS \"movies.fsk\"\nFROM `KinoTicketSystem`.actors\n     LEFT JOIN `KinoTicketSystem`.actor_pictures ON (actor_pictures.actor_id = actors.id)\n     LEFT JOIN `KinoTicketSystem`.movie_actors ON (movie_actors.actor_id = actors.id)\n     LEFT JOIN `KinoTicketSystem`.movies ON (movies.id = movie_actors.movie_id)\nWHERE actors.id = ?;\n"

	actor := (*GetActors())[0]

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
					WillReturnRows(sqlmock.NewRows([]string{"actors.id", "actors.name", "actors.age", "actors.description", "actor_pictures.id", "actor_pictures.actor_id", "actor_pictures.pic_url", "movies.id", "movies.title", "movies.description", "movies.banner_pic_url", "movies.cover_pic_url", "movies.trailer_url", "movies.rating", "movies.release_date", "movies.time_in_min", "movies.fsk"}).
						AddRow(actor.ID, actor.Name, actor.Age, actor.Description, actor.Pictures[0].ID, actor.Pictures[0].ActorID, actor.Pictures[0].PicURL, actor.Movies[0].ID, actor.Movies[0].Title, actor.Movies[0].Description, actor.Movies[0].BannerPicURL, actor.Movies[0].CoverPicURL, actor.Movies[0].TrailerURL, actor.Movies[0].Rating, actor.Movies[0].ReleaseDate, actor.Movies[0].TimeInMin, actor.Movies[0].Fsk).
						AddRow(actor.ID, actor.Name, actor.Age, actor.Description, actor.Pictures[1].ID, actor.Pictures[1].ActorID, actor.Pictures[1].PicURL, actor.Movies[1].ID, actor.Movies[1].Title, actor.Movies[1].Description, actor.Movies[1].BannerPicURL, actor.Movies[1].CoverPicURL, actor.Movies[1].TrailerURL, actor.Movies[1].Rating, actor.Movies[1].ReleaseDate, actor.Movies[1].TimeInMin, actor.Movies[1].Fsk))

			},
			expectedActor: &actor,
			expectedError: nil,
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

			log.Println("Expected Actor: " + kts_err.ErrorMessage)

			assert.NotNil(t, actor)

			// Verify the results
			// assert.Equal(t, tc.expectedActor, actor)
			// assert.Equal(t, tc.expectedError, err)

			// Print the JSON of tc.expectedActor and actor to the console

			jsonActor, err := json.MarshalIndent(actor, "", "\t")
			if err != nil {
				log.Fatal(err)
			}
			jsonExpectedActor, _ := json.MarshalIndent(tc.expectedActor, "", "\t")

			// Write both json into a file
			ioutil.WriteFile("expectedActor.json", jsonExpectedActor, 0644)
			ioutil.WriteFile("actor.json", jsonActor, 0644)

			log.Println("Expected Actor: " + string(jsonExpectedActor))
			log.Println("Actor: " + string(jsonActor))

			assert.Equal(t, string(jsonExpectedActor), string(jsonActor))

			// Verify that all expectations were met

		})
	}
}
