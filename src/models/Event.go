package main

type Event struct {
	Title   string    `json:"title"`
	Start   string    `json:"start"` //TODO: fix datatype
	End     string    `json:"end"`   //TODO: fix datatype
	Tickets *[]Ticket `json:"tickets"`
}

type SpecialEvent struct {
	Event
	Movies *[]Movie `json:"movies"`
}

type Showing struct {
	Event
	Movie Movie `json:"movie"`
}