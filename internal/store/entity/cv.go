package entity

import (
	"encoding/json"

	"github.com/jackc/pgx/v5/pgtype"
)

type Experience struct {
	Id           pgtype.Int8   `db:"id"`
	CompanyName  pgtype.Text   `db:"company_name"`
	PositionName pgtype.Text   `db:"position_name"`
	Skills       []pgtype.Text `db:"skills"`
	Location     pgtype.Text   `db:"location"`
	StartDt      pgtype.Date   `db:"start_dt"`
	EndDt        pgtype.Date   `db:"end_dt"`
}

type Profile struct {
	Id                pgtype.Int8     `db:"id"`
	Email             pgtype.Text     `db:"email"`
	Phone             pgtype.Text     `db:"phone"`
	Salary            pgtype.Numeric  `db:"salary"`
	SalaryExpectation pgtype.Text     `db:"salary_exp"`
	FullName          pgtype.Text     `db:"fullname"`
	Residence         pgtype.Text     `db:"residence"`
	Summary           pgtype.Text     `db:"summary"`
	Skills            json.RawMessage `db:"main_skills"`
}

type Education struct {
	Id             pgtype.Int8   `db:"id"`
	UniversityName pgtype.Text   `db:"university"`
	Achievements   []pgtype.Text `db:"achievements"`
	StartDt        pgtype.Date   `db:"start_dt"`
	EndDt          pgtype.Date   `db:"end_dt"`
}

type Language struct {
	Id    pgtype.Int8 `db:"id"`
	Txt   pgtype.Text `db:"value"`
	Level pgtype.Int8 `db:"level"`
}

type Project struct {
	Id          pgtype.Int8   `db:"id"`
	ProjectName pgtype.Text   `db:"pname"`
	Link        pgtype.Text   `db:"link"`
	LinkTxt     pgtype.Text   `db:"link_txt"`
	Points      []pgtype.Text `db:"points"`
}
