package schemas

type Theatre struct {
	ID                   int                   `json:"id"`
	Name                 string                `json:"name"`
	Address              Address               `json:"address"`
	Theatres             []CinemaHall          `json:"theatres"`
	PresentationSchedule *PresentationSchedule `json:"presentationSchedule"`
}