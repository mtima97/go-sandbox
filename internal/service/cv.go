package service

import (
	"context"
	"errors"
	"fmt"
	"log"
	"strconv"
	"test/internal/models/responses"
	"test/internal/store"
	"test/internal/utils"
	"time"
)

type Cv struct {
	store store.Store
}

func New(store store.Store) Cv {
	return Cv{store: store}
}

var ErrDb = errors.New("db error")

func (s Cv) Profile(ctx context.Context, lang string) (responses.Profile, error) {
	var res responses.Profile

	p, err := s.store.GetProfile(ctx, lang)
	if err != nil {
		log.Print(fmt.Sprintf("GetProfile: %v", err))
		return res, ErrDb
	}

	res.Id = p.Id.Int64
	res.Email = p.Email.String
	res.Phone = p.Phone.String
	res.SalaryExpectation = p.SalaryExpectation.String
	res.FullName = p.FullName.String
	res.Residence = p.Residence.String
	res.Skills = p.Skills

	v, err := p.Salary.Value()
	if err != nil {
		return res, err
	}

	res.Salary, err = strconv.ParseFloat(v.(string), 64)
	if err != nil {
		return res, err
	}

	return res, nil
}

func (s Cv) Experience(ctx context.Context, lang string) (responses.Experience, error) {
	var res responses.Experience

	pgmodel, err := s.store.GetExperience(ctx, lang)
	if err != nil {
		log.Print(fmt.Sprintf("GetExperience: %v", err))
		return res, ErrDb
	}

	res.Id = pgmodel.Id.Int64
	res.CompanyName = pgmodel.CompanyName.String
	res.PositionName = pgmodel.PositionName.String

	for _, sk := range pgmodel.Skills {
		res.Skills = append(res.Skills, sk.String)
	}

	res.StartDt = utils.Nullable[time.Time]{
		Valid: pgmodel.StartDt.Valid,
		Value: pgmodel.StartDt.Time,
	}

	res.EndDt = utils.Nullable[time.Time]{
		Valid: pgmodel.EndDt.Valid,
		Value: pgmodel.EndDt.Time,
	}

	return res, nil
}

func (s Cv) Education(ctx context.Context, lang string) (responses.Education, error) {
	var res responses.Education

	pgmodel, err := s.store.GetEducation(ctx, lang)
	if err != nil {
		log.Print(fmt.Sprintf("GetEducation: %v", err))
		return res, ErrDb
	}

	res.Id = pgmodel.Id.Int64
	res.University = pgmodel.UniversityName.String

	for _, a := range pgmodel.Achievements {
		res.Achievements = append(res.Achievements, a.String)
	}

	res.StartDt = pgmodel.StartDt.Time
	res.EndDt = pgmodel.EndDt.Time

	return res, nil
}

func (s Cv) Languages(ctx context.Context, lang string) ([]responses.Language, error) {
	var languages []responses.Language

	pgmodels, err := s.store.GetLanguages(ctx, lang)
	if err != nil {
		log.Print(fmt.Sprintf("GetLanguages: %v", err))
		return nil, ErrDb
	}

	for _, l := range pgmodels {
		languages = append(languages, responses.Language{
			Id:    l.Id.Int64,
			Lang:  utils.Nullable[string]{Valid: l.Txt.Valid, Value: l.Txt.String},
			Level: l.Level.Int64,
		})
	}

	return languages, nil
}

func (s Cv) Projects(ctx context.Context, lang string) ([]responses.Project, error) {
	var projects []responses.Project

	pgmodels, err := s.store.GetProjects(ctx, lang)
	if err != nil {
		log.Print(fmt.Sprintf("GetProjects: %v", err))
		return nil, ErrDb
	}

	for _, p := range pgmodels {
		projects = append(projects, responses.Project{
			Id:          p.Id.Int64,
			ProjectName: utils.Nullable[string]{Valid: p.ProjectName.Valid, Value: p.ProjectName.String},
			Link:        p.Link.String,
			LinkTxt:     utils.Nullable[string]{Valid: p.LinkTxt.Valid, Value: p.LinkTxt.String},
		})
	}

	return projects, nil
}
