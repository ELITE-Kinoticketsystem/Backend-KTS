package utils

const BLOCKED_TICKET_TIME = 15 * 60 // 15 minutes

type SeatType string

const (
	EMPTY        SeatType = "empty"
	EMPTY_DOUBLE SeatType = "emptyDouble"
	REGULAR      SeatType = "regular"
	DOUBLE       SeatType = "double"
)
