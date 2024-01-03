package models

type CreateTheatreRequest struct {
	Name    string
	LogoUrl string
	Address struct {
		Street   string
		StreetNr string
		Zipcode  string
		City     string
		Country  string
	}
}
