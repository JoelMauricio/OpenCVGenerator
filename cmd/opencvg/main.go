package main

import (
	"fmt"
	"log/slog"
	"net/http"

	"github.com/JoelMauricio/OpenCVGenerator/models"
	"github.com/JoelMauricio/OpenCVGenerator/web/templates"
)

type Theme struct {
	Value bool
}

func main() {
	data := models.UserData{
		"Joel",
		"Mauricio",
		"joeldavidmauriciohdez@gmail.com",
		"+1 809-571-0824",
		"",
		[]models.Academic{
			models.Academic{"2020", "TODAY", "INTEC", "Ing. Software", "sa"},
		},
		[]models.Professional{
			models.Professional{"2023", "TODAY", "HECMAPP", "Flutter Developer", "yes"},
		},
		[]models.Ability{
			models.Ability{"as"},
			models.Ability{"sas"},
			models.Ability{"ass"},
		},
		[]models.Language{
			models.Language{"English"},
			models.Language{"Spanish"},
		},
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		err := templates.Body().Render(r.Context(), w)
		if err != nil {
			slog.Error("no funciona")
			return
		}
	})
	http.HandleFunc("/Editor", func(w http.ResponseWriter, r *http.Request) {
		err := templates.Layout(data).Render(r.Context(), w)
		if err != nil {
			slog.Error("no funciona")
			return
		}
	})

	fmt.Println("Running...")
	http.ListenAndServe(":80", nil)

}
