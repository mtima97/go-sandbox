package sqlite

import (
	"context"
	"database/sql"
	"fmt"
	"test/internal/domain"
)

func (d Db) GetResume(ctx context.Context, langID int) (resume domain.Resume, err error) {
	q := `
select
	r.email,
	r.phone,
	r.salary,
	rt.salary_sfx,
	rt.fullname,
	rt.residence,
	rt.summary
from main.resume as r
join main.resume_t9n as rt on r.id = rt.resume_id
where rt.lang_id = ?
`

	var (
		email        sql.NullString
		phone        sql.NullString
		salary       sql.NullFloat64
		salarySuffix sql.NullString
		fullName     sql.NullString
		residence    sql.NullString
		summary      sql.NullString
	)

	err = d.db.QueryRowContext(ctx, q, langID).Scan(
		&email,
		&phone,
		&salary,
		&salarySuffix,
		&fullName,
		&residence,
		&summary,
	)

	if err != nil {
		return
	}

	resume.Email = domain.Nullable[string]{
		Value: email.String,
		Valid: email.Valid,
	}
	resume.Phone = domain.Nullable[string]{
		Value: phone.String,
		Valid: phone.Valid,
	}
	resume.Salary = domain.Nullable[float64]{
		Value: salary.Float64,
		Valid: salary.Valid,
	}
	resume.SalarySuffix = domain.Nullable[string]{
		Value: salarySuffix.String,
		Valid: salarySuffix.Valid,
	}
	resume.FullName = domain.Nullable[string]{
		Value: fullName.String,
		Valid: fullName.Valid,
	}
	resume.Residence = domain.Nullable[string]{
		Value: residence.String,
		Valid: residence.Valid,
	}
	resume.Summary = domain.Nullable[string]{
		Value: summary.String,
		Valid: summary.Valid,
	}

	return
}

func (d Db) GetProjects(ctx context.Context, langID int) ([]domain.Project, error) {
	const q = `
select p.link, pt.name, pt.points
from main.project as p
join main.project_t9n as pt on pt.project_id = p.id
where pt.lang_id = ?
order by p.id
`

	rows, err := d.db.QueryContext(ctx, q, langID)
	if err != nil {
		return nil, fmt.Errorf("query projects: %w", err)
	}
	defer rows.Close()

	projects := make([]domain.Project, 0)
	for rows.Next() {
		var (
			project domain.Project
			link    sql.NullString
			name    sql.NullString
		)

		if err := rows.Scan(&link, &name, &project.Points); err != nil {
			return nil, fmt.Errorf("scan project: %w", err)
		}

		project.Link = domain.Nullable[string]{Value: link.String, Valid: link.Valid}
		project.Name = domain.Nullable[string]{Value: name.String, Valid: name.Valid}
		projects = append(projects, project)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("iterate projects: %w", err)
	}

	return projects, nil
}

//goland:noinspection GoUnhandledErrorResult
func (d Db) GetExperiences(ctx context.Context, langID int) ([]domain.Experience, error) {
	const q = `
select e.company, e.start_dt, e.end_dt, et.achiev, et.position, et.location
from main.experience as e
join main.experience_t9n as et on et.experience_id = e.id
where et.lang_id = ?
order by e.id
`

	rows, err := d.db.QueryContext(ctx, q, langID)
	if err != nil {
		return nil, fmt.Errorf("query experiences: %w", err)
	}
	defer rows.Close()

	experiences := make([]domain.Experience, 0)
	for rows.Next() {
		var (
			experience domain.Experience
			company    sql.NullString
			startDate  sql.NullString
			endDate    sql.NullString
			position   sql.NullString
			location   sql.NullString
		)

		if err := rows.Scan(
			&company,
			&startDate,
			&endDate,
			&experience.Achievement,
			&position,
			&location,
		); err != nil {
			return nil, fmt.Errorf("scan experience: %w", err)
		}

		experience.Company = domain.Nullable[string]{Value: company.String, Valid: company.Valid}
		experience.StartDate = domain.Nullable[string]{Value: startDate.String, Valid: startDate.Valid}
		experience.EndDate = domain.Nullable[string]{Value: endDate.String, Valid: endDate.Valid}
		experience.Position = domain.Nullable[string]{Value: position.String, Valid: position.Valid}
		experience.Location = domain.Nullable[string]{Value: location.String, Valid: location.Valid}
		experiences = append(experiences, experience)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("iterate experiences: %w", err)
	}

	return experiences, nil
}

//goland:noinspection GoUnhandledErrorResult
func (d Db) GetEducation(ctx context.Context, langID int) ([]domain.Education, error) {
	const q = `
select e.university, e.start_dt, e.end_dt, et.achiev
from main.education as e
join main.education_t9n as et on et.education_id = e.id
where et.lang_id = ?
order by e.id
`

	rows, err := d.db.QueryContext(ctx, q, langID)
	if err != nil {
		return nil, fmt.Errorf("query education: %w", err)
	}
	defer rows.Close()

	education := make([]domain.Education, 0)
	for rows.Next() {
		var (
			item       domain.Education
			university sql.NullString
			startDate  sql.NullString
			endDate    sql.NullString
		)

		if err := rows.Scan(&university, &startDate, &endDate, &item.Achievement); err != nil {
			return nil, fmt.Errorf("scan education: %w", err)
		}

		item.University = domain.Nullable[string]{Value: university.String, Valid: university.Valid}
		item.StartDate = domain.Nullable[string]{Value: startDate.String, Valid: startDate.Valid}
		item.EndDate = domain.Nullable[string]{Value: endDate.String, Valid: endDate.Valid}
		education = append(education, item)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("iterate education: %w", err)
	}

	return education, nil
}

//goland:noinspection GoUnhandledErrorResult
func (d Db) GetProficiencies(ctx context.Context, langID int) ([]domain.Proficiency, error) {
	const q = `
select p.level, pt.name
from main.proficiency as p
join main.proficiency_t9n as pt on pt.proficiency_id = p.id
where pt.lang_id = ?
order by p.id
`

	rows, err := d.db.QueryContext(ctx, q, langID)
	if err != nil {
		return nil, fmt.Errorf("query proficiencies: %w", err)
	}
	defer rows.Close()

	proficiencies := make([]domain.Proficiency, 0)
	for rows.Next() {
		var (
			proficiency domain.Proficiency
			name        sql.NullString
		)

		if err := rows.Scan(&proficiency.Level, &name); err != nil {
			return nil, fmt.Errorf("scan proficiency: %w", err)
		}

		proficiency.Name = domain.Nullable[string]{Value: name.String, Valid: name.Valid}
		proficiencies = append(proficiencies, proficiency)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("iterate proficiencies: %w", err)
	}

	return proficiencies, nil
}

//goland:noinspection GoUnhandledErrorResult
func (d Db) GetSkills(ctx context.Context) ([]domain.Skill, error) {
	const q = `
select s.name, s.sort_order
from main.skills as s
order by s.sort_order
`

	rows, err := d.db.QueryContext(ctx, q)
	if err != nil {
		return nil, fmt.Errorf("query skills: %w", err)
	}
	defer rows.Close()

	skills := make([]domain.Skill, 0)
	for rows.Next() {
		var skill domain.Skill
		if err := rows.Scan(&skill.Name, &skill.SortOrder); err != nil {
			return nil, fmt.Errorf("scan skill: %w", err)
		}
		skills = append(skills, skill)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("iterate skills: %w", err)
	}

	return skills, nil
}
