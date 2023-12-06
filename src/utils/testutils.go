package utils

import (
	"fmt"
	"reflect"
	"time"

	"github.com/ELITE-Kinoticketsystem/Backend-KTS/src/.gen/KinoTicketSystem/model"
	"github.com/ELITE-Kinoticketsystem/Backend-KTS/src/models"
	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"github.com/google/uuid"
	"github.com/ELITE-Kinoticketsystem/Backend-KTS/src/.gen/KinoTicketSystem/model"
)

func GetSampleRegistrationData() models.RegistrationRequest {
	return models.RegistrationRequest{
		Username:  "Collinho el niño",
		Email:     "collin.forslund@gmail.com",
		Password:  "Passwort",
		FirstName: "Collin",
		LastName:  "Forslund",
	}
}

func GetSampleLoginData() models.LoginRequest {
	return models.LoginRequest{
		Username: "Collinho el niño",
		Password: "Passwort",
	}
}

func GetSampleUser() model.Users {
	id, _ := uuid.Parse("47CF752501DF45B7A3A9D3CB25AE939F")
	username := "Collinho el niño" 
	firstname := "Collin"
	lastname := "Forslund"
	return model.Users{
		ID:        &id,
		Username: &username,
		Email:     "collin.forslund@gmail.com",
		Password:  "$2a$10$vxXPPpLp5baQ7mzS1pNSEuk6ZW3mbx1Ej7u0tJnF5wferEFqT.qlK",
		Firstname: &firstname,
		Lastname:  &lastname,
	}
}

// Compare two users while ignoring their ids and hashed passwords.
func UserEqual(user1 model.Users, user2 model.Users) bool {
	return cmp.Equal(user1, user2, cmpopts.IgnoreFields(model.Users{}, "ID", "Password"))
}

type UserMatcher struct {
	user     model.Users
	password string
}

func (m UserMatcher) Matches(x interface{}) bool {
	user, ok := x.(model.Users)
	if !ok {
		return false
	}
	if !ComparePasswordHash(m.password, user.Password) {
		return false
	}
	m.user.Password = user.Password

	// ignore user_id
	m.user.ID = user.ID

	return reflect.DeepEqual(m.user, user)
}

func (m UserMatcher) String() string {
	return fmt.Sprintf("matches user %v and password %s", m.user, m.password)
}

func EqUserMatcher(u model.Users, password string) UserMatcher {
	return UserMatcher{user: u, password: password}
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

//

func GetSampleMovieById() *model.Movies {
	modelMovies := model.Movies{}
	uuid1 := uuid.MustParse("6ba7b826-9dad-11d1-80b4-00c04fd430c0")
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
	uuid1 := uuid.New()
	uuid2 := uuid.New()

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

func GetSampleGenreByName() *model.Genres {
	modelGenres := model.Genres{}
	uuid1 := uuid.MustParse("6ba7b820-9dad-11d1-80b4-00c04fd430c4")

	modelGenres = model.Genres{
		ID:        uuid1,
		GenreName: "Action",
	}

	return &modelGenres
}

func GetSampleMovieByIdWithGenre() *models.MovieWithGenres {
	movieWithGenre := models.MovieWithGenres{}

	banner := ""
	cover := ""
	trailer := ""
	rating := 5.0

	movieWithGenre = models.MovieWithGenres{
		Movies: model.Movies{
			ID:           uuid.MustParse("6ba7b827-9dad-11d1-80b4-00c04fd430c1"),
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
					ID:        uuid.MustParse("6ba7b820-9dad-11d1-80b4-00c04fd430c4"),
					GenreName: "Action",
				},
			},
			{
				model.Genres{
					ID:        uuid.MustParse("6ba7b821-9dad-11d1-80b4-00c04fd430c5"),
					GenreName: "Drama",
				},
			},
		},
	}

	return &movieWithGenre

}

func GetSampleGenreByNameWithMovies() *models.GenreWithMovies {
	genreByNameWithMovies := models.GenreWithMovies{}

	banner := ""
	cover := ""
	trailer := ""
	rating := 5.0

	genreByNameWithMovies = models.GenreWithMovies{
		Genres: model.Genres{
			ID:        uuid.MustParse("6ba7b820-9dad-11d1-80b4-00c04fd430c4"),
			GenreName: "Action",
		},
		Movies: []struct {
			model.Movies
		}{
			{
				model.Movies{
					ID:           uuid.MustParse("6ba7b827-9dad-11d1-80b4-00c04fd430c1"),
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
					ID:           uuid.MustParse("6ba7b828-9dad-11d1-80b4-00c04fd430c2"),
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

	banner := ""
	cover := ""
	trailer := ""
	rating := 5.0

	genresWithMovies = append(genresWithMovies, models.GenreWithMovies{
		Genres: model.Genres{
			ID:        uuid.MustParse("6ba7b820-9dad-11d1-80b4-00c04fd430c4"),
			GenreName: "Action",
		},
		Movies: []struct {
			model.Movies
		}{
			{
				model.Movies{
					ID:           uuid.MustParse("6ba7b827-9dad-11d1-80b4-00c04fd430c1"),
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
					ID:           uuid.MustParse("6ba7b828-9dad-11d1-80b4-00c04fd430c2"),
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
			ID:        uuid.MustParse("6ba7b821-9dad-11d1-80b4-00c04fd430c5"),
			GenreName: "Drama",
		},
		Movies: []struct {
			model.Movies
		}{
			{
				model.Movies{
					ID:           uuid.MustParse("6ba7b829-9dad-11d1-80b4-00c04fd430c3"),
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
					ID:           uuid.MustParse("6ba7b82a-9dad-11d1-80b4-00c04fd430c4"),
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
