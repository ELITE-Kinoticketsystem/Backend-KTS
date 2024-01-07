package samples

import (
	"github.com/ELITE-Kinoticketsystem/Backend-KTS/src/gen/KinoTicketSystem/model"
	"github.com/ELITE-Kinoticketsystem/Backend-KTS/src/myid"
	"github.com/ELITE-Kinoticketsystem/Backend-KTS/src/utils"
)

func GetPriceCategories() *[]model.PriceCategories {
	return &[]model.PriceCategories{
		{
			ID:           myid.New(),
			CategoryName: utils.ADULT.String(),
			Price:        0,
		},
		{
			ID:           myid.New(),
			CategoryName: utils.CHILDREN.String(),
			Price:        10,
		},
		{
			ID:           myid.New(),
			CategoryName: utils.SENIOR.String(),
			Price:        5,
		},
	}
}

func GetSamplePriceCategories() *[]model.PriceCategories {
	priceCategories := []model.PriceCategories{}

	uuid1 := myid.New()
	uuid2 := myid.New()

	priceCategories = append(priceCategories, model.PriceCategories{
		ID:           uuid1,
		CategoryName: "StudentDiscount",
		Price:        1000,
	})

	priceCategories = append(priceCategories, model.PriceCategories{
		ID:           uuid2,
		CategoryName: "regular_price",
		Price:        500,
	})

	return &priceCategories
}

func GetSamplePriceCategory() *model.PriceCategories {
	priceCategory := model.PriceCategories{}

	uuid1 := myid.MustParse("6ba7b820-9dad-11d1-80b4-00c04fd430c4")

	priceCategory = model.PriceCategories{
		ID:           uuid1,
		CategoryName: "StudentDiscount",
		Price:        1000,
	}

	return &priceCategory
}
