package utils

import (
	"database/sql/driver"
	"fmt"
	"reflect"
	"time"

	"github.com/ELITE-Kinoticketsystem/Backend-KTS/src/.gen/KinoTicketSystem/model"
	"github.com/ELITE-Kinoticketsystem/Backend-KTS/src/models"
	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"github.com/google/uuid"
)

// Compare two users while ignoring their ids and hashed passwords.
func UserEqual(user1 model.Users, user2 model.Users) bool {
	return cmp.Equal(user1, user2, cmpopts.IgnoreFields(model.Users{}, "ID", "Password"))
}

type UserMatcher struct {
	user     model.Users
	password string
}

func (m UserMatcher) Matches(x interface{}) bool {
	user, ok := x.(model.Users)
	if !ok {
		return false
	}
	if !ComparePasswordHash(m.password, user.Password) {
		return false
	}
	m.user.Password = user.Password

	// ignore user_id
	m.user.ID = user.ID

	return reflect.DeepEqual(m.user, user)
}

func (m UserMatcher) String() string {
	return fmt.Sprintf("matches user %v and password %s", m.user, m.password)
}

func EqUserMatcher(u model.Users, password string) UserMatcher {
	return UserMatcher{user: u, password: password}
}

// for matching a struct except for uuid fields
type IdMatcher struct {
	value interface{}
}

func (m IdMatcher) Matches(otherValue interface{}) bool {
	return cmp.Equal(m.value, otherValue, cmpopts.IgnoreTypes(&uuid.UUID{}))
}

func (m IdMatcher) String() string {
	return fmt.Sprintf("matches %v", m.value)
}

// Returns a matcher that matches the struct except for the uuid fields.
func EqExceptId(value interface{}) IdMatcher {
	return IdMatcher{value: value}
}

// for matching a uuid with its binary representation
type UUIDMatcher struct {
	id *uuid.UUID
}

func (m UUIDMatcher) Match(v driver.Value) bool {
	bytes, ok := v.(string)
	if !ok {
		return false
	}
	id, err := m.id.MarshalBinary()
	if err != nil {
		return false
	}
	return string(id) == bytes
}

// Returns a matcher that matches the uuid with its binary representation.
func EqUUID(id *uuid.UUID) UUIDMatcher {
	return UUIDMatcher{id: id}
}

func GetStringPointer(s string) *string {
	return &s
}

func GetSamplePriceCategories() *[]model.PriceCategories {
	priceCategories := []model.PriceCategories{}

	uuid1 := uuid.New()
	uuid2 := uuid.New()

	priceCategories = append(priceCategories, model.PriceCategories{
		ID:           &uuid1,
		CategoryName: "StudentDiscount",
		Price:        1000,
	})

	priceCategories = append(priceCategories, model.PriceCategories{
		ID:           &uuid2,
		CategoryName: "regular_price",
		Price:        500,
	})

	return &priceCategories
}

func GetSamplePriceCategory() *model.PriceCategories {
	priceCategory := model.PriceCategories{}

	uuid1 := uuid.New()

	priceCategory = model.PriceCategories{
		ID:           &uuid1,
		CategoryName: "StudentDiscount",
		Price:        1000,
	}

	return &priceCategory
}

func GetSampleTicket() *models.TicketDTO {
	ticket := models.TicketDTO{}

	id := uuid.New()
	seatId := uuid.New()
	eventId := uuid.New()
	orderId := uuid.New()
	cinemaHallId := uuid.New()
	userId := uuid.New()
	paymentMethodId := uuid.New()
	seatCategoryId := uuid.New()
	description := "Test Description"

	ticket = models.TicketDTO{
		ID:        &id,
		Validated: *new(bool),
		Price:     10,
		Seats: &model.Seats{
			ID:             &seatId,
			RowNr:          1,
			ColumnNr:       1,
			SeatCategoryID: &seatCategoryId,
			Type:           "Test Type",
		},
		Event: &model.Events{
			ID:           &eventId,
			Title:        "Test Event",
			Start:        time.Now(),
			End:          time.Now(),
			Description:  &description,
			EventType:    "Test EventType",
			CinemaHallID: &cinemaHallId,
		},
		Order: &model.Orders{
			ID:              &orderId,
			Totalprice:      1000,
			IsPaid:          false,
			PaymentMethodID: &paymentMethodId,
			UserID:          &userId,
		},
	}

	return &ticket
}

func GetSampleCreateTicket() *model.Tickets {
	ticket := model.Tickets{}

	id := uuid.New()
	eventSeatId := uuid.New()
	priceCategoryID := uuid.New()
	orderId := uuid.New()

	ticket = model.Tickets{
		ID:              &id,
		Validated:       *new(bool),
		Price:           10,
		OrderID:         &orderId,
		PriceCategoryID: &priceCategoryID,
		EventSeatID:     &eventSeatId,
	}

	return &ticket
}
