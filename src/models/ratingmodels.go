package models

type NewRating struct {
	Rating float64 `alias:"SUM(reviews.rating)"`
}
