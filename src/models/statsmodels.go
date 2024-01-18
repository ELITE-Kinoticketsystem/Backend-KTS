package models

import "time"

type StatsStruct struct {
	Count    int       `alias:"COUNT(orders.id)"`
	Duration int      
	Date     time.Time `alias:"(MIN(events.end))"`
}
