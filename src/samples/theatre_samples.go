package samples

import (
	"github.com/ELITE-Kinoticketsystem/Backend-KTS/src/gen/KinoTicketSystem/model"
	"github.com/ELITE-Kinoticketsystem/Backend-KTS/src/models"
	"github.com/ELITE-Kinoticketsystem/Backend-KTS/src/utils"
	"github.com/google/uuid"
)

const (
	addressId       = "b7f445cc-2fa5-42bb-967e-9d9585a87f0d"
	theatreId       = "aceb8bf0-8ff8-416a-9199-df6f64f4ce6d"
	cinemaHallId    = "4fe38ff6-6498-4832-92db-ac1c0a00e8df"
	seatId          = "795ac16e-db79-4178-bdd0-dc553b76900b"
	seatCategoryId1 = "84df6b9f-6633-437d-a17c-d75a82d61c90"
	seatCategoryId2 = "77e4f7fa-b3a7-4a50-9b8f-5543e70fd8d0"
	seatCategoryId3 = "feb75348-9dfd-4680-8e54-18d89d02b50a"
)

func GetSampleAddress() model.Addresses {
	id := uuid.MustParse(addressId)
	return model.Addresses{
		ID:       &id,
		Street:   "Street",
		StreetNr: "StreetNr",
		Zipcode:  "Zipcode",
		City:     "City",
		Country:  "Country",
	}
}

func GetSampleTheatres() []models.GetTheatreWithAddress {
	return []models.GetTheatreWithAddress{
		{
			ID:      utils.NewUUID(),
			Name:    "Theatre1",
			LogoUrl: utils.GetStringPointer("LogoUr1"),
			Address: model.Addresses{
				ID:       utils.NewUUID(),
				Street:   "Street1",
				StreetNr: "StreetNr1",
				Zipcode:  "26382",
				City:     "City1",
				Country:  "Country1",
			},
		},
		{
			ID:      utils.NewUUID(),
			Name:    "Theatre2",
			LogoUrl: utils.GetStringPointer("LogoUrl2"),
			Address: model.Addresses{
				ID:       utils.NewUUID(),
				Street:   "Street2",
				StreetNr: "StreetNr2",
				Zipcode:  "26382",
				City:     "City2",
				Country:  "Country2",
			},
		},
	}
}

