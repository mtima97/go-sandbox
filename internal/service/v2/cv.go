package v2

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"test/internal/domain"
	"test/internal/models/responses"
)

type Db interface {
	GetEntityByName(ctx context.Context, name, language string) (json.RawMessage, error)
}

type CvSvc struct {
	db Db
}

func NewCvService(db Db) CvSvc {
	return CvSvc{db: db}
}

func (svc CvSvc) GetCV(ctx context.Context, language string) (responses.CV, error) {
	var (
		resp responses.CV
	)

	if language != domain.LangEn && language != domain.LangRu {
		return resp, errors.New("invalid language")
	}

	entities := []string{"profile", "projects", "languages", "experience", "education"}

	for _, entity := range entities {
		result, err := svc.db.GetEntityByName(ctx, entity, language)
		if err != nil {
			return resp, fmt.Errorf("db error: %w", err)
		}

		switch entity {
		case "profile":
			resp.Profile = result
		case "projects":
			resp.Projects = result
		case "languages":
			resp.Languages = result
		case "experience":
			resp.Experience = result
		case "education":
			resp.Education = result
		}
	}

	return resp, nil
}
