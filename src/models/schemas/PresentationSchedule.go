package models

type PresentationSchedule struct {
	Events *[]Event `json:"events"`
}