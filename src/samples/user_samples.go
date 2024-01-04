package samples

import (
	"github.com/ELITE-Kinoticketsystem/Backend-KTS/src/gen/KinoTicketSystem/model"
	"github.com/ELITE-Kinoticketsystem/Backend-KTS/src/models"
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
	id := uuid.MustParse("47cf7525-01df-45b7-a3a9-d3cb25ae939f")
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
