package domain

type Resume struct {
	Email        Nullable[string]
	Phone        Nullable[string]
	Salary       Nullable[float64]
	SalarySuffix Nullable[string]
	FullName     Nullable[string]
	Residence    Nullable[string]
	Summary      Nullable[string]
}

type Project struct {
	Link   Nullable[string]
	Name   Nullable[string]
	Points string
}

type Experience struct {
	Company     Nullable[string]
	StartDate   Nullable[string]
	EndDate     Nullable[string]
	Achievement string
	Position    Nullable[string]
	Location    Nullable[string]
}

type Education struct {
	University  Nullable[string]
	StartDate   Nullable[string]
	EndDate     Nullable[string]
	Achievement string
}

type Proficiency struct {
	Level int
	Name  Nullable[string]
}

type Skill struct {
	Name      string
	SortOrder int
}
