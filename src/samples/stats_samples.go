package samples

import (
	"time"

	"github.com/ELITE-Kinoticketsystem/Backend-KTS/src/models"
)

func GetSampleDayVisitsStats() *[]models.StatsVisits {
	day1 := time.Date(2024, 1, 2, 0, 0, 0, 0, time.UTC)
	day2 := time.Date(2024, 1, 3, 0, 0, 0, 0, time.UTC)

	return &[]models.StatsVisits{
		{
			Count:   2,
			Date:    day1,
			Revenue: 10,
		},
		{
			Count:   5,
			Date:    day2,
			Revenue: 25,
		},
	}
}

func GetSampleMonthVisitsStats() *[]models.StatsVisits {
	month1 := time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC)
	month2 := time.Date(2024, 3, 1, 0, 0, 0, 0, time.UTC)

	return &[]models.StatsVisits{
		{
			Count:   3,
			Date:    month1,
			Revenue: 15,
		},
		{
			Count:   1,
			Date:    month2,
			Revenue: 5,
		},
	}
}

func GetSampleYearVisitsStats() *[]models.StatsVisits {
	year1 := time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC)
	year2 := time.Date(2024, 1, 2, 0, 0, 0, 0, time.UTC)

	return &[]models.StatsVisits{
		{
			Count:   2,
			Date:    year1,
			Revenue: 10,
		},
		{
			Count:   5,
			Date:    year2,
			Revenue: 25,
		},
	}
}

func GetSampleDayVisitsStatsTwoArrays() *models.StatsVisitsTwoArrays {
	day1 := time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC)
	day2 := time.Date(2024, 1, 2, 0, 0, 0, 0, time.UTC)
	day3 := time.Date(2024, 1, 3, 0, 0, 0, 0, time.UTC)

	return &models.StatsVisitsTwoArrays{
		Count:   []int{0, 2, 5},
		Revenue: []int{0, 10, 25},
		Date:    []time.Time{day1, day2, day3},
	}
}

func GetSampleMonthVisitsStatsTwoArrays() *models.StatsVisitsTwoArrays {
	month1 := time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC)
	month2 := time.Date(2024, 2, 1, 0, 0, 0, 0, time.UTC)
	month3 := time.Date(2024, 3, 1, 0, 0, 0, 0, time.UTC)

	return &models.StatsVisitsTwoArrays{
		Count:   []int{3, 0, 1},
		Revenue: []int{15, 0, 5},
		Date:    []time.Time{month1, month2, month3},
	}
}

func GetSampleYearVisitsStatsTwoArrays() *models.StatsVisitsTwoArrays {
	year1 := time.Date(2022, 1, 1, 0, 0, 0, 0, time.UTC)
	year2 := time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC)
	year3 := time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC)

	return &models.StatsVisitsTwoArrays{
		Count:   []int{0, 2, 5},
		Revenue: []int{0, 10, 25},
		Date:    []time.Time{year1, year2, year3},
	}
}

func GetSampleEventWithTicketCount() *[]models.GetEventWithTicketCount {
	return &[]models.GetEventWithTicketCount{
		{
			EventName:   "Event1",
			TicketCount: 2,
		},
		{
			EventName:   "Event2",
			TicketCount: 5,
		},
	}
}

func GetSampleAllEvents() *[]models.GetEventsTitle {
	return &[]models.GetEventsTitle{
		{
			EventName: "Event1",
		},
		{
			EventName: "Event2",
		},
		{
			EventName: "Event3",
		},
	}
}

func GetSamplePreparedEvents() *[]models.GetEventWithTicketCount {
	return &[]models.GetEventWithTicketCount{
		{
			EventName:   "Event2",
			TicketCount: 5,
		},
		{
			EventName:   "Event1",
			TicketCount: 2,
		},
		{
			EventName:   "Event3",
			TicketCount: 0,
		},
	}
}
