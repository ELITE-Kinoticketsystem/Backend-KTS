package main

type Movie struct {
	Title       string      `json:"title"`
	Description string      `json:"description"`
	ReleaseDate string      `json:"releaseDate"`
	TimeInMin   int         `json:"timeInMin"`
	FSK         FSK         `json:"fsk"`
	Producers   *[]Producer `json:"producers"`
	Actors      *[]Actor    `json:"actors"`
}

type FSK string

const (
	Zero     FSK = "FSK ab 0 freigegeben"
	Six      FSK = "FSK ab 6 freigegeben"
	Twelve   FSK = "FSK ab 12 freigegeben"
	Sixteen  FSK = "FSK ab 16 freigegeben"
	Eighteen FSK = "FSK ab 18 freigegeben"
)