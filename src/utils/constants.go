package utils

import "time"

const BLOCKED_TICKET_TIME time.Duration = time.Minute * 15 // 15 minutes

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
	SENIOR   PriceCategories = "pensioner"
	STUDENT  PriceCategories = "student"
)

const URL = "cinemika.germanywestcentral.cloudapp.azure.com"