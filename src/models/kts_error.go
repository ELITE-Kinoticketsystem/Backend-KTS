package models

type KTSErrorMessage struct {
	ErrorMessage string `json:"errorMessage"`
}

type KTSError struct {
	KTSErrorMessage
	Status int
}
