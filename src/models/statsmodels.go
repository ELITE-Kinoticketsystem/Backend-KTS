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

type GetEventWithTicketCount struct {
	EventName   string `alias:"events.title"`
	TicketCount int    `alias:"COUNT(tickets.id)"`
}

type GetEventsTitle struct {
	EventName   string `alias:"events.title"`
}
