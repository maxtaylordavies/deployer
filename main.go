package main

import (
	"net/http"
	"os/exec"
	"time"
)

func registerRoutes() http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		cmd := exec.Command("git pull && go build && sudo systemctl restart maxtaylordavi.es.service")
		_, err := cmd.Output()

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	})

	return mux
}

func main() {
	server := http.Server{
		Addr:         ":8000",
		Handler:      registerRoutes(),
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 5 * time.Second,
		IdleTimeout:  120 * time.Second,
	}

	server.ListenAndServe()
}
