package samples

import (
	"time"

	"github.com/ELITE-Kinoticketsystem/Backend-KTS/src/gen/KinoTicketSystem/model"
	"github.com/ELITE-Kinoticketsystem/Backend-KTS/src/models"
	"github.com/google/uuid"
)

func GetSampleActor() *models.ActorDTO {

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


func GetSampleActors() *[]models.GetActorsDTO {

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
