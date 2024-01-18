package samples

import (
	"time"

	"github.com/ELITE-Kinoticketsystem/Backend-KTS/src/models"
)

func GetSampleDayVisitsStats() []models.StatsVisits {

	day1 := time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC)
	day2 := time.Date(2022, 4, 21, 0, 0, 0, 0, time.UTC)

	return []models.StatsVisits{
		{
			Count: 2,
			Date:  day1,
		},
		{
			Count: 5,
			Date:  day2,
		},
	}

}

func GetSampleMonthVisitsStats() []models.StatsVisits {
	month1 := time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC)
	month2 := time.Date(2022, 5, 1, 0, 0, 0, 0, time.UTC)

	return []models.StatsVisits{
		{
			Count: 3,
			Date:  month1,
		},
		{
			Count: 1,
			Date:  month2,
		},
	}

}

func GetSampleYearVisitsStats() []models.StatsVisits {
	year1 := time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC)
	year2 := time.Date(2022, 1, 1, 0, 0, 0, 0, time.UTC)

	return []models.StatsVisits{
		{
			Count: 2,
			Date:  year1,
		},
		{
			Count: 5,
			Date:  year2,
		},
	}

}
