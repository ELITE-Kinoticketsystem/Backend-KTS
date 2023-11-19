package models

type Movie struct {
	Id          string      `json:"id"`
	Title       string      `json:"title"`
	Description string      `json:"description"`
	ReleaseDate string      `json:"releaseDate"`
	TimeInMin   int         `json:"timeInMin"`
	FSK         FSK         `json:"fsk"`
	Genre       Genre       `json:"genre"`
	Producers   *[]Producer `json:"producers"`
	Actors      *[]Actor    `json:"actors"`
}

type Genre string
type FSK string

const (
	Zero     FSK = "FSK ab 0 freigegeben"
	Six      FSK = "FSK ab 6 freigegeben"
	Twelve   FSK = "FSK ab 12 freigegeben"
	Sixteen  FSK = "FSK ab 16 freigegeben"
	Eighteen FSK = "FSK ab 18 freigegeben"

	Action  Genre = "Action"
	Drama   Genre = "Drama"
	Crime   Genre = "Crime"
	Fantasy Genre = "Fantasy"
	Western Genre = "Western"
	Romance Genre = "Romance"
)
