package handlers

import (
	"context"

	"github.com/JoelMauricio/OpenCVGenerator/models"
)

type contextKey string

var User contextKey = "user"

var emptyData = models.UserData{
	FirstName:              "Error",
	LastName:               "",
	Email:                  "",
	ContactPrimary:         "",
	ContactSecondary:       "",
	AcademicExperience:     []models.Academic{},
	ProfessionalExperience: []models.Professional{},
	Abilities:              []models.Ability{},
	Languages:              []models.Language{},
}

func GetUserContext(ctx context.Context) models.UserData {
	if udata, ok := ctx.Value(User).(models.UserData); ok {
		return udata
	}
	return emptyData
}

var CTX = context.WithValue(context.Background(), User, data)

var data = models.UserData{
	FirstName:        "Joel",
	LastName:         "Mauricio",
	Email:            "joeldavidmauriciohdez@gmail.com",
	ContactPrimary:   "+1 809-571-0824",
	ContactSecondary: "",
	AcademicExperience: []models.Academic{
		{StartDate: "2020", EndDate: "TODAY", Institution: "INTEC", Career: "Ing. Software", Description: "sa"},
	},
	ProfessionalExperience: []models.Professional{
		{StartDate: "2023", EndDate: "TODAY", BusinessName: "HECMAPP", PositionName: "Flutter Developer", JobDescription: "yes"},
	},
	Abilities: []models.Ability{
		{Name: "Flutter"},
		{Name: "Git"},
		{Name: "VsCode"},
	},
	Languages: []models.Language{
		{Name: "English"},
		{Name: "Spanish"},
	},
}
