package store

import (
	"context"
	"test/internal/store/entity"
)

type Store interface {
	GetProfile(ctx context.Context, lang string) (entity.Profile, error)
	GetExperience(ctx context.Context, lang string) (entity.Experience, error)
	GetEducation(ctx context.Context, lang string) (entity.Education, error)
	GetLanguages(ctx context.Context, lang string) ([]entity.Language, error)
	GetProjects(ctx context.Context, lang string) ([]entity.Project, error)
}
