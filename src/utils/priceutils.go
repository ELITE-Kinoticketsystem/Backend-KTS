package utils

func CalculatePrice(seatPrice int32, priceCategoryDiscount int32, seatType string) int32 {
	test := float64(priceCategoryDiscount)
	discount := 1. - test/100.
	singleSeatPrice := int32(float64(seatPrice) * discount)
	if seatType == string(DOUBLE) {
		return singleSeatPrice * 2
	}
	return singleSeatPrice
}
