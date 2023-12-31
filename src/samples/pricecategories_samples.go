package samples

import (
	"github.com/ELITE-Kinoticketsystem/Backend-KTS/src/.gen/KinoTicketSystem/model"
	"github.com/ELITE-Kinoticketsystem/Backend-KTS/src/utils"
	"github.com/google/uuid"
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

func GetSamplePriceCategories() *[]model.PriceCategories {
	priceCategories := []model.PriceCategories{}

	uuid1 := uuid.New()
	uuid2 := uuid.New()

	priceCategories = append(priceCategories, model.PriceCategories{
		ID:           &uuid1,
		CategoryName: "StudentDiscount",
		Price:        1000,
	})

	priceCategories = append(priceCategories, model.PriceCategories{
		ID:           &uuid2,
		CategoryName: "regular_price",
		Price:        500,
	})

	return &priceCategories
}

func GetSamplePriceCategory() *model.PriceCategories {
	priceCategory := model.PriceCategories{}

	uuid1 := uuid.MustParse("6ba7b820-9dad-11d1-80b4-00c04fd430c4")

	priceCategory = model.PriceCategories{
		ID:           &uuid1,
		CategoryName: "StudentDiscount",
		Price:        1000,
	}

	return &priceCategory
}

