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
