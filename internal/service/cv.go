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

	if p.Summary.Valid {
		res.Summary = utils.Nullable[string]{
			Value: p.Summary.String,
			Valid: true,
		}
	}

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

func (s Cv) Experience(ctx context.Context, lang string) ([]responses.Experience, error) {
	var resp []responses.Experience

	pgmodels, err := s.store.GetExperience(ctx, lang)
	if err != nil {
		log.Print(fmt.Sprintf("GetExperience: %v", err))
		return nil, ErrDb
	}

	for _, m := range pgmodels {
		var r responses.Experience

		r.Id = m.Id.Int64
		r.CompanyName = m.CompanyName.String
		r.PositionName = m.PositionName.String

		for _, sk := range m.Skills {
			r.Skills = append(r.Skills, sk.String)
		}

		r.Location = m.Location.String

		r.StartDt = utils.Nullable[time.Time]{
			Valid: m.StartDt.Valid,
			Value: m.StartDt.Time,
		}

		r.EndDt = utils.Nullable[time.Time]{
			Valid: m.EndDt.Valid,
			Value: m.EndDt.Time,
		}

		resp = append(resp, r)
	}

	return resp, nil
}

func (s Cv) Education(ctx context.Context, lang string) ([]responses.Education, error) {
	var resp []responses.Education

	educations, err := s.store.GetEducation(ctx, lang)
	if err != nil {
		log.Print(fmt.Sprintf("GetEducation: %v", err))
		return nil, ErrDb
	}

	for _, e := range educations {
		var r responses.Education

		r.Id = e.Id.Int64
		r.University = e.UniversityName.String

		for _, a := range e.Achievements {
			r.Achievements = append(r.Achievements, a.String)
		}

		r.StartDt = e.StartDt.Time
		r.EndDt = e.EndDt.Time

		resp = append(resp, r)
	}

	return resp, nil
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
		points := make([]string, len(p.Points))

		for i, point := range p.Points {
			points[i] = point.String
		}

		projects = append(projects, responses.Project{
			Id:          p.Id.Int64,
			ProjectName: utils.Nullable[string]{Valid: p.ProjectName.Valid, Value: p.ProjectName.String},
			Link:        p.Link.String,
			LinkTxt:     utils.Nullable[string]{Valid: p.LinkTxt.Valid, Value: p.LinkTxt.String},
			Points:      points,
		})
	}

	return projects, nil
}
