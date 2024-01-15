package models

type NewRating struct {
	TotalRatings int64   `alias:"COUNT(reviews.rating)"`
	Rating       float64 `alias:"SUM(reviews.rating)"`
}
