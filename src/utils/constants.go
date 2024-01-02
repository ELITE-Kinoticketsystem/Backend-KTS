package utils

import "time"

const BLOCKED_TICKET_TIME time.Duration = 9e+11 // 15 minutes

type SeatType string

const (
	EMPTY        SeatType = "empty"
	EMPTY_DOUBLE SeatType = "emptyDouble"
	REGULAR      SeatType = "regular"
	DOUBLE       SeatType = "double"
)

type PriceCategories string

func (pc PriceCategories) String() string {
	return string(pc)
}

const (
	CHILDREN PriceCategories = "children"
	ADULT    PriceCategories = "adult"
	SENIOR   PriceCategories = "senior"
)
