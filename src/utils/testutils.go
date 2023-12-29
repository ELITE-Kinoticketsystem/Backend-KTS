package utils

import (
	"database/sql/driver"
	"fmt"
	"reflect"

	"github.com/ELITE-Kinoticketsystem/Backend-KTS/src/.gen/KinoTicketSystem/model"
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
