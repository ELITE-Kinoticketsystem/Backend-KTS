package models

import (
	"github.com/ELITE-Kinoticketsystem/Backend-KTS/src/gen/KinoTicketSystem/model"
	"github.com/google/uuid"
)

type RegistrationRequest struct {
	Username  string `json:"username"`
	Email     string `json:"email"`
	Password  string `json:"password"`
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
}

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type LoginResponse struct {
	User         model.Users
	Token        string
	RefreshToken string
}

type CheckEmailRequest struct {
	Email string `json:"email" binding:"required"`
}

type CheckUsernameRequest struct {
	Username string `json:"username" binding:"required"`
}

type LoggedInResponse struct {
	LoggedIn bool       `json:"loggedIn"`
	Id       *uuid.UUID `json:"id"`
}
