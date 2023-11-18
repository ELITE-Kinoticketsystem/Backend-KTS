package schemas

type PresentationSchedule struct {
	Events *[]Event `json:"events"`
}