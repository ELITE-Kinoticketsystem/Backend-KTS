package samples

import (
	"log"
	"time"

	"github.com/ELITE-Kinoticketsystem/Backend-KTS/src/gen/KinoTicketSystem/model"
	"github.com/ELITE-Kinoticketsystem/Backend-KTS/src/models"
	"github.com/ELITE-Kinoticketsystem/Backend-KTS/src/myid"
)

func GetSampleMovies() *[]model.Movies {
	modelMovies := []model.Movies{}
	uuid1 := myid.New()
	uuid2 := myid.New()
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

func GetSampleMovieById() *model.Movies {
	modelMovies := model.Movies{}
	uuid1, err := myid.Parse("6ba7b826-9dad-11d1-80b4-00c04fd430c0")
	if err != nil {
		log.Print("GetSampleMovieById: Parsing UUID does not work")
	}
	banner := ""
	cover := ""
	trailer := ""
	rating := 5.0

	modelMovies = model.Movies{
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
	}

	return &modelMovies
}

func GetSampleGenres() *[]model.Genres {
	modelGenres := []model.Genres{}
	uuid1 := myid.New()
	uuid2 := myid.New()

	modelGenres = append(modelGenres, model.Genres{
		ID:        uuid1,
		GenreName: "Action",
	})

	modelGenres = append(modelGenres, model.Genres{
		ID:        uuid2,
		GenreName: "Drama",
	})

	return &modelGenres
}

func GetSampleGenre() *model.Genres {
	modelGenres := model.Genres{}
	uuid1 := myid.MustParse("6ba7b820-9dad-11d1-80b4-00c04fd430c4")

	modelGenres = model.Genres{
		ID:        uuid1,
		GenreName: "Action",
	}

	return &modelGenres
}

func GetSampleMovieByIdWithGenre() *models.MovieWithGenres {
	movieWithGenre := models.MovieWithGenres{}

	uuid1 := myid.MustParse("6ba7b827-9dad-11d1-80b4-00c04fd430c1")
	uuid2 := myid.MustParse("6ba7b820-9dad-11d1-80b4-00c04fd430c4")
	uuid3 := myid.MustParse("6ba7b821-9dad-11d1-80b4-00c04fd430c5")

	banner := ""
	cover := ""
	trailer := ""
	rating := 5.0

	movieWithGenre = models.MovieWithGenres{
		Movies: model.Movies{
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
		},
		Genres: []struct {
			model.Genres
		}{
			{
				model.Genres{
					ID:        uuid2,
					GenreName: "Action",
				},
			},
			{
				model.Genres{
					ID:        uuid3,
					GenreName: "Drama",
				},
			},
		},
	}

	return &movieWithGenre

}

func GetSampleGenreByNameWithMovies() *models.GenreWithMovies {
	genreByNameWithMovies := models.GenreWithMovies{}

	uuid1 := myid.MustParse("6ba7b820-9dad-11d1-80b4-00c04fd430c4")
	uuid2 := myid.MustParse("6ba7b827-9dad-11d1-80b4-00c04fd430c1")
	uuid3 := myid.MustParse("6ba7b828-9dad-11d1-80b4-00c04fd430c2")
	banner := ""
	cover := ""
	trailer := ""
	rating := 5.0

	genreByNameWithMovies = models.GenreWithMovies{
		Genres: model.Genres{
			ID:        uuid1,
			GenreName: "Action",
		},
		Movies: []struct {
			model.Movies
		}{
			{
				model.Movies{
					ID:           uuid2,
					Title:        "Test Movie 1",
					Description:  "Test Description 1",
					BannerPicURL: &banner,
					CoverPicURL:  &cover,
					TrailerURL:   &trailer,
					Rating:       &rating,
					ReleaseDate:  time.Now(),
					TimeInMin:    120,
					Fsk:          18,
				},
			},
			{
				model.Movies{
					ID:           uuid3,
					Title:        "Test Movie 2",
					Description:  "Test Description 2",
					BannerPicURL: &banner,
					CoverPicURL:  &cover,
					TrailerURL:   &trailer,
					Rating:       &rating,
					ReleaseDate:  time.Now(),
					TimeInMin:    120,
					Fsk:          18,
				},
			},
		},
	}

	return &genreByNameWithMovies
}

func GetSampleGenresWithMovies() *[]models.GenreWithMovies {
	genresWithMovies := []models.GenreWithMovies{}

	uuid1 := myid.MustParse("6ba7b820-9dad-11d1-80b4-00c04fd430c4")
	uuid2 := myid.MustParse("6ba7b827-9dad-11d1-80b4-00c04fd430c1")
	uuid3 := myid.MustParse("6ba7b828-9dad-11d1-80b4-00c04fd430c2")
	uuid4 := myid.MustParse("6ba7b821-9dad-11d1-80b4-00c04fd430c5")
	uuid5 := myid.MustParse("6ba7b829-9dad-11d1-80b4-00c04fd430c3")
	uuid6 := myid.MustParse("6ba7b82a-9dad-11d1-80b4-00c04fd430c4")
	banner := ""
	cover := ""
	trailer := ""
	rating := 5.0

	genresWithMovies = append(genresWithMovies, models.GenreWithMovies{
		Genres: model.Genres{
			ID:        uuid1,
			GenreName: "Action",
		},
		Movies: []struct {
			model.Movies
		}{
			{
				model.Movies{
					ID:           uuid2,
					Title:        "Test Movie 1",
					Description:  "Test Description 1",
					BannerPicURL: &banner,
					CoverPicURL:  &cover,
					TrailerURL:   &trailer,
					Rating:       &rating,
					ReleaseDate:  time.Now(),
					TimeInMin:    120,
					Fsk:          18,
				},
			},
			{
				model.Movies{
					ID:           uuid3,
					Title:        "Test Movie 2",
					Description:  "Test Description 2",
					BannerPicURL: &banner,
					CoverPicURL:  &cover,
					TrailerURL:   &trailer,
					Rating:       &rating,
					ReleaseDate:  time.Now(),
					TimeInMin:    120,
					Fsk:          18,
				},
			},
		},
	})

	genresWithMovies = append(genresWithMovies, models.GenreWithMovies{
		Genres: model.Genres{
			ID:        uuid4,
			GenreName: "Drama",
		},
		Movies: []struct {
			model.Movies
		}{
			{
				model.Movies{
					ID:           uuid5,
					Title:        "Test Movie 3",
					Description:  "Test Description 3",
					BannerPicURL: &banner,
					CoverPicURL:  &cover,
					TrailerURL:   &trailer,
					Rating:       &rating,
					ReleaseDate:  time.Now(),
					TimeInMin:    120,
					Fsk:          18,
				},
			},
			{
				model.Movies{
					ID:           uuid6,
					Title:        "Test Movie 4",
					Description:  "Test Description 4",
					BannerPicURL: &banner,
					CoverPicURL:  &cover,
					TrailerURL:   &trailer,
					Rating:       &rating,
					ReleaseDate:  time.Now(),
					TimeInMin:    120,
					Fsk:          18,
				},
			},
		},
	})

	return &genresWithMovies
}

func GetSampleMoviesWithGenres() *[]models.MovieWithGenres {
	movieWithGenres := []models.MovieWithGenres{}

	uuid1 := myid.MustParse("6ba7b820-9dad-11d1-80b4-00c04fd430c4")
	uuid2 := myid.MustParse("6ba7b827-9dad-11d1-80b4-00c04fd430c1")
	uuid3 := myid.MustParse("6ba7b828-9dad-11d1-80b4-00c04fd430c2")
	uuid4 := myid.MustParse("6ba7b821-9dad-11d1-80b4-00c04fd430c5")
	banner := ""
	cover := ""
	trailer := ""
	rating := 5.0

	movieWithGenres = append(movieWithGenres, models.MovieWithGenres{
		Movies: model.Movies{
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
		},
		Genres: []struct {
			model.Genres
		}{
			{
				model.Genres{
					ID:        uuid2,
					GenreName: "Action",
				},
			},
			{
				model.Genres{
					ID:        uuid3,
					GenreName: "Drama",
				},
			},
		},
	})

	movieWithGenres = append(movieWithGenres, models.MovieWithGenres{
		Movies: model.Movies{
			ID:           uuid4,
			Title:        "Test Movie 2",
			Description:  "Test Description 2",
			BannerPicURL: &banner,
			CoverPicURL:  &cover,
			TrailerURL:   &trailer,
			Rating:       &rating,
			ReleaseDate:  time.Now(),
			TimeInMin:    120,
			Fsk:          18,
		},
		Genres: []struct {
			model.Genres
		}{
			{
				model.Genres{
					ID:        uuid3,
					GenreName: "Drama",
				},
			},
		},
	})

	return &movieWithGenres
}

func GetSampleMovieByIdWithEverything() *models.MovieWithEverything {
	movieWithEverything := models.MovieWithEverything{}

	// uuid1 := myid.MustParse("6ba7b820-9dad-11d1-80b4-00c04fd430c4")
	movieId := myid.NewUUID()

	banner := ""
	cover := ""
	trailer := ""
	rating := 5.0
	isSpoiler := false

	movieWithEverything = models.MovieWithEverything{
		Movies: model.Movies{
			ID:           *movieId,
			Title:        "Test Movie 1",
			Description:  "Test Description 1",
			BannerPicURL: &banner,
			CoverPicURL:  &cover,
			TrailerURL:   &trailer,
			Rating:       &rating,
			ReleaseDate:  time.Now(),
			TimeInMin:    120,
			Fsk:          18,
		},
		Genres: []struct {
			model.Genres
		}{
			{
				model.Genres{
					ID:        myid.New(),
					GenreName: "Action",
				},
			},
		},
		Actors: []struct {
			model.Actors
		}{
			{
				model.Actors{
					ID:          myid.New(),
					Name:        "MaxActor Mustermann",
					Birthdate:   time.Now(),
					Description: "This is a description",
				},
			},
		},
		Producers: []struct {
			model.Producers
		}{
			{
				model.Producers{
					ID:          myid.New(),
					Name:        "MaxProducer Mustermann",
					Birthdate:   time.Now(),
					Description: "This is a description",
				},
			},
		},

		Reviews: []struct {
			Review   model.Reviews
			Username string `alias:"users.username"`
		}{
			{
				Review: model.Reviews{
					ID:        myid.New(),
					Rating:    4.0,
					Comment:   "This is a comment",
					Datetime:  time.Now(),
					IsSpoiler: &isSpoiler,
					UserID:    myid.New(),
					MovieID:   *movieId,
				},
				Username: "Max Mustermann",
			},
		},
	}

	return &movieWithEverything
}

func GetSampleMovieDTOCreate() *models.MovieDTOCreate {
	movieDTO := models.MovieDTOCreate{}

	uuid1 := myid.MustParse("6ba7b820-9dad-11d1-80b4-00c04fd430c4")
	uuid2 := myid.MustParse("6ba7b827-9dad-11d1-80b4-00c04fd430c1")
	uuid5 := myid.MustParse("6ba7b829-9dad-11d1-80b4-00c04fd430c3")
	uuid7 := myid.MustParse("6ba7b82b-9dad-11d1-80b4-00c04fd430c5")

	banner := ""
	cover := ""
	trailer := ""
	rating := 5.0

	movieDTO = models.MovieDTOCreate{
		Movies: model.Movies{
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
		},
		GenresID: []struct {
			ID *myid.UUID
		}{
			{
				ID: &uuid2,
			},
		},
		ActorsID: []struct {
			ID *myid.UUID
		}{
			{
				ID: &uuid5,
			},
		},
		ProducersID: []struct {
			ID *myid.UUID
		}{
			{
				ID: &uuid7,
			},
		},
	}

	return &movieDTO
}
