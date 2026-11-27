package sqlite

import (
	"context"
	"database/sql"
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
