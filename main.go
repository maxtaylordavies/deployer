package main

import (
	"fmt"
	"net/http"
	"os/exec"
	"time"
)

func registerRoutes() http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("/deploy", func(w http.ResponseWriter, r *http.Request) {
		repo := r.URL.Query().Get("repo")
		if repo == "" {
			http.Error(w, "repo query param is required", http.StatusBadRequest)
			return
		}

		cmd := exec.Command("/bin/sh", "-c", fmt.Sprintf("git pull && go build && sudo systemctl restart %s.service", repo))
		cmd.Dir = "/home/pi/code/" + repo

		out, err := cmd.Output()
		if err != nil {
			http.Error(w, "output: "+string(out)+", error: "+err.Error(), http.StatusInternalServerError)
			return
		}
	})

	return mux
}

func main() {
	server := http.Server{
		Addr:         ":9000",
		Handler:      registerRoutes(),
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 5 * time.Second,
		IdleTimeout:  120 * time.Second,
	}

	server.ListenAndServe()
}
