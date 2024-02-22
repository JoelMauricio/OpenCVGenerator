package main

import (
	"fmt"
	"log/slog"
	"net/http"

	"github.com/JoelMauricio/OpenCVGenerator/web/templates"
)

type Theme struct {
	Value bool
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		err := templates.Body().Render(r.Context(), w)
		if err != nil {
			slog.Error("no funciona")
			return
		}
	})
	http.HandleFunc("/Editor", func(w http.ResponseWriter, r *http.Request) {
		err := templates.Layout().Render(r.Context(), w)
		if err != nil {
			slog.Error("no funciona")
			return
		}
	})
	fmt.Println("Running...")
	http.ListenAndServe(":80", nil)

}
