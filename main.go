package main

import (
	"net/http"
	"os/exec"
	"time"
)

func registerRoutes() http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		dir := "/home/pi/code/maxtaylordavi.es"

		cmd := exec.Command("git", "pull")
		cmd.Dir = dir

		err := cmd.Run()

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
