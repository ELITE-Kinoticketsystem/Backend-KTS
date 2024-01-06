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
		Seats: [][]struct {
			RowNr    int
			ColumnNr int
			Type     string
			Category string
		}{
			{
				{
					Category: "regular",
					Type:     "regular",
					ColumnNr: 0,
					RowNr:    0,
				},
				{
					Category: "regular",
					Type:     "regular",
					ColumnNr: 1,
					RowNr:    0,
				},
				{
					Category: "regular",
					Type:     "regular",
					ColumnNr: 2,
					RowNr:    0,
				},
				{
					Category: "regular",
					Type:     "regular",
					ColumnNr: 3,
					RowNr:    0,
				},
				{
					Category: "regular",
					Type:     "regular",
					ColumnNr: 4,
					RowNr:    0,
				},
				{
					Category: "regular",
					Type:     "regular",
					ColumnNr: 5,
					RowNr:    0,
				},
				{
					Category: "regular",
					Type:     "regular",
					ColumnNr: 6,
					RowNr:    0,
				},
				{
					Category: "regular",
					Type:     "regular",
					ColumnNr: 7,
					RowNr:    0,
				},
				{
					Category: "regular",
					Type:     "regular",
					ColumnNr: 8,
					RowNr:    0,
				},
				{
					Category: "regular",
					Type:     "regular",
					ColumnNr: 9,
					RowNr:    0,
				},
				{
					Category: "regular",
					Type:     "regular",
					ColumnNr: 10,
					RowNr:    0,
				},
				{
					Category: "regular",
					Type:     "regular",
					ColumnNr: 11,
					RowNr:    0,
				},
				{
					Category: "regular",
					Type:     "regular",
					ColumnNr: 12,
					RowNr:    0,
				},
				{
					Category: "regular",
					Type:     "regular",
					ColumnNr: 13,
					RowNr:    0,
				},
				{
					Category: "regular",
					Type:     "regular",
					ColumnNr: 14,
					RowNr:    0,
				},
				{
					Category: "regular",
					Type:     "regular",
					ColumnNr: 15,
					RowNr:    0,
				},
				{
					Category: "regular",
					Type:     "regular",
					ColumnNr: 16,
					RowNr:    0,
				},
				{
					Category: "regular",
					Type:     "regular",
					ColumnNr: 17,
					RowNr:    0,
				},
				{
					Category: "regular",
					Type:     "regular",
					ColumnNr: 18,
					RowNr:    0,
				},
				{
					Category: "regular",
					Type:     "regular",
					ColumnNr: 19,
					RowNr:    0,
				},
				{
					Category: "regular",
					Type:     "regular",
					ColumnNr: 20,
					RowNr:    0,
				},
				{
					Category: "regular",
					Type:     "regular",
					ColumnNr: 21,
					RowNr:    0,
				},
				{
					Category: "regular",
					Type:     "regular",
					ColumnNr: 22,
					RowNr:    0,
				},
			},
			{

				{
					Category: "regular",
					Type:     "regular",
					ColumnNr: 0,
					RowNr:    1,
				},
				{
					Category: "regular",
					Type:     "regular",
					ColumnNr: 1,
					RowNr:    1,
				},
				{
					Category: "regular",
					Type:     "regular",
					ColumnNr: 2,
					RowNr:    1,
				},
				{
					Category: "regular",
					Type:     "regular",
					ColumnNr: 3,
					RowNr:    1,
				},
				{
					Category: "regular",
					Type:     "regular",
					ColumnNr: 4,
					RowNr:    1,
				},
				{
					Category: "regular",
					Type:     "regular",
					ColumnNr: 5,
					RowNr:    1,
				},
				{
					Category: "regular",
					Type:     "regular",
					ColumnNr: 6,
					RowNr:    1,
				},
				{
					Category: "regular",
					Type:     "regular",
					ColumnNr: 7,
					RowNr:    1,
				},
				{
					Category: "regular",
					Type:     "regular",
					ColumnNr: 8,
					RowNr:    1,
				},
				{
					Category: "regular",
					Type:     "regular",
					ColumnNr: 9,
					RowNr:    1,
				},
				{
					Category: "regular",
					Type:     "regular",
					ColumnNr: 10,
					RowNr:    1,
				},
				{
					Category: "regular",
					Type:     "regular",
					ColumnNr: 11,
					RowNr:    1,
				},
				{
					Category: "regular",
					Type:     "regular",
					ColumnNr: 12,
					RowNr:    1,
				},
				{
					Category: "regular",
					Type:     "regular",
					ColumnNr: 13,
					RowNr:    1,
				},
				{
					Category: "regular",
					Type:     "regular",
					ColumnNr: 14,
					RowNr:    1,
				},
				{
					Category: "regular",
					Type:     "regular",
					ColumnNr: 15,
					RowNr:    1,
				},
				{
					Category: "regular",
					Type:     "regular",
					ColumnNr: 16,
					RowNr:    1,
				},
				{
					Category: "regular",
					Type:     "regular",
					ColumnNr: 17,
					RowNr:    1,
				},
				{
					Category: "regular",
					Type:     "regular",
					ColumnNr: 18,
					RowNr:    1,
				},
				{
					Category: "regular",
					Type:     "regular",
					ColumnNr: 19,
					RowNr:    1,
				},
				{
					Category: "regular",
					Type:     "regular",
					ColumnNr: 20,
					RowNr:    1,
				},
				{
					Category: "regular",
					Type:     "regular",
					ColumnNr: 21,
					RowNr:    1,
				},
				{
					Category: "regular",
					Type:     "regular",
					ColumnNr: 22,
					RowNr:    1,
				},
			},
			{
				{
					Category: "regular",
					Type:     "regular",
					ColumnNr: 0,
					RowNr:    2,
				},
				{
					Category: "regular",
					Type:     "regular",
					ColumnNr: 1,
					RowNr:    2,
				},
				{
					Category: "regular",
					Type:     "regular",
					ColumnNr: 2,
					RowNr:    2,
				},
				{
					Category: "regular",
					Type:     "regular",
					ColumnNr: 3,
					RowNr:    2,
				},
				{
					Category: "regular",
					Type:     "regular",
					ColumnNr: 4,
					RowNr:    2,
				},
				{
					Category: "regular",
					Type:     "regular",
					ColumnNr: 5,
					RowNr:    2,
				},
				{
					Category: "regular",
					Type:     "regular",
					ColumnNr: 6,
					RowNr:    2,
				},
				{
					Category: "regular",
					Type:     "regular",
					ColumnNr: 7,
					RowNr:    2,
				},
				{
					Category: "regular",
					Type:     "regular",
					ColumnNr: 8,
					RowNr:    2,
				},
				{
					Category: "regular",
					Type:     "regular",
					ColumnNr: 9,
					RowNr:    2,
				},
				{
					Category: "regular",
					Type:     "regular",
					ColumnNr: 10,
					RowNr:    2,
				},
				{
					Category: "regular",
					Type:     "regular",
					ColumnNr: 11,
					RowNr:    2,
				},
				{
					Category: "regular",
					Type:     "regular",
					ColumnNr: 12,
					RowNr:    2,
				},
				{
					Category: "regular",
					Type:     "regular",
					ColumnNr: 13,
					RowNr:    2,
				},
				{
					Category: "regular",
					Type:     "regular",
					ColumnNr: 14,
					RowNr:    2,
				},
				{
					Category: "regular",
					Type:     "regular",
					ColumnNr: 15,
					RowNr:    2,
				},
				{
					Category: "regular",
					Type:     "regular",
					ColumnNr: 16,
					RowNr:    2,
				},
				{
					Category: "regular",
					Type:     "regular",
					ColumnNr: 17,
					RowNr:    2,
				},
				{
					Category: "regular",
					Type:     "regular",
					ColumnNr: 18,
					RowNr:    2,
				},
				{
					Category: "regular",
					Type:     "regular",
					ColumnNr: 19,
					RowNr:    2,
				},
				{
					Category: "regular",
					Type:     "regular",
					ColumnNr: 20,
					RowNr:    2,
				},
				{
					Category: "regular",
					Type:     "regular",
					ColumnNr: 21,
					RowNr:    2,
				},
				{
					Category: "regular",
					Type:     "regular",
					ColumnNr: 22,
					RowNr:    2,
				},
			},
			{
				{
					Category: "regular",
					Type:     "regular",
					ColumnNr: 0,
					RowNr:    3,
				},
				{
					Category: "regular",
					Type:     "regular",
					ColumnNr: 1,
					RowNr:    3,
				},
				{
					Category: "regular",
					Type:     "regular",
					ColumnNr: 2,
					RowNr:    3,
				},
				{
					Category: "regular",
					Type:     "regular",
					ColumnNr: 3,
					RowNr:    3,
				},
				{
					Category: "regular",
					Type:     "regular",
					ColumnNr: 4,
					RowNr:    3,
				},
				{
					Category: "regular",
					Type:     "regular",
					ColumnNr: 5,
					RowNr:    3,
				},
				{
					Category: "regular",
					Type:     "regular",
					ColumnNr: 6,
					RowNr:    3,
				},
				{
					Category: "regular",
					Type:     "regular",
					ColumnNr: 7,
					RowNr:    3,
				},
				{
					Category: "regular",
					Type:     "regular",
					ColumnNr: 8,
					RowNr:    3,
				},
				{
					Category: "regular",
					Type:     "regular",
					ColumnNr: 9,
					RowNr:    3,
				},
				{
					Category: "regular",
					Type:     "regular",
					ColumnNr: 10,
					RowNr:    3,
				},
				{
					Category: "regular",
					Type:     "regular",
					ColumnNr: 11,
					RowNr:    3,
				},
				{
					Category: "regular",
					Type:     "regular",
					ColumnNr: 12,
					RowNr:    3,
				},
				{
					Category: "regular",
					Type:     "regular",
					ColumnNr: 13,
					RowNr:    3,
				},
				{
					Category: "regular",
					Type:     "regular",
					ColumnNr: 14,
					RowNr:    3,
				},
				{
					Category: "regular",
					Type:     "regular",
					ColumnNr: 15,
					RowNr:    3,
				},
				{
					Category: "regular",
					Type:     "regular",
					ColumnNr: 16,
					RowNr:    3,
				},
				{
					Category: "regular",
					Type:     "regular",
					ColumnNr: 17,
					RowNr:    3,
				},
				{
					Category: "regular",
					Type:     "regular",
					ColumnNr: 18,
					RowNr:    3,
				},
				{
					Category: "regular",
					Type:     "regular",
					ColumnNr: 19,
					RowNr:    3,
				},
				{
					Category: "regular",
					Type:     "regular",
					ColumnNr: 20,
					RowNr:    3,
				},
				{
					Category: "regular",
					Type:     "regular",
					ColumnNr: 21,
					RowNr:    3,
				},
				{
					Category: "regular",
					Type:     "regular",
					ColumnNr: 22,
					RowNr:    3,
				},
			},
			{
				{
					Category: "regular",
					Type:     "regular",
					ColumnNr: 0,
					RowNr:    4,
				},
				{
					Category: "regular",
					Type:     "regular",
					ColumnNr: 1,
					RowNr:    4,
				},
				{
					Category: "regular",
					Type:     "regular",
					ColumnNr: 2,
					RowNr:    4,
				},
				{
					Category: "regular",
					Type:     "regular",
					ColumnNr: 3,
					RowNr:    4,
				},
				{
					Category: "regular",
					Type:     "regular",
					ColumnNr: 4,
					RowNr:    4,
				},
				{
					Category: "regular",
					Type:     "regular",
					ColumnNr: 5,
					RowNr:    4,
				},
				{
					Category: "regular",
					Type:     "regular",
					ColumnNr: 6,
					RowNr:    4,
				},
				{
					Category: "regular",
					Type:     "regular",
					ColumnNr: 7,
					RowNr:    4,
				},
				{
					Category: "regular",
					Type:     "regular",
					ColumnNr: 8,
					RowNr:    4,
				},
				{
					Category: "regular",
					Type:     "regular",
					ColumnNr: 9,
					RowNr:    4,
				},
				{
					Category: "regular",
					Type:     "regular",
					ColumnNr: 10,
					RowNr:    4,
				},
				{
					Category: "regular",
					Type:     "regular",
					ColumnNr: 11,
					RowNr:    4,
				},
				{
					Category: "regular",
					Type:     "regular",
					ColumnNr: 12,
					RowNr:    4,
				},
				{
					Category: "regular",
					Type:     "regular",
					ColumnNr: 13,
					RowNr:    4,
				},
				{
					Category: "regular",
					Type:     "regular",
					ColumnNr: 14,
					RowNr:    4,
				},
				{
					Category: "regular",
					Type:     "regular",
					ColumnNr: 15,
					RowNr:    4,
				},
				{
					Category: "regular",
					Type:     "regular",
					ColumnNr: 16,
					RowNr:    4,
				},
				{
					Category: "regular",
					Type:     "regular",
					ColumnNr: 17,
					RowNr:    4,
				},
				{
					Category: "regular",
					Type:     "regular",
					ColumnNr: 18,
					RowNr:    4,
				},
				{
					Category: "regular",
					Type:     "regular",
					ColumnNr: 19,
					RowNr:    4,
				},
				{
					Category: "regular",
					Type:     "regular",
					ColumnNr: 20,
					RowNr:    4,
				},
				{
					Category: "regular",
					Type:     "regular",
					ColumnNr: 21,
					RowNr:    4,
				},
				{
					Category: "regular",
					Type:     "regular",
					ColumnNr: 22,
					RowNr:    4,
				},
			},
			{
				{
					Category: "regular",
					Type:     "regular",
					ColumnNr: 0,
					RowNr:    5,
				},
				{
					Category: "regular",
					Type:     "regular",
					ColumnNr: 1,
					RowNr:    5,
				},
				{
					Category: "regular",
					Type:     "regular",
					ColumnNr: 2,
					RowNr:    5,
				},
				{
					Category: "regular",
					Type:     "regular",
					ColumnNr: 3,
					RowNr:    5,
				},
				{
					Category: "regular",
					Type:     "regular",
					ColumnNr: 4,
					RowNr:    5,
				},
				{
					Category: "regular",
					Type:     "regular",
					ColumnNr: 5,
					RowNr:    5,
				},
				{
					Category: "regular",
					Type:     "regular",
					ColumnNr: 6,
					RowNr:    5,
				},
				{
					Category: "regular",
					Type:     "regular",
					ColumnNr: 7,
					RowNr:    5,
				},
				{
					Category: "regular",
					Type:     "regular",
					ColumnNr: 8,
					RowNr:    5,
				},
				{
					Category: "regular",
					Type:     "regular",
					ColumnNr: 9,
					RowNr:    5,
				},
				{
					Category: "regular",
					Type:     "regular",
					ColumnNr: 10,
					RowNr:    5,
				},
				{
					Category: "regular",
					Type:     "regular",
					ColumnNr: 11,
					RowNr:    5,
				},
				{
					Category: "regular",
					Type:     "regular",
					ColumnNr: 12,
					RowNr:    5,
				},
				{
					Category: "regular",
					Type:     "regular",
					ColumnNr: 13,
					RowNr:    5,
				},
				{
					Category: "regular",
					Type:     "regular",
					ColumnNr: 14,
					RowNr:    5,
				},
				{
					Category: "regular",
					Type:     "regular",
					ColumnNr: 15,
					RowNr:    5,
				},
				{
					Category: "regular",
					Type:     "regular",
					ColumnNr: 16,
					RowNr:    5,
				},
				{
					Category: "regular",
					Type:     "regular",
					ColumnNr: 17,
					RowNr:    5,
				},
				{
					Category: "regular",
					Type:     "regular",
					ColumnNr: 18,
					RowNr:    5,
				},
				{
					Category: "regular",
					Type:     "regular",
					ColumnNr: 19,
					RowNr:    5,
				},
				{
					Category: "regular",
					Type:     "regular",
					ColumnNr: 20,
					RowNr:    5,
				},
				{
					Category: "regular",
					Type:     "regular",
					ColumnNr: 21,
					RowNr:    5,
				},
				{
					Category: "regular",
					Type:     "regular",
					ColumnNr: 22,
					RowNr:    5,
				},
			},
			{
				{
					Category: "regular",
					Type:     "regular",
					ColumnNr: 0,
					RowNr:    6,
				},
				{
					Category: "regular",
					Type:     "regular",
					ColumnNr: 1,
					RowNr:    6,
				},
				{
					Category: "regular",
					Type:     "regular",
					ColumnNr: 2,
					RowNr:    6,
				},
				{
					Category: "regular",
					Type:     "regular",
					ColumnNr: 3,
					RowNr:    6,
				},
				{
					Category: "regular",
					Type:     "regular",
					ColumnNr: 4,
					RowNr:    6,
				},
				{
					Category: "regular",
					Type:     "regular",
					ColumnNr: 5,
					RowNr:    6,
				},
				{
					Category: "regular",
					Type:     "regular",
					ColumnNr: 6,
					RowNr:    6,
				},
				{
					Category: "regular",
					Type:     "regular",
					ColumnNr: 7,
					RowNr:    6,
				},
				{
					Category: "regular",
					Type:     "regular",
					ColumnNr: 8,
					RowNr:    6,
				},
				{
					Category: "regular",
					Type:     "regular",
					ColumnNr: 9,
					RowNr:    6,
				},
				{
					Category: "regular",
					Type:     "regular",
					ColumnNr: 10,
					RowNr:    6,
				},
				{
					Category: "regular",
					Type:     "regular",
					ColumnNr: 11,
					RowNr:    6,
				},
				{
					Category: "regular",
					Type:     "regular",
					ColumnNr: 12,
					RowNr:    6,
				},
				{
					Category: "regular",
					Type:     "regular",
					ColumnNr: 13,
					RowNr:    6,
				},
				{
					Category: "regular",
					Type:     "regular",
					ColumnNr: 14,
					RowNr:    6,
				},
				{
					Category: "regular",
					Type:     "regular",
					ColumnNr: 15,
					RowNr:    6,
				},
				{
					Category: "regular",
					Type:     "regular",
					ColumnNr: 16,
					RowNr:    6,
				},
				{
					Category: "regular",
					Type:     "regular",
					ColumnNr: 17,
					RowNr:    6,
				},
				{
					Category: "regular",
					Type:     "regular",
					ColumnNr: 18,
					RowNr:    6,
				},
				{
					Category: "regular",
					Type:     "regular",
					ColumnNr: 19,
					RowNr:    6,
				},
				{
					Category: "regular",
					Type:     "regular",
					ColumnNr: 20,
					RowNr:    6,
				},
				{
					Category: "regular",
					Type:     "regular",
					ColumnNr: 21,
					RowNr:    6,
				},
				{
					Category: "regular",
					Type:     "regular",
					ColumnNr: 22,
					RowNr:    6,
				},
			},
			{
				{
					Category: "regular",
					Type:     "regular",
					ColumnNr: 0,
					RowNr:    7,
				},
				{
					Category: "regular",
					Type:     "regular",
					ColumnNr: 1,
					RowNr:    7,
				},
				{
					Category: "regular",
					Type:     "regular",
					ColumnNr: 2,
					RowNr:    7,
				},
				{
					Category: "regular",
					Type:     "regular",
					ColumnNr: 3,
					RowNr:    7,
				},
				{
					Category: "regular",
					Type:     "regular",
					ColumnNr: 4,
					RowNr:    7,
				},
				{
					Category: "regular",
					Type:     "regular",
					ColumnNr: 5,
					RowNr:    7,
				},
				{
					Category: "regular",
					Type:     "regular",
					ColumnNr: 6,
					RowNr:    7,
				},
				{
					Category: "regular",
					Type:     "regular",
					ColumnNr: 7,
					RowNr:    7,
				},
				{
					Category: "regular",
					Type:     "regular",
					ColumnNr: 8,
					RowNr:    7,
				},
				{
					Category: "regular",
					Type:     "regular",
					ColumnNr: 9,
					RowNr:    7,
				},
				{
					Category: "regular",
					Type:     "regular",
					ColumnNr: 10,
					RowNr:    7,
				},
				{
					Category: "regular",
					Type:     "regular",
					ColumnNr: 11,
					RowNr:    7,
				},
				{
					Category: "regular",
					Type:     "regular",
					ColumnNr: 12,
					RowNr:    7,
				},
				{
					Category: "regular",
					Type:     "regular",
					ColumnNr: 13,
					RowNr:    7,
				},
				{
					Category: "regular",
					Type:     "regular",
					ColumnNr: 14,
					RowNr:    7,
				},
				{
					Category: "regular",
					Type:     "regular",
					ColumnNr: 15,
					RowNr:    7,
				},
				{
					Category: "regular",
					Type:     "regular",
					ColumnNr: 16,
					RowNr:    7,
				},
				{
					Category: "regular",
					Type:     "regular",
					ColumnNr: 17,
					RowNr:    7,
				},
				{
					Category: "regular",
					Type:     "regular",
					ColumnNr: 18,
					RowNr:    7,
				},
				{
					Category: "regular",
					Type:     "regular",
					ColumnNr: 19,
					RowNr:    7,
				},
				{
					Category: "regular",
					Type:     "regular",
					ColumnNr: 20,
					RowNr:    7,
				},
				{
					Category: "regular",
					Type:     "regular",
					ColumnNr: 21,
					RowNr:    7,
				},
				{
					Category: "regular",
					Type:     "regular",
					ColumnNr: 22,
					RowNr:    7,
				},
			},
			{
				{
					Category: "regular",
					Type:     "regular",
					ColumnNr: 0,
					RowNr:    8,
				},
				{
					Category: "regular",
					Type:     "regular",
					ColumnNr: 1,
					RowNr:    8,
				},
				{
					Category: "regular",
					Type:     "regular",
					ColumnNr: 2,
					RowNr:    8,
				},
				{
					Category: "regular",
					Type:     "regular",
					ColumnNr: 3,
					RowNr:    8,
				},
				{
					Category: "regular",
					Type:     "regular",
					ColumnNr: 4,
					RowNr:    8,
				},
				{
					Category: "regular",
					Type:     "regular",
					ColumnNr: 5,
					RowNr:    8,
				},
				{
					Category: "regular",
					Type:     "regular",
					ColumnNr: 6,
					RowNr:    8,
				},
				{
					Category: "regular",
					Type:     "regular",
					ColumnNr: 7,
					RowNr:    8,
				},
				{
					Category: "regular",
					Type:     "regular",
					ColumnNr: 8,
					RowNr:    8,
				},
				{
					Category: "regular",
					Type:     "regular",
					ColumnNr: 9,
					RowNr:    8,
				},
				{
					Category: "regular",
					Type:     "regular",
					ColumnNr: 10,
					RowNr:    8,
				},
				{
					Category: "regular",
					Type:     "regular",
					ColumnNr: 11,
					RowNr:    8,
				},
				{
					Category: "regular",
					Type:     "regular",
					ColumnNr: 12,
					RowNr:    8,
				},
				{
					Category: "regular",
					Type:     "regular",
					ColumnNr: 13,
					RowNr:    8,
				},
				{
					Category: "regular",
					Type:     "regular",
					ColumnNr: 14,
					RowNr:    8,
				},
				{
					Category: "regular",
					Type:     "regular",
					ColumnNr: 15,
					RowNr:    8,
				},
				{
					Category: "regular",
					Type:     "regular",
					ColumnNr: 16,
					RowNr:    8,
				},
				{
					Category: "regular",
					Type:     "regular",
					ColumnNr: 17,
					RowNr:    8,
				},
				{
					Category: "regular",
					Type:     "regular",
					ColumnNr: 18,
					RowNr:    8,
				},
				{
					Category: "regular",
					Type:     "regular",
					ColumnNr: 19,
					RowNr:    8,
				},
				{
					Category: "regular",
					Type:     "regular",
					ColumnNr: 20,
					RowNr:    8,
				},
				{
					Category: "regular",
					Type:     "regular",
					ColumnNr: 21,
					RowNr:    8,
				},
				{
					Category: "regular",
					Type:     "regular",
					ColumnNr: 22,
					RowNr:    8,
				},
			},
			{
				{
					Category: "regular",
					Type:     "regular",
					ColumnNr: 0,
					RowNr:    9,
				},
				{
					Category: "regular",
					Type:     "regular",
					ColumnNr: 1,
					RowNr:    9,
				},
				{
					Category: "regular",
					Type:     "regular",
					ColumnNr: 2,
					RowNr:    9,
				},
				{
					Category: "regular",
					Type:     "regular",
					ColumnNr: 3,
					RowNr:    9,
				},
				{
					Category: "regular",
					Type:     "regular",
					ColumnNr: 4,
					RowNr:    9,
				},
				{
					Category: "regular",
					Type:     "regular",
					ColumnNr: 5,
					RowNr:    9,
				},
				{
					Category: "regular",
					Type:     "regular",
					ColumnNr: 6,
					RowNr:    9,
				},
				{
					Category: "regular",
					Type:     "regular",
					ColumnNr: 7,
					RowNr:    9,
				},
				{
					Category: "regular",
					Type:     "regular",
					ColumnNr: 8,
					RowNr:    9,
				},
				{
					Category: "regular",
					Type:     "regular",
					ColumnNr: 9,
					RowNr:    9,
				},
				{
					Category: "regular",
					Type:     "regular",
					ColumnNr: 10,
					RowNr:    9,
				},
				{
					Category: "regular",
					Type:     "regular",
					ColumnNr: 11,
					RowNr:    9,
				},
				{
					Category: "regular",
					Type:     "regular",
					ColumnNr: 12,
					RowNr:    9,
				},
				{
					Category: "regular",
					Type:     "regular",
					ColumnNr: 13,
					RowNr:    9,
				},
				{
					Category: "regular",
					Type:     "regular",
					ColumnNr: 14,
					RowNr:    9,
				},
				{
					Category: "regular",
					Type:     "regular",
					ColumnNr: 15,
					RowNr:    9,
				},
				{
					Category: "regular",
					Type:     "regular",
					ColumnNr: 16,
					RowNr:    9,
				},
				{
					Category: "regular",
					Type:     "regular",
					ColumnNr: 17,
					RowNr:    9,
				},
				{
					Category: "regular",
					Type:     "regular",
					ColumnNr: 18,
					RowNr:    9,
				},
				{
					Category: "regular",
					Type:     "regular",
					ColumnNr: 19,
					RowNr:    9,
				},
				{
					Category: "regular",
					Type:     "regular",
					ColumnNr: 20,
					RowNr:    9,
				},
				{
					Category: "regular",
					Type:     "regular",
					ColumnNr: 21,
					RowNr:    9,
				},
				{
					Category: "regular",
					Type:     "regular",
					ColumnNr: 22,
					RowNr:    9,
				},
			},
			{
				{
					Category: "regular",
					Type:     "regular",
					ColumnNr: 0,
					RowNr:    10,
				},
				{
					Category: "regular",
					Type:     "regular",
					ColumnNr: 1,
					RowNr:    10,
				},
				{
					Category: "regular",
					Type:     "regular",
					ColumnNr: 2,
					RowNr:    10,
				},
				{
					Category: "regular",
					Type:     "regular",
					ColumnNr: 3,
					RowNr:    10,
				},
				{
					Category: "regular",
					Type:     "regular",
					ColumnNr: 4,
					RowNr:    10,
				},
				{
					Category: "regular",
					Type:     "regular",
					ColumnNr: 5,
					RowNr:    10,
				},
				{
					Category: "regular",
					Type:     "regular",
					ColumnNr: 6,
					RowNr:    10,
				},
				{
					Category: "regular",
					Type:     "regular",
					ColumnNr: 7,
					RowNr:    10,
				},
				{
					Category: "regular",
					Type:     "regular",
					ColumnNr: 8,
					RowNr:    10,
				},
				{
					Category: "regular",
					Type:     "regular",
					ColumnNr: 9,
					RowNr:    10,
				},
				{
					Category: "regular",
					Type:     "regular",
					ColumnNr: 10,
					RowNr:    10,
				},
				{
					Category: "regular",
					Type:     "regular",
					ColumnNr: 11,
					RowNr:    10,
				},
				{
					Category: "regular",
					Type:     "regular",
					ColumnNr: 12,
					RowNr:    10,
				},
				{
					Category: "regular",
					Type:     "regular",
					ColumnNr: 13,
					RowNr:    10,
				},
				{
					Category: "regular",
					Type:     "regular",
					ColumnNr: 14,
					RowNr:    10,
				},
				{
					Category: "regular",
					Type:     "regular",
					ColumnNr: 15,
					RowNr:    10,
				},
				{
					Category: "regular",
					Type:     "regular",
					ColumnNr: 16,
					RowNr:    10,
				},
				{
					Category: "regular",
					Type:     "regular",
					ColumnNr: 17,
					RowNr:    10,
				},
				{
					Category: "regular",
					Type:     "regular",
					ColumnNr: 18,
					RowNr:    10,
				},
				{
					Category: "regular",
					Type:     "regular",
					ColumnNr: 19,
					RowNr:    10,
				},
				{
					Category: "regular",
					Type:     "regular",
					ColumnNr: 20,
					RowNr:    10,
				},
				{
					Category: "regular",
					Type:     "regular",
					ColumnNr: 21,
					RowNr:    10,
				},
				{
					Category: "regular",
					Type:     "regular",
					ColumnNr: 22,
					RowNr:    10,
				},
			},
			{
				{
					Category: "regular",
					Type:     "regular",
					ColumnNr: 0,
					RowNr:    11,
				},
				{
					Category: "regular",
					Type:     "regular",
					ColumnNr: 1,
					RowNr:    11,
				},
				{
					Category: "regular",
					Type:     "regular",
					ColumnNr: 2,
					RowNr:    11,
				},
				{
					Category: "regular",
					Type:     "regular",
					ColumnNr: 3,
					RowNr:    11,
				},
				{
					Category: "regular",
					Type:     "regular",
					ColumnNr: 4,
					RowNr:    11,
				},
				{
					Category: "regular",
					Type:     "regular",
					ColumnNr: 5,
					RowNr:    11,
				},
				{
					Category: "regular",
					Type:     "regular",
					ColumnNr: 6,
					RowNr:    11,
				},
				{
					Category: "regular",
					Type:     "regular",
					ColumnNr: 7,
					RowNr:    11,
				},
				{
					Category: "regular",
					Type:     "regular",
					ColumnNr: 8,
					RowNr:    11,
				},
				{
					Category: "regular",
					Type:     "regular",
					ColumnNr: 9,
					RowNr:    11,
				},
				{
					Category: "regular",
					Type:     "regular",
					ColumnNr: 10,
					RowNr:    11,
				},
				{
					Category: "regular",
					Type:     "regular",
					ColumnNr: 11,
					RowNr:    11,
				},
				{
					Category: "regular",
					Type:     "regular",
					ColumnNr: 12,
					RowNr:    11,
				},
				{
					Category: "regular",
					Type:     "regular",
					ColumnNr: 13,
					RowNr:    11,
				},
				{
					Category: "regular",
					Type:     "regular",
					ColumnNr: 14,
					RowNr:    11,
				},
				{
					Category: "regular",
					Type:     "regular",
					ColumnNr: 15,
					RowNr:    11,
				},
				{
					Category: "regular",
					Type:     "regular",
					ColumnNr: 16,
					RowNr:    11,
				},
				{
					Category: "regular",
					Type:     "regular",
					ColumnNr: 17,
					RowNr:    11,
				},
				{
					Category: "regular",
					Type:     "regular",
					ColumnNr: 18,
					RowNr:    11,
				},
				{
					Category: "regular",
					Type:     "regular",
					ColumnNr: 19,
					RowNr:    11,
				},
				{
					Category: "regular",
					Type:     "regular",
					ColumnNr: 20,
					RowNr:    11,
				},
				{
					Category: "regular",
					Type:     "regular",
					ColumnNr: 21,
					RowNr:    11,
				},
				{
					Category: "regular",
					Type:     "regular",
					ColumnNr: 22,
					RowNr:    11,
				},
			},
			{
				{
					Category: "regular",
					Type:     "regular",
					ColumnNr: 0,
					RowNr:    12,
				},
				{
					Category: "regular",
					Type:     "regular",
					ColumnNr: 1,
					RowNr:    12,
				},
				{
					Category: "regular",
					Type:     "regular",
					ColumnNr: 2,
					RowNr:    12,
				},
				{
					Category: "regular",
					Type:     "regular",
					ColumnNr: 3,
					RowNr:    12,
				},
				{
					Category: "regular",
					Type:     "regular",
					ColumnNr: 4,
					RowNr:    12,
				},
				{
					Category: "regular",
					Type:     "regular",
					ColumnNr: 5,
					RowNr:    12,
				},
				{
					Category: "regular",
					Type:     "regular",
					ColumnNr: 6,
					RowNr:    12,
				},
				{
					Category: "regular",
					Type:     "regular",
					ColumnNr: 7,
					RowNr:    12,
				},
				{
					Category: "regular",
					Type:     "regular",
					ColumnNr: 8,
					RowNr:    12,
				},
				{
					Category: "regular",
					Type:     "regular",
					ColumnNr: 9,
					RowNr:    12,
				},
				{
					Category: "regular",
					Type:     "regular",
					ColumnNr: 10,
					RowNr:    12,
				},
				{
					Category: "regular",
					Type:     "regular",
					ColumnNr: 11,
					RowNr:    12,
				},
				{
					Category: "regular",
					Type:     "regular",
					ColumnNr: 12,
					RowNr:    12,
				},
				{
					Category: "regular",
					Type:     "regular",
					ColumnNr: 13,
					RowNr:    12,
				},
				{
					Category: "regular",
					Type:     "regular",
					ColumnNr: 14,
					RowNr:    12,
				},
				{
					Category: "regular",
					Type:     "regular",
					ColumnNr: 15,
					RowNr:    12,
				},
				{
					Category: "regular",
					Type:     "regular",
					ColumnNr: 16,
					RowNr:    12,
				},
				{
					Category: "regular",
					Type:     "regular",
					ColumnNr: 17,
					RowNr:    12,
				},
				{
					Category: "regular",
					Type:     "regular",
					ColumnNr: 18,
					RowNr:    12,
				},
				{
					Category: "regular",
					Type:     "regular",
					ColumnNr: 19,
					RowNr:    12,
				},
				{
					Category: "regular",
					Type:     "regular",
					ColumnNr: 20,
					RowNr:    12,
				},
				{
					Category: "regular",
					Type:     "regular",
					ColumnNr: 21,
					RowNr:    12,
				},
				{
					Category: "regular",
					Type:     "regular",
					ColumnNr: 22,
					RowNr:    12,
				},
			},
			{
				{
					Category: "regular",
					Type:     "regular",
					ColumnNr: 0,
					RowNr:    13,
				},
				{
					Category: "regular",
					Type:     "regular",
					ColumnNr: 1,
					RowNr:    13,
				},
				{
					Category: "regular",
					Type:     "regular",
					ColumnNr: 2,
					RowNr:    13,
				},
				{
					Category: "regular",
					Type:     "regular",
					ColumnNr: 3,
					RowNr:    13,
				},
				{
					Category: "regular",
					Type:     "regular",
					ColumnNr: 4,
					RowNr:    13,
				},
				{
					Category: "regular",
					Type:     "regular",
					ColumnNr: 5,
					RowNr:    13,
				},
				{
					Category: "regular",
					Type:     "regular",
					ColumnNr: 6,
					RowNr:    13,
				},
				{
					Category: "regular",
					Type:     "regular",
					ColumnNr: 7,
					RowNr:    13,
				},
				{
					Category: "regular",
					Type:     "regular",
					ColumnNr: 8,
					RowNr:    13,
				},
				{
					Category: "regular",
					Type:     "regular",
					ColumnNr: 9,
					RowNr:    13,
				},
				{
					Category: "regular",
					Type:     "regular",
					ColumnNr: 10,
					RowNr:    13,
				},
				{
					Category: "regular",
					Type:     "regular",
					ColumnNr: 11,
					RowNr:    13,
				},
				{
					Category: "regular",
					Type:     "regular",
					ColumnNr: 12,
					RowNr:    13,
				},
				{
					Category: "regular",
					Type:     "regular",
					ColumnNr: 13,
					RowNr:    13,
				},
				{
					Category: "regular",
					Type:     "regular",
					ColumnNr: 14,
					RowNr:    13,
				},
				{
					Category: "regular",
					Type:     "regular",
					ColumnNr: 15,
					RowNr:    13,
				},
				{
					Category: "regular",
					Type:     "regular",
					ColumnNr: 16,
					RowNr:    13,
				},
				{
					Category: "regular",
					Type:     "regular",
					ColumnNr: 17,
					RowNr:    13,
				},
				{
					Category: "regular",
					Type:     "regular",
					ColumnNr: 18,
					RowNr:    13,
				},
				{
					Category: "regular",
					Type:     "regular",
					ColumnNr: 19,
					RowNr:    13,
				},
				{
					Category: "regular",
					Type:     "regular",
					ColumnNr: 20,
					RowNr:    13,
				},
				{
					Category: "regular",
					Type:     "regular",
					ColumnNr: 21,
					RowNr:    13,
				},
				{
					Category: "regular",
					Type:     "regular",
					ColumnNr: 22,
					RowNr:    13,
				},
			},
			{
				{
					Category: "regular",
					Type:     "regular",
					ColumnNr: 0,
					RowNr:    14,
				},
				{
					Category: "regular",
					Type:     "double",
					ColumnNr: 1,
					RowNr:    14,
				},
				{
					Category: "regular",
					Type:     "emptyDouble",
					ColumnNr: 2,
					RowNr:    14,
				},
				{
					Category: "regular",
					Type:     "double",
					ColumnNr: 3,
					RowNr:    14,
				},
				{
					Category: "regular",
					Type:     "emptyDouble",
					ColumnNr: 4,
					RowNr:    14,
				},
				{
					Category: "regular",
					Type:     "double",
					ColumnNr: 5,
					RowNr:    14,
				},
				{
					Category: "regular",
					Type:     "emptyDouble",
					ColumnNr: 6,
					RowNr:    14,
				},
				{
					Category: "regular",
					Type:     "double",
					ColumnNr: 7,
					RowNr:    14,
				},
				{
					Category: "regular",
					Type:     "emptyDouble",
					ColumnNr: 8,
					RowNr:    14,
				},
				{
					Category: "regular",
					Type:     "double",
					ColumnNr: 9,
					RowNr:    14,
				},
				{
					Category: "regular",
					Type:     "emptyDouble",
					ColumnNr: 10,
					RowNr:    14,
				},
				{
					Category: "regular",
					Type:     "double",
					ColumnNr: 11,
					RowNr:    14,
				},
				{
					Category: "regular",
					Type:     "emptyDouble",
					ColumnNr: 12,
					RowNr:    14,
				},
				{
					Category: "regular",
					Type:     "double",
					ColumnNr: 13,
					RowNr:    14,
				},
				{
					Category: "regular",
					Type:     "emptyDouble",
					ColumnNr: 14,
					RowNr:    14,
				},
				{
					Category: "regular",
					Type:     "double",
					ColumnNr: 15,
					RowNr:    14,
				},
				{
					Category: "regular",
					Type:     "emptyDouble",
					ColumnNr: 16,
					RowNr:    14,
				},
				{
					Category: "regular",
					Type:     "double",
					ColumnNr: 17,
					RowNr:    14,
				},
				{
					Category: "regular",
					Type:     "emptyDouble",
					ColumnNr: 18,
					RowNr:    14,
				},
				{
					Category: "regular",
					Type:     "double",
					ColumnNr: 19,
					RowNr:    14,
				},
				{
					Category: "regular",
					Type:     "emptyDouble",
					ColumnNr: 20,
					RowNr:    14,
				},
				{
					Category: "regular",
					Type:     "double",
					ColumnNr: 21,
					RowNr:    14,
				},
				{
					Category: "regular",
					Type:     "emptyDouble",
					ColumnNr: 22,
					RowNr:    14,
				},
			},
		},

		TheatreId: &theatreId,
	}
}

