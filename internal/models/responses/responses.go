package responses

import (
	"encoding/json"
	"test/internal/utils"
	"time"
)

type Experience struct {
	Id           int64                     `json:"id"`
	CompanyName  string                    `json:"company_name"`
	PositionName string                    `json:"position_name"`
	Skills       []string                  `json:"skills"`
	StartDt      utils.Nullable[time.Time] `json:"start_dt"`
	EndDt        utils.Nullable[time.Time] `json:"end_dt"`
}

type Profile struct {
	Id                int64           `json:"id"`
	Email             string          `json:"email"`
	Phone             string          `json:"phone"`
	Salary            float64         `json:"salary"`
	SalaryExpectation string          `json:"salary_exp"`
	FullName          string          `json:"fullname"`
	Residence         string          `json:"residence"`
	Skills            json.RawMessage `json:"skills"`
}

type Education struct {
	Id           int64     `json:"id"`
	University   string    `json:"university"`
	Achievements []string  `json:"achievements"`
	StartDt      time.Time `json:"start_dt"`
	EndDt        time.Time `json:"end_dt"`
}

type Language struct {
	Id    int64                  `json:"id"`
	Lang  utils.Nullable[string] `json:"language"`
	Level int64                  `json:"level"`
}

type Project struct {
	Id          int64                  `json:"id"`
	ProjectName utils.Nullable[string] `json:"project_name"`
	Link        string                 `json:"link"`
	LinkTxt     utils.Nullable[string] `json:"link_txt"`
}
