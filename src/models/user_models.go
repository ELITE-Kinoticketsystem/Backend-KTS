package models

type RegistrationRequest struct {
	Username  string `json:"username"`
	Email     string `json:"email"`
	Password  string `json:"password"`
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
}

type CheckEmailRequest struct {
	Email string `json:"email"`
}

type CheckUsernameRequest struct {
	Username string `json:"username"`
}