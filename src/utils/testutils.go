package utils

import (
	"database/sql/driver"
	"fmt"
	"reflect"
	"time"

	"github.com/ELITE-Kinoticketsystem/Backend-KTS/src/.gen/KinoTicketSystem/model"
	"github.com/ELITE-Kinoticketsystem/Backend-KTS/src/models"
	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"github.com/google/uuid"
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
		Username:  &username,
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

// for matching a struct except for uuid fields
type IdMatcher struct {
	value interface{}
}

func (m IdMatcher) Matches(otherValue interface{}) bool {
	return cmp.Equal(m.value, otherValue, cmpopts.IgnoreTypes(&uuid.UUID{}))
}

func (m IdMatcher) String() string {
	return fmt.Sprintf("matches %v", m.value)
}

// Returns a matcher that matches the struct except for the uuid fields.
func EqExceptId(value interface{}) IdMatcher {
	return IdMatcher{value: value}
}

// for matching a uuid with its binary representation
type UUIDMatcher struct {
	id *uuid.UUID
}

func (m UUIDMatcher) Match(v driver.Value) bool {
	bytes, ok := v.(string)
	if !ok {
		return false
	}
	id, err := m.id.MarshalBinary()
	if err != nil {
		return false
	}
	return string(id) == bytes
}

// Returns a matcher that matches the uuid with its binary representation.
func EqUUID(id *uuid.UUID) UUIDMatcher {
	return UUIDMatcher{id: id}
}

func GetSampleProducers() *[]models.GetProducersDTO {
	var producers []models.GetProducersDTO

	uuid1 := uuid.New()
	uuid2 := uuid.New()
	uuid3 := uuid.New()
	uuid4 := uuid.New()
	uuid5 := uuid.New()
	uuid6 := uuid.New()

	picUrl := "https://www.picture.google.com"

	producers = append(producers, models.GetProducersDTO{
		Producers: model.Producers{
			ID:          &uuid1,
			Name:        "Producer 1",
			Birthdate:   time.Now(),
			Description: "Description 1",
			PicURL:      &picUrl,
		},
		Pictures: []model.ProducerPictures{
			{
				ID:         &uuid2,
				ProducerID: &uuid1,
				PicURL:     &picUrl,
			},
			{
				ID:         &uuid3,
				ProducerID: &uuid1,
				PicURL:     &picUrl,
			},
		},
	})

	producers = append(producers, models.GetProducersDTO{
		Producers: model.Producers{
			ID:          &uuid4,
			Name:        "Producer 1",
			Birthdate:   time.Now(),
			Description: "Description 1",
			PicURL:      &picUrl,
		},
		Pictures: []model.ProducerPictures{
			{
				ID:         &uuid5,
				ProducerID: &uuid4,
				PicURL:     &picUrl,
			},
			{
				ID:         &uuid6,
				ProducerID: &uuid4,
				PicURL:     &picUrl,
			},
		},
	})

	return &producers
}

func GetSampleProducer() *model.Producers {
	uuid1 := uuid.New()

	picUrl := "https://www.picture.google.com"

	return &model.Producers{
		ID:          &uuid1,
		Name:        "Producer 1",
		Birthdate:   time.Now(),
		Description: "Description 1",
		PicURL:      &picUrl,
	}
}

func GetSampleProducerDTO() *models.ProducerDTO {
	uuid1 := uuid.New()
	uuid2 := uuid.New()
	uuid3 := uuid.New()
	uuid4 := uuid.New()
	uuid5 := uuid.New()

	banner := "https://www.banner.google.com"
	cover := "https://www.cover.google.com"
	trailer_url := "https://www.trailer.google.com"
	rating := 5.0

	picUrl := "https://www.picture.google.com"

	return &models.ProducerDTO{
		Producers: model.Producers{
			ID:          &uuid1,
			Name:        "Producer 1",
			Birthdate:   time.Now(),
			Description: "Description 1",
			PicURL:      &picUrl,
		},
		Pictures: []model.ProducerPictures{
			{
				ID:         &uuid2,
				ProducerID: &uuid1,
				PicURL:     &picUrl,
			},
			{
				ID:         &uuid3,
				ProducerID: &uuid1,
				PicURL:     &picUrl,
			},
		},
		Movies: []model.Movies{
			{
				ID:           &uuid4,
				Title:        "Movie 1",
				BannerPicURL: &banner,
				CoverPicURL:  &cover,
				TrailerURL:   &trailer_url,
				Rating:       &rating,
				Description:  "Description 1",
				ReleaseDate:  time.Now(),
				TimeInMin:    120,
				Fsk:          12,
			},
			{
				ID:           &uuid5,
				Title:        "Movie 2",
				BannerPicURL: &banner,
				CoverPicURL:  &cover,
				TrailerURL:   &trailer_url,
				Rating:       &rating,
				Description:  "Description 2",
				ReleaseDate:  time.Now(),
				TimeInMin:    120,
				Fsk:          12,
			},
		},
	}
}
