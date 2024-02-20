package main

import (
	"log/slog"
	"net/http"
)

type Theme struct {
	Value bool
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		err := Body().Render(r.Context(), w)
		if err != nil {
			slog.Error("no funciona")
			return
		}
	})
	http.ListenAndServe(":80", nil)
}
