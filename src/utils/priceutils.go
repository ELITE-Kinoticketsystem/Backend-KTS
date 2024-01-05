package utils

func CalculatePrice(seatPrice int32, priceCategoryDiscount int32) int32 {
	return seatPrice * (1 - priceCategoryDiscount / 100)
}
