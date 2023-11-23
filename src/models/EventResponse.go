package models

import (
	"time"

	"github.com/google/uuid"
)

type EventResponse struct {
	Id              *uuid.UUID              `json:"id"`
	Title           string                  `json:"title"`
	Start           time.Time               `json:"start"`
	End             time.Time               `json:"end"`
	EventTypeId     string                  `json:"eventTypeId"`
	Movie           MovieResponse           `json:"movie"`
	PriceCategories []PriceCategoryResponse `json:"priceCategories"`
}

type PriceCategoryResponse struct {
	CategoryName string `json:"categoryName"`
	Price        int    `json:"price"`
}
