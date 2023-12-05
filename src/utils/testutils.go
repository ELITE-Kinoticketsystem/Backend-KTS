package utils

import (
	"fmt"
	"reflect"

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
