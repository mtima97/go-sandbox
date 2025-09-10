package store

import (
	"context"
	"encoding/json"
)

type Store interface {
	GetProfile(ctx context.Context, lang string) (json.RawMessage, error)
}
