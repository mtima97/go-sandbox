package utils

import "encoding/json"

type Nullable[T any] struct {
	Value T
	Valid bool
}

func (r Nullable[T]) MarshalJSON() ([]byte, error) {
	if !r.Valid {
		return []byte("null"), nil
	}

	return json.Marshal(r.Value)
}
