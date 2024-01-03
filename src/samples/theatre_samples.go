package samples

import (
	"github.com/ELITE-Kinoticketsystem/Backend-KTS/src/gen/KinoTicketSystem/model"
)

func GetSampleAddress() model.Addresses {
	return model.Addresses{
		/* Id */
		Street:   "Street",
		StreetNr: "StreetNr",
		Zipcode:  "Zipcode",
		City:     "City",
		Country:  "Country",
	}
}

func GetSampleTheatre() model.Theatres {
	return model.Theatres{
		/* ID */
		Name: "Theatre",
		/* AddressID */
	}
}