func GetSampleTheatre() model.Theatres {
	id := uuid.MustParse(theatreId)
	addressId := uuid.MustParse(addressId)
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

func GetSampleCreateCinemaHallRequest() models.CreateCinemaHallRequest {
	theatreId := uuid.MustParse(theatreId)
	return models.CreateCinemaHallRequest{
		HallName: "HallName",
		Width:    23,
		Height:   15,
		Seats: []struct {
			X        int
			Y        int
			Type     string
			Category string
		}{

			{
				Category: "regular",
				Type:     "regular",
				X:        0,
				Y:        0,
			},
			{
				Category: "regular",
				Type:     "regular",
				X:        1,
				Y:        0,
			},
			{
				Category: "regular",
				Type:     "regular",
				X:        2,
				Y:        0,
			},
			{
				Category: "regular",
				Type:     "regular",
				X:        3,
				Y:        0,
			},
			{
				Category: "regular",
				Type:     "regular",
				X:        4,
				Y:        0,
			},
			{
				Category: "regular",
				Type:     "regular",
				X:        5,
				Y:        0,
			},
			{
				Category: "regular",
				Type:     "regular",
				X:        6,
				Y:        0,
			},
			{
				Category: "regular",
				Type:     "regular",
				X:        7,
				Y:        0,
			},
			{
				Category: "regular",
				Type:     "regular",
				X:        8,
				Y:        0,
			},
			{
				Category: "regular",
				Type:     "regular",
				X:        9,
				Y:        0,
			},
			{
				Category: "regular",
				Type:     "regular",
				X:        10,
				Y:        0,
			},
			{
				Category: "regular",
				Type:     "regular",
				X:        11,
				Y:        0,
			},
			{
				Category: "regular",
				Type:     "regular",
				X:        12,
				Y:        0,
			},
			{
				Category: "regular",
				Type:     "regular",
				X:        13,
				Y:        0,
			},
			{
				Category: "regular",
				Type:     "regular",
				X:        14,
				Y:        0,
			},
			{
				Category: "regular",
				Type:     "regular",
				X:        15,
				Y:        0,
			},
			{
				Category: "regular",
				Type:     "regular",
				X:        16,
				Y:        0,
			},
			{
				Category: "regular",
				Type:     "regular",
				X:        17,
				Y:        0,
			},
			{
				Category: "regular",
				Type:     "regular",
				X:        18,
				Y:        0,
			},
			{
				Category: "regular",
				Type:     "regular",
				X:        19,
				Y:        0,
			},
			{
				Category: "regular",
				Type:     "regular",
				X:        20,
				Y:        0,
			},
			{
				Category: "regular",
				Type:     "regular",
				X:        21,
				Y:        0,
			},
			{
				Category: "regular",
				Type:     "regular",
				X:        22,
				Y:        0,
			},
			{
				Category: "regular",
				Type:     "regular",
				X:        0,
				Y:        1,
			},
			{
				Category: "regular",
				Type:     "regular",
				X:        1,
				Y:        1,
			},
			{
				Category: "regular",
				Type:     "regular",
				X:        2,
				Y:        1,
			},
			{
				Category: "regular",
				Type:     "regular",
				X:        3,
				Y:        1,
			},
			{
				Category: "regular",
				Type:     "regular",
				X:        4,
				Y:        1,
			},
			{
				Category: "regular",
				Type:     "regular",
				X:        5,
				Y:        1,
			},
			{
				Category: "regular",
				Type:     "regular",
				X:        6,
				Y:        1,
			},
			{
				Category: "regular",
				Type:     "regular",
				X:        7,
				Y:        1,
			},
			{
				Category: "regular",
				Type:     "regular",
				X:        8,
				Y:        1,
			},
			{
				Category: "regular",
				Type:     "regular",
				X:        9,
				Y:        1,
			},
			{
				Category: "regular",
				Type:     "regular",
				X:        10,
				Y:        1,
			},
			{
				Category: "regular",
				Type:     "regular",
				X:        11,
				Y:        1,
			},
			{
				Category: "regular",
				Type:     "regular",
				X:        12,
				Y:        1,
			},
			{
				Category: "regular",
				Type:     "regular",
				X:        13,
				Y:        1,
			},
			{
				Category: "regular",
				Type:     "regular",
				X:        14,
				Y:        1,
			},
			{
				Category: "regular",
				Type:     "regular",
				X:        15,
				Y:        1,
			},
			{
				Category: "regular",
				Type:     "regular",
				X:        16,
				Y:        1,
			},
			{
				Category: "regular",
				Type:     "regular",
				X:        17,
				Y:        1,
			},
			{
				Category: "regular",
				Type:     "regular",
				X:        18,
				Y:        1,
			},
			{
				Category: "regular",
				Type:     "regular",
				X:        19,
				Y:        1,
			},
			{
				Category: "regular",
				Type:     "regular",
				X:        20,
				Y:        1,
			},
			{
				Category: "regular",
				Type:     "regular",
				X:        21,
				Y:        1,
			},
			{
				Category: "regular",
				Type:     "regular",
				X:        22,
				Y:        1,
			},
			{
				Category: "regular",
				Type:     "regular",
				X:        0,
				Y:        2,
			},
			{
				Category: "regular",
				Type:     "regular",
				X:        1,
				Y:        2,
			},
			{
				Category: "regular",
				Type:     "regular",
				X:        2,
				Y:        2,
			},
			{
				Category: "regular",
				Type:     "regular",
				X:        3,
				Y:        2,
			},
			{
				Category: "regular",
				Type:     "regular",
				X:        4,
				Y:        2,
			},
			{
				Category: "regular",
				Type:     "regular",
				X:        5,
				Y:        2,
			},
			{
				Category: "regular",
				Type:     "regular",
				X:        6,
				Y:        2,
			},
			{
				Category: "regular",
				Type:     "regular",
				X:        7,
				Y:        2,
			},
			{
				Category: "regular",
				Type:     "regular",
				X:        8,
				Y:        2,
			},
			{
				Category: "regular",
				Type:     "regular",
				X:        9,
				Y:        2,
			},
			{
				Category: "regular",
				Type:     "regular",
				X:        10,
				Y:        2,
			},
			{
				Category: "regular",
				Type:     "regular",
				X:        11,
				Y:        2,
			},
			{
				Category: "regular",
				Type:     "regular",
				X:        12,
				Y:        2,
			},
			{
				Category: "regular",
				Type:     "regular",
				X:        13,
				Y:        2,
			},
			{
				Category: "regular",
				Type:     "regular",
				X:        14,
				Y:        2,
			},
			{
				Category: "regular",
				Type:     "regular",
				X:        15,
				Y:        2,
			},
			{
				Category: "regular",
				Type:     "regular",
				X:        16,
				Y:        2,
			},
			{
				Category: "regular",
				Type:     "regular",
				X:        17,
				Y:        2,
			},
			{
				Category: "regular",
				Type:     "regular",
				X:        18,
				Y:        2,
			},
			{
				Category: "regular",
				Type:     "regular",
				X:        19,
				Y:        2,
			},
			{
				Category: "regular",
				Type:     "regular",
				X:        20,
				Y:        2,
			},
			{
				Category: "regular",
				Type:     "regular",
				X:        21,
				Y:        2,
			},
			{
				Category: "regular",
				Type:     "regular",
				X:        22,
				Y:        2,
			},
			{
				Category: "regular",
				Type:     "regular",
				X:        0,
				Y:        3,
			},
			{
				Category: "regular",
				Type:     "regular",
				X:        1,
				Y:        3,
			},
			{
				Category: "regular",
				Type:     "regular",
				X:        2,
				Y:        3,
			},
			{
				Category: "regular",
				Type:     "regular",
				X:        3,
				Y:        3,
			},
			{
				Category: "regular",
				Type:     "regular",
				X:        4,
				Y:        3,
			},
			{
				Category: "regular",
				Type:     "regular",
				X:        5,
				Y:        3,
			},
			{
				Category: "regular",
				Type:     "regular",
				X:        6,
				Y:        3,
			},
			{
				Category: "regular",
				Type:     "regular",
				X:        7,
				Y:        3,
			},
			{
				Category: "regular",
				Type:     "regular",
				X:        8,
				Y:        3,
			},
			{
				Category: "regular",
				Type:     "regular",
				X:        9,
				Y:        3,
			},
			{
				Category: "regular",
				Type:     "regular",
				X:        10,
				Y:        3,
			},
			{
				Category: "regular",
				Type:     "regular",
				X:        11,
				Y:        3,
			},
			{
				Category: "regular",
				Type:     "regular",
				X:        12,
				Y:        3,
			},
			{
				Category: "regular",
				Type:     "regular",
				X:        13,
				Y:        3,
			},
			{
				Category: "regular",
				Type:     "regular",
				X:        14,
				Y:        3,
			},
			{
				Category: "regular",
				Type:     "regular",
				X:        15,
				Y:        3,
			},
			{
				Category: "regular",
				Type:     "regular",
				X:        16,
				Y:        3,
			},
			{
				Category: "regular",
				Type:     "regular",
				X:        17,
				Y:        3,
			},
			{
				Category: "regular",
				Type:     "regular",
				X:        18,
				Y:        3,
			},
			{
				Category: "regular",
				Type:     "regular",
				X:        19,
				Y:        3,
			},
			{
				Category: "regular",
				Type:     "regular",
				X:        20,
				Y:        3,
			},
			{
				Category: "regular",
				Type:     "regular",
				X:        21,
				Y:        3,
			},
			{
				Category: "regular",
				Type:     "regular",
				X:        22,
				Y:        3,
			},
			{
				Category: "regular",
				Type:     "regular",
				X:        0,
				Y:        4,
			},
			{
				Category: "regular",
				Type:     "regular",
				X:        1,
				Y:        4,
			},
			{
				Category: "regular",
				Type:     "regular",
				X:        2,
				Y:        4,
			},
			{
				Category: "regular",
				Type:     "regular",
				X:        3,
				Y:        4,
			},
			{
				Category: "regular",
				Type:     "regular",
				X:        4,
				Y:        4,
			},
			{
				Category: "regular",
				Type:     "regular",
				X:        5,
				Y:        4,
			},
			{
				Category: "regular",
				Type:     "regular",
				X:        6,
				Y:        4,
			},
			{
				Category: "regular",
				Type:     "regular",
				X:        7,
				Y:        4,
			},
			{
				Category: "regular",
				Type:     "regular",
				X:        8,
				Y:        4,
			},
			{
				Category: "regular",
				Type:     "regular",
				X:        9,
				Y:        4,
			},
			{
				Category: "regular",
				Type:     "regular",
				X:        10,
				Y:        4,
			},
			{
				Category: "regular",
				Type:     "regular",
				X:        11,
				Y:        4,
			},
			{
				Category: "regular",
				Type:     "regular",
				X:        12,
				Y:        4,
			},
			{
				Category: "regular",
				Type:     "regular",
				X:        13,
				Y:        4,
			},
			{
				Category: "regular",
				Type:     "regular",
				X:        14,
				Y:        4,
			},
			{
				Category: "regular",
				Type:     "regular",
				X:        15,
				Y:        4,
			},
			{
				Category: "regular",
				Type:     "regular",
				X:        16,
				Y:        4,
			},
			{
				Category: "regular",
				Type:     "regular",
				X:        17,
				Y:        4,
			},
			{
				Category: "regular",
				Type:     "regular",
				X:        18,
				Y:        4,
			},
			{
				Category: "regular",
				Type:     "regular",
				X:        19,
				Y:        4,
			},
			{
				Category: "regular",
				Type:     "regular",
				X:        20,
				Y:        4,
			},
			{
				Category: "regular",
				Type:     "regular",
				X:        21,
				Y:        4,
			},
			{
				Category: "regular",
				Type:     "regular",
				X:        22,
				Y:        4,
			},
			{
				Category: "regular",
				Type:     "regular",
				X:        0,
				Y:        5,
			},
			{
				Category: "regular",
				Type:     "regular",
				X:        1,
				Y:        5,
			},
			{
				Category: "regular",
				Type:     "regular",
				X:        2,
				Y:        5,
			},
			{
				Category: "regular",
				Type:     "regular",
				X:        3,
				Y:        5,
			},
			{
				Category: "regular",
				Type:     "regular",
				X:        4,
				Y:        5,
			},
			{
				Category: "regular",
				Type:     "regular",
				X:        5,
				Y:        5,
			},
			{
				Category: "regular",
				Type:     "regular",
				X:        6,
				Y:        5,
			},
			{
				Category: "regular",
				Type:     "regular",
				X:        7,
				Y:        5,
			},
			{
				Category: "regular",
				Type:     "regular",
				X:        8,
				Y:        5,
			},
			{
				Category: "regular",
				Type:     "regular",
				X:        9,
				Y:        5,
			},
			{
				Category: "regular",
				Type:     "regular",
				X:        10,
				Y:        5,
			},
			{
				Category: "regular",
				Type:     "regular",
				X:        11,
				Y:        5,
			},
			{
				Category: "regular",
				Type:     "regular",
				X:        12,
				Y:        5,
			},
			{
				Category: "regular",
				Type:     "regular",
				X:        13,
				Y:        5,
			},
			{
				Category: "regular",
				Type:     "regular",
				X:        14,
				Y:        5,
			},
			{
				Category: "regular",
				Type:     "regular",
				X:        15,
				Y:        5,
			},
			{
				Category: "regular",
				Type:     "regular",
				X:        16,
				Y:        5,
			},
			{
				Category: "regular",
				Type:     "regular",
				X:        17,
				Y:        5,
			},
			{
				Category: "regular",
				Type:     "regular",
				X:        18,
				Y:        5,
			},
			{
				Category: "regular",
				Type:     "regular",
				X:        19,
				Y:        5,
			},
			{
				Category: "regular",
				Type:     "regular",
				X:        20,
				Y:        5,
			},
			{
				Category: "regular",
				Type:     "regular",
				X:        21,
				Y:        5,
			},
			{
				Category: "regular",
				Type:     "regular",
				X:        22,
				Y:        5,
			},
			{
				Category: "regular",
				Type:     "regular",
				X:        0,
				Y:        6,
			},
			{
				Category: "regular",
				Type:     "regular",
				X:        1,
				Y:        6,
			},
			{
				Category: "regular",
				Type:     "regular",
				X:        2,
				Y:        6,
			},
			{
				Category: "regular",
				Type:     "regular",
				X:        3,
				Y:        6,
			},
			{
				Category: "regular",
				Type:     "regular",
				X:        4,
				Y:        6,
			},
			{
				Category: "regular",
				Type:     "regular",
				X:        5,
				Y:        6,
			},
			{
				Category: "regular",
				Type:     "regular",
				X:        6,
				Y:        6,
			},
			{
				Category: "regular",
				Type:     "regular",
				X:        7,
				Y:        6,
			},
			{
				Category: "regular",
				Type:     "regular",
				X:        8,
				Y:        6,
			},
			{
				Category: "regular",
				Type:     "regular",
				X:        9,
				Y:        6,
			},
			{
				Category: "regular",
				Type:     "regular",
				X:        10,
				Y:        6,
			},
			{
				Category: "regular",
				Type:     "regular",
				X:        11,
				Y:        6,
			},
			{
				Category: "regular",
				Type:     "regular",
				X:        12,
				Y:        6,
			},
			{
				Category: "regular",
				Type:     "regular",
				X:        13,
				Y:        6,
			},
			{
				Category: "regular",
				Type:     "regular",
				X:        14,
				Y:        6,
			},
			{
				Category: "regular",
				Type:     "regular",
				X:        15,
				Y:        6,
			},
			{
				Category: "regular",
				Type:     "regular",
				X:        16,
				Y:        6,
			},
			{
				Category: "regular",
				Type:     "regular",
				X:        17,
				Y:        6,
			},
			{
				Category: "regular",
				Type:     "regular",
				X:        18,
				Y:        6,
			},
			{
				Category: "regular",
				Type:     "regular",
				X:        19,
				Y:        6,
			},
			{
				Category: "regular",
				Type:     "regular",
				X:        20,
				Y:        6,
			},
			{
				Category: "regular",
				Type:     "regular",
				X:        21,
				Y:        6,
			},
			{
				Category: "regular",
				Type:     "regular",
				X:        22,
				Y:        6,
			},
			{
				Category: "regular",
				Type:     "regular",
				X:        0,
				Y:        8,
			},
			{
				Category: "regular",
				Type:     "regular",
				X:        1,
				Y:        8,
			},
			{
				Category: "regular",
				Type:     "regular",
				X:        2,
				Y:        8,
			},
			{
				Category: "regular",
				Type:     "regular",
				X:        3,
				Y:        8,
			},
			{
				Category: "regular",
				Type:     "regular",
				X:        4,
				Y:        8,
			},
			{
				Category: "regular",
				Type:     "regular",
				X:        5,
				Y:        8,
			},
			{
				Category: "regular",
				Type:     "regular",
				X:        6,
				Y:        8,
			},
			{
				Category: "regular",
				Type:     "regular",
				X:        7,
				Y:        8,
			},
			{
				Category: "regular",
				Type:     "regular",
				X:        8,
				Y:        8,
			},
			{
				Category: "regular",
				Type:     "regular",
				X:        9,
				Y:        8,
			},
			{
				Category: "regular",
				Type:     "regular",
				X:        10,
				Y:        8,
			},
			{
				Category: "regular",
				Type:     "regular",
				X:        11,
				Y:        8,
			},
			{
				Category: "regular",
				Type:     "regular",
				X:        12,
				Y:        8,
			},
			{
				Category: "regular",
				Type:     "regular",
				X:        13,
				Y:        8,
			},
			{
				Category: "regular",
				Type:     "regular",
				X:        14,
				Y:        8,
			},
			{
				Category: "regular",
				Type:     "regular",
				X:        15,
				Y:        8,
			},
			{
				Category: "regular",
				Type:     "regular",
				X:        16,
				Y:        8,
			},
			{
				Category: "regular",
				Type:     "regular",
				X:        17,
				Y:        8,
			},
			{
				Category: "regular",
				Type:     "regular",
				X:        18,
				Y:        8,
			},
			{
				Category: "regular",
				Type:     "regular",
				X:        19,
				Y:        8,
			},
			{
				Category: "regular",
				Type:     "regular",
				X:        20,
				Y:        8,
			},
			{
				Category: "regular",
				Type:     "regular",
				X:        21,
				Y:        8,
			},
			{
				Category: "regular",
				Type:     "regular",
				X:        22,
				Y:        8,
			},
			{
				Category: "regular",
				Type:     "regular",
				X:        0,
				Y:        9,
			},
			{
				Category: "regular",
				Type:     "regular",
				X:        1,
				Y:        9,
			},
			{
				Category: "regular",
				Type:     "regular",
				X:        2,
				Y:        9,
			},
			{
				Category: "regular",
				Type:     "regular",
				X:        3,
				Y:        9,
			},
			{
				Category: "regular",
				Type:     "regular",
				X:        4,
				Y:        9,
			},
			{
				Category: "regular",
				Type:     "regular",
				X:        5,
				Y:        9,
			},
			{
				Category: "regular",
				Type:     "regular",
				X:        6,
				Y:        9,
			},
			{
				Category: "regular",
				Type:     "regular",
				X:        7,
				Y:        9,
			},
			{
				Category: "regular",
				Type:     "regular",
				X:        8,
				Y:        9,
			},
			{
				Category: "regular",
				Type:     "regular",
				X:        9,
				Y:        9,
			},
			{
				Category: "regular",
				Type:     "regular",
				X:        10,
				Y:        9,
			},
			{
				Category: "regular",
				Type:     "regular",
				X:        11,
				Y:        9,
			},
			{
				Category: "regular",
				Type:     "regular",
				X:        12,
				Y:        9,
			},
			{
				Category: "regular",
				Type:     "regular",
				X:        13,
				Y:        9,
			},
			{
				Category: "regular",
				Type:     "regular",
				X:        14,
				Y:        9,
			},
			{
				Category: "regular",
				Type:     "regular",
				X:        15,
				Y:        9,
			},
			{
				Category: "regular",
				Type:     "regular",
				X:        16,
				Y:        9,
			},
			{
				Category: "regular",
				Type:     "regular",
				X:        17,
				Y:        9,
			},
			{
				Category: "regular",
				Type:     "regular",
				X:        18,
				Y:        9,
			},
			{
				Category: "regular",
				Type:     "regular",
				X:        19,
				Y:        9,
			},
			{
				Category: "regular",
				Type:     "regular",
				X:        20,
				Y:        9,
			},
			{
				Category: "regular",
				Type:     "regular",
				X:        21,
				Y:        9,
			},
			{
				Category: "regular",
				Type:     "regular",
				X:        22,
				Y:        9,
			},
			{
				Category: "regular",
				Type:     "regular",
				X:        0,
				Y:        10,
			},
			{
				Category: "regular",
				Type:     "regular",
				X:        1,
				Y:        10,
			},
			{
				Category: "regular",
				Type:     "regular",
				X:        2,
				Y:        10,
			},
			{
				Category: "regular",
				Type:     "regular",
				X:        3,
				Y:        10,
			},
			{
				Category: "regular",
				Type:     "regular",
				X:        4,
				Y:        10,
			},
			{
				Category: "regular",
				Type:     "regular",
				X:        5,
				Y:        10,
			},
			{
				Category: "regular",
				Type:     "regular",
				X:        6,
				Y:        10,
			},
			{
				Category: "regular",
				Type:     "regular",
				X:        7,
				Y:        10,
			},
			{
				Category: "regular",
				Type:     "regular",
				X:        8,
				Y:        10,
			},
			{
				Category: "regular",
				Type:     "regular",
				X:        9,
				Y:        10,
			},
			{
				Category: "regular",
				Type:     "regular",
				X:        10,
				Y:        10,
			},
			{
				Category: "regular",
				Type:     "regular",
				X:        11,
				Y:        10,
			},
			{
				Category: "regular",
				Type:     "regular",
				X:        12,
				Y:        10,
			},
			{
				Category: "regular",
				Type:     "regular",
				X:        13,
				Y:        10,
			},
			{
				Category: "regular",
				Type:     "regular",
				X:        14,
				Y:        10,
			},
			{
				Category: "regular",
				Type:     "regular",
				X:        15,
				Y:        10,
			},
			{
				Category: "regular",
				Type:     "regular",
				X:        16,
				Y:        10,
			},
			{
				Category: "regular",
				Type:     "regular",
				X:        17,
				Y:        10,
			},
			{
				Category: "regular",
				Type:     "regular",
				X:        18,
				Y:        10,
			},
			{
				Category: "regular",
				Type:     "regular",
				X:        19,
				Y:        10,
			},
			{
				Category: "regular",
				Type:     "regular",
				X:        20,
				Y:        10,
			},
			{
				Category: "regular",
				Type:     "regular",
				X:        21,
				Y:        10,
			},
			{
				Category: "regular",
				Type:     "regular",
				X:        22,
				Y:        10,
			},
			{
				Category: "regular",
				Type:     "regular",
				X:        0,
				Y:        11,
			},
			{
				Category: "regular",
				Type:     "regular",
				X:        1,
				Y:        11,
			},
			{
				Category: "regular",
				Type:     "regular",
				X:        2,
				Y:        11,
			},
			{
				Category: "regular",
				Type:     "regular",
				X:        3,
				Y:        11,
			},
			{
				Category: "regular",
				Type:     "regular",
				X:        4,
				Y:        11,
			},
			{
				Category: "regular",
				Type:     "regular",
				X:        5,
				Y:        11,
			},
			{
				Category: "regular",
				Type:     "regular",
				X:        6,
				Y:        11,
			},
			{
				Category: "regular",
				Type:     "regular",
				X:        7,
				Y:        11,
			},
			{
				Category: "regular",
				Type:     "regular",
				X:        8,
				Y:        11,
			},
			{
				Category: "regular",
				Type:     "regular",
				X:        9,
				Y:        11,
			},
			{
				Category: "regular",
				Type:     "regular",
				X:        10,
				Y:        11,
			},
			{
				Category: "regular",
				Type:     "regular",
				X:        11,
				Y:        11,
			},
			{
				Category: "regular",
				Type:     "regular",
				X:        12,
				Y:        11,
			},
			{
				Category: "regular",
				Type:     "regular",
				X:        13,
				Y:        11,
			},
			{
				Category: "regular",
				Type:     "regular",
				X:        14,
				Y:        11,
			},
			{
				Category: "regular",
				Type:     "regular",
				X:        15,
				Y:        11,
			},
			{
				Category: "regular",
				Type:     "regular",
				X:        16,
				Y:        11,
			},
			{
				Category: "regular",
				Type:     "regular",
				X:        17,
				Y:        11,
			},
			{
				Category: "regular",
				Type:     "regular",
				X:        18,
				Y:        11,
			},
			{
				Category: "regular",
				Type:     "regular",
				X:        19,
				Y:        11,
			},
			{
				Category: "regular",
				Type:     "regular",
				X:        20,
				Y:        11,
			},
			{
				Category: "regular",
				Type:     "regular",
				X:        21,
				Y:        11,
			},
			{
				Category: "regular",
				Type:     "regular",
				X:        22,
				Y:        11,
			},
			{
				Category: "regular",
				Type:     "regular",
				X:        0,
				Y:        12,
			},
			{
				Category: "regular",
				Type:     "regular",
				X:        1,
				Y:        12,
			},
			{
				Category: "regular",
				Type:     "regular",
				X:        2,
				Y:        12,
			},
			{
				Category: "regular",
				Type:     "regular",
				X:        3,
				Y:        12,
			},
			{
				Category: "regular",
				Type:     "regular",
				X:        4,
				Y:        12,
			},
			{
				Category: "regular",
				Type:     "regular",
				X:        5,
				Y:        12,
			},
			{
				Category: "regular",
				Type:     "regular",
				X:        6,
				Y:        12,
			},
			{
				Category: "regular",
				Type:     "regular",
				X:        7,
				Y:        12,
			},
			{
				Category: "regular",
				Type:     "regular",
				X:        8,
				Y:        12,
			},
			{
				Category: "regular",
				Type:     "regular",
				X:        9,
				Y:        12,
			},
			{
				Category: "regular",
				Type:     "regular",
				X:        10,
				Y:        12,
			},
			{
				Category: "regular",
				Type:     "regular",
				X:        11,
				Y:        12,
			},
			{
				Category: "regular",
				Type:     "regular",
				X:        12,
				Y:        12,
			},
			{
				Category: "regular",
				Type:     "regular",
				X:        13,
				Y:        12,
			},
			{
				Category: "regular",
				Type:     "regular",
				X:        14,
				Y:        12,
			},
			{
				Category: "regular",
				Type:     "regular",
				X:        15,
				Y:        12,
			},
			{
				Category: "regular",
				Type:     "regular",
				X:        16,
				Y:        12,
			},
			{
				Category: "regular",
				Type:     "regular",
				X:        17,
				Y:        12,
			},
			{
				Category: "regular",
				Type:     "regular",
				X:        18,
				Y:        12,
			},
			{
				Category: "regular",
				Type:     "regular",
				X:        19,
				Y:        12,
			},
			{
				Category: "regular",
				Type:     "regular",
				X:        20,
				Y:        12,
			},
			{
				Category: "regular",
				Type:     "regular",
				X:        21,
				Y:        12,
			},
			{
				Category: "regular",
				Type:     "regular",
				X:        22,
				Y:        12,
			},
			{
				Category: "regular",
				Type:     "regular",
				X:        0,
				Y:        13,
			},
			{
				Category: "regular",
				Type:     "regular",
				X:        1,
				Y:        13,
			},
			{
				Category: "regular",
				Type:     "regular",
				X:        2,
				Y:        13,
			},
			{
				Category: "regular",
				Type:     "regular",
				X:        3,
				Y:        13,
			},
			{
				Category: "regular",
				Type:     "regular",
				X:        4,
				Y:        13,
			},
			{
				Category: "regular",
				Type:     "regular",
				X:        5,
				Y:        13,
			},
			{
				Category: "regular",
				Type:     "regular",
				X:        6,
				Y:        13,
			},
			{
				Category: "regular",
				Type:     "regular",
				X:        7,
				Y:        13,
			},
			{
				Category: "regular",
				Type:     "regular",
				X:        8,
				Y:        13,
			},
			{
				Category: "regular",
				Type:     "regular",
				X:        9,
				Y:        13,
			},
			{
				Category: "regular",
				Type:     "regular",
				X:        10,
				Y:        13,
			},
			{
				Category: "regular",
				Type:     "regular",
				X:        11,
				Y:        13,
			},
			{
				Category: "regular",
				Type:     "regular",
				X:        12,
				Y:        13,
			},
			{
				Category: "regular",
				Type:     "regular",
				X:        13,
				Y:        13,
			},
			{
				Category: "regular",
				Type:     "regular",
				X:        14,
				Y:        13,
			},
			{
				Category: "regular",
				Type:     "regular",
				X:        15,
				Y:        13,
			},
			{
				Category: "regular",
				Type:     "regular",
				X:        16,
				Y:        13,
			},
			{
				Category: "regular",
				Type:     "regular",
				X:        17,
				Y:        13,
			},
			{
				Category: "regular",
				Type:     "regular",
				X:        18,
				Y:        13,
			},
			{
				Category: "regular",
				Type:     "regular",
				X:        19,
				Y:        13,
			},
			{
				Category: "regular",
				Type:     "regular",
				X:        20,
				Y:        13,
			},
			{
				Category: "regular",
				Type:     "regular",
				X:        21,
				Y:        13,
			},
			{
				Category: "regular",
				Type:     "regular",
				X:        22,
				Y:        13,
			},
			{
				Category: "regular",
				Type:     "regular",
				X:        0,
				Y:        14,
			},
			{
				Category: "regular",
				Type:     "double",
				X:        1,
				Y:        14,
			},
			{
				Category: "regular",
				Type:     "double",
				X:        3,
				Y:        14,
			},
			{
				Category: "regular",
				Type:     "double",
				X:        5,
				Y:        14,
			},
			{
				Category: "regular",
				Type:     "double",
				X:        7,
				Y:        14,
			},
			{
				Category: "regular",
				Type:     "double",
				X:        9,
				Y:        14,
			},
			{
				Category: "regular",
				Type:     "double",
				X:        11,
				Y:        14,
			},
			{
				Category: "regular",
				Type:     "double",
				X:        13,
				Y:        14,
			},
			{
				Category: "regular",
				Type:     "double",
				X:        15,
				Y:        14,
			},
			{
				Category: "regular",
				Type:     "double",
				X:        17,
				Y:        14,
			},
			{
				Category: "regular",
				Type:     "double",
				X:        19,
				Y:        14,
			},
			{
				Category: "regular",
				Type:     "double",
				X:        21,
				Y:        14,
			},
		},

		TheatreId: &theatreId,
	}
}

func GetSampleCreateCinemaHallRequestInvalidWidth() models.CreateCinemaHallRequest {
	validHall := GetSampleCreateCinemaHallRequest()
	validHall.Width = 0
	return validHall
}

func GetSampleCreateCinemaHallRequestInvalidSeatCategory() models.CreateCinemaHallRequest {
	validHall := GetSampleCreateCinemaHallRequest()
	validHall.Seats[0].Category = "invalid"
	return validHall
}

func GetSampleCreateCinemaHallRequestEmptySeats() models.CreateCinemaHallRequest {
	validHall := GetSampleCreateCinemaHallRequest()
	validHall.Seats = []struct {
		X        int
		Y        int
		Type     string
		Category string
	}{}
	return validHall
}

func GetSampleCreateCinemaHallRequestInvalidCoordinates() models.CreateCinemaHallRequest {
	validHall := GetSampleCreateCinemaHallRequest()
	validHall.Seats[0].X = validHall.Width
	return validHall
}

func GetSampleCreateCinemaHallRequestRowsNotAscending() models.CreateCinemaHallRequest {
	validHall := GetSampleCreateCinemaHallRequest()
	temp := validHall.Seats[0]
	validHall.Seats[0] = validHall.Seats[23]
	validHall.Seats[23] = temp
	return validHall
}

func GetSampleCreateCinemaHallRequestDoubleNoSpace() models.CreateCinemaHallRequest {
	validHall := GetSampleCreateCinemaHallRequest()
	validHall.Seats = append(validHall.Seats, struct {
		X        int
		Y        int
		Type     string
		Category string
	}{Category: "regular",
		Type: "regular",
		X:    22,
		Y:    14,
	})
	return validHall
}

func GetSampleCreateCinemaHallRequestDoubleAtEnd() models.CreateCinemaHallRequest {
	validHall := GetSampleCreateCinemaHallRequest()
	validHall.Seats[22].Type = "double"
	return validHall
}

func GetSampleCreateCinemaHallRequestSeatsNotAscending() models.CreateCinemaHallRequest {
	validHall := GetSampleCreateCinemaHallRequest()
	temp := validHall.Seats[0]
	validHall.Seats[0] = validHall.Seats[1]
	validHall.Seats[1] = temp
	return validHall
}

func GetSampleCreateCinemaHallRequestLastSeatsOutOfBounds() models.CreateCinemaHallRequest {
	validHall := GetSampleCreateCinemaHallRequest()
	validHall.Seats = append(validHall.Seats, struct {
		X        int
		Y        int
		Type     string
		Category string
	}{
		Category: "regular",
		Type:     "regular",
		X:        23,
		Y:        14,
	})
	return validHall
}

func GetSampleCreateCinemaHallRequestLastRowOutOfBounds() models.CreateCinemaHallRequest {
	validHall := GetSampleCreateCinemaHallRequest()
	validHall.Seats = append(validHall.Seats,
		struct {
			X        int
			Y        int
			Type     string
			Category string
		}{
			Category: "regular",
			Type:     "regular",
			X:        0,
			Y:        15,
		},
	)
	return validHall
}

func GetSampleCinemaHall() model.CinemaHalls {
	id := uuid.MustParse(cinemaHallId)
	theatreId := uuid.MustParse(theatreId)
	return model.CinemaHalls{
		ID:        &id,
		Name:      "HallName",
		Capacity:  23 * 14,
		TheatreID: &theatreId,
		Width:     23,
		Height:    15,
	}
}

func GetSampleCinemaHalls() []model.CinemaHalls {
	theatreId := uuid.MustParse(theatreId)
	return []model.CinemaHalls{
		{
			ID:        utils.NewUUID(),
			TheatreID: &theatreId,
			Name:      "Hall 1",
			Capacity:  100,
			Width:    10,
			Height:   10,
		},
		{
			ID:        utils.NewUUID(),
			TheatreID: &theatreId,
			Name:      "Hall 2",
			Capacity:  200,
			Width:   20,
			Height:  10,
		},
	}
}

func GetSampleSeat() model.Seats {
	id := uuid.MustParse(seatId)
	cinemaHallId := uuid.MustParse(cinemaHallId)
	seatCategoryId := uuid.MustParse(seatCategoryId1)
	return model.Seats{
		ID:             &id,
		X:              1,
		Y:              1,
		ColumnNr:       1,
		RowNr:          1,
		CinemaHallID:   &cinemaHallId,
		SeatCategoryID: &seatCategoryId,
	}
}

func GetSampleSeatCategories() []model.SeatCategories {
	seatCategoryId1 := uuid.MustParse(seatCategoryId1)
	seatCategoryId2 := uuid.MustParse(seatCategoryId2)
	seatCategoryId3 := uuid.MustParse(seatCategoryId3)
	return []model.SeatCategories{
		{
			ID:           &seatCategoryId1,
			CategoryName: "regular",
		},
		{
			ID:           &seatCategoryId2,
			CategoryName: "loge",
		},
		{
			ID:           &seatCategoryId3,
			CategoryName: "vip",
		},
	}
}
