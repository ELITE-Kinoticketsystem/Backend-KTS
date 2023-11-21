package utils

import (
	"fmt"
	"reflect"

	"github.com/ELITE-Kinoticketsystem/Backend-KTS/src/models/schemas"
)

type UserMatcher struct {
	user schemas.User
	password string
}

func (m UserMatcher) Matches(x interface{}) bool {
	user, ok := x.(schemas.User)
	if !ok {
		return false
	}
	if !ComparePasswordHash(m.password, user.Password) {
		return false
	}
	m.user.Password = user.Password

	// ignore user_id
	m.user.Id = user.Id

	return reflect.DeepEqual(m.user, user)
}

func (m UserMatcher) String() string {
	return fmt.Sprintf("matches user %v and password %s", m.user, m.password)
}

func EqUserMatcher(u schemas.User, password string) UserMatcher {
	return UserMatcher{user: u, password: password}
}