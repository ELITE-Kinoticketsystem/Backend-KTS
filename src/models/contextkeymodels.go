package models

type ContextKey string

const (
	ContextKeyUserID = ContextKey("userId")
)

func (c ContextKey) String() string {
	return string(c)
}
