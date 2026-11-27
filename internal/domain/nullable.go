package domain

type Nullable[T any] struct {
	Value T
	Valid bool
}
