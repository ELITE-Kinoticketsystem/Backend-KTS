package models

import "time"

type StatsVisits struct {
	Count   int       `alias:"COUNT(tickets.id)"`
	Date    time.Time `alias:"MIN(events.end)"`
	Revenue int       `alias:"SUM(orders.totalprice)"`
}

type StatsVisitsTwoArrays struct {
	Count   []int
	Date    []time.Time
	Revenue []int
}
