package service

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"test/internal/store"
)

type Cv struct {
	store store.Store
}

func New(store store.Store) Cv {
	return Cv{store: store}
}

var ErrDb = errors.New("db error")

func (s Cv) Profile(ctx context.Context, lang string) (json.RawMessage, error) {
	raw, err := s.store.GetProfile(ctx, lang)
	if err != nil {
		log.Print(fmt.Sprintf("GetProfile: %v", err))
		return nil, ErrDb
	}

	return raw, nil
}
