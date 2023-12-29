package samples

import (
	"github.com/ELITE-Kinoticketsystem/Backend-KTS/src/.gen/KinoTicketSystem/model"
	"github.com/ELITE-Kinoticketsystem/Backend-KTS/src/utils"
)

func GetPriceCategories() *[]model.PriceCategories {
	return &[]model.PriceCategories{
		{
			ID:           utils.NewUUID(),
			CategoryName: utils.ADULT.String(),
			Price:        0,
		},
		{
			ID:           utils.NewUUID(),
			CategoryName: utils.CHILDREN.String(),
			Price:        10,
		},
		{
			ID:           utils.NewUUID(),
			CategoryName: utils.SENIOR.String(),
			Price:        5,
		},
	}
}
