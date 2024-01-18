package models

import "time"

type StatsVisits struct {
	Count    int       `alias:"COUNT(orders.id)"`   
	Date     time.Time `alias:"MIN(events.end)"`
}

type StatsVisitsTwoArrays struct {
	Count    []int       
	Date     []time.Time 
}
