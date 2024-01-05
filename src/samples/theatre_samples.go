package samples

import (
	"github.com/ELITE-Kinoticketsystem/Backend-KTS/src/gen/KinoTicketSystem/model"
	"github.com/ELITE-Kinoticketsystem/Backend-KTS/src/models"
	"github.com/ELITE-Kinoticketsystem/Backend-KTS/src/utils"
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

func GetSampleTheatres() []model.Theatres {
	return []model.Theatres{
		{
			ID:        utils.NewUUID(),
			Name:      "Theatre1",
			LogoURL:   utils.GetStringPointer("LogoUr1"),
			AddressID: utils.NewUUID(),
		},
		{
			ID:        utils.NewUUID(),
			Name:      "Theatre2",
			LogoURL:   utils.GetStringPointer("LogoUrl2"),
			AddressID: utils.NewUUID(),
		},
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

func GetSampleTheatreCreate() models.CreateTheatreRequest {
	logoUrl := "LogoURL"
	return models.CreateTheatreRequest{
		Name:    "Theatre",
		LogoUrl: logoUrl,

		Address: struct {
			Street   string
			StreetNr string
			Zipcode  string
			City     string
			Country  string
		}{
			Street:   "Street",
			StreetNr: "StreetNr",
			Zipcode:  "Zipcode",
			City:     "City",
			Country:  "Country",
		},
	}
}
