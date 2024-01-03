package samples

import (
	"github.com/ELITE-Kinoticketsystem/Backend-KTS/src/gen/KinoTicketSystem/model"
	"github.com/google/uuid"
)

func GetSampleAddress() model.Addresses {
	id := uuid.MustParse("b7f445cc-2fa5-42bb-967e-9d9585a87f0d")
	return model.Addresses{
		ID:       &id,
		Street:   "Street",
		StreetNr: "StreetNr",
		Zipcode:  "Zipcode",
		City:     "City",
		Country:  "Country",
	}
}

func GetSampleTheatre() model.Theatres {
	id := uuid.New()
	addressId := uuid.MustParse("b7f445cc-2fa5-42bb-967e-9d9585a87f0d")
	logoUrl := "LogoURL"
	return model.Theatres{
		ID:        &id,
		Name:      "Theatre",
		LogoURL:   &logoUrl,
		AddressID: &addressId,
	}
}