func GetSampleCreateCinemaHallRequestNotRectangular() models.CreateCinemaHallRequest {
	validHall := GetSampleCreateCinemaHallRequest()
	validHall.Seats[0] = validHall.Seats[0][:len(validHall.Seats[0])-1]
	return validHall
}

func GetSampleCreateCinemaHallRequestInvalidSeatCategory() models.CreateCinemaHallRequest {
	validHall := GetSampleCreateCinemaHallRequest()
	validHall.Seats[0][0].Category = "invalid"
	return validHall
}

func GetSampleCreateCinemaHallRequestInvalidDoubleSeats() models.CreateCinemaHallRequest {
	validHall := GetSampleCreateCinemaHallRequest()
	validHall.Seats[14][2].Type = "regular"
	return validHall
}

func GetSampleCinemaHall() model.CinemaHalls {
	id := uuid.MustParse(cinemaHallId)
	theatreId := uuid.MustParse(theatreId)
	return model.CinemaHalls{
		ID:        &id,
		Name:      "HallName",
		Capacity:  23 * 15,
		TheatreID: &theatreId,
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
		},
		{
			ID:        utils.NewUUID(),
			TheatreID: &theatreId,
			Name:      "Hall 2",
			Capacity:  200,
		},
	}
}

func GetSampleSeat() model.Seats {
	id := uuid.MustParse(seatId)
	cinemaHallId := uuid.MustParse(cinemaHallId)
	seatCategoryId := uuid.MustParse(seatCategoryId1)
	return model.Seats{
		ID:             &id,
		ColumnNr:       0,
		RowNr:          0,
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
