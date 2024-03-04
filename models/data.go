package models

type UserData struct {
	FirstName              string
	LastName               string
	Email                  string
	ContactPrimary         string
	ContactSecondary       string
	AcademicExperience     []Academic
	ProfessionalExperience []Professional
	Abilities              []Ability
	Languages              []Language
}

type Academic struct {
	StartDate   string
	EndDate     string
	Institution string
	Career      string
	Description string
}

type Professional struct {
	StartDate      string
	EndDate        string
	BusinessName   string
	PositionName   string
	JobDescription string
}

type Ability struct {
	Name string
}

type Language struct {
	Name string
}
