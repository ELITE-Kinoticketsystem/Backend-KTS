package utils

func CalculatePrice(seatPrice int32, priceCategoryDiscount int32, seatType string) int32 {
	singleSeatPrice := seatPrice * (1 - priceCategoryDiscount/100)
	if seatType == string(DOUBLE) {
		return singleSeatPrice * 2
	}
	return singleSeatPrice
}
