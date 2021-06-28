package main

import (
	"fmt"
	"net/http"
	"os/exec"
	"time"
)

func registerRoutes() http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		dir := "/home/pi/code/maxtaylordavi.es"

		cmds := []*exec.Cmd{
			exec.Command("git", "pull"),
			exec.Command("go", "build"),
			exec.Command("systemctl", "restart", "maxtaylordavi.es.service"),
		}

		for _, cmd := range cmds {
			cmd.Dir = dir
			err := cmd.Run()
			if err != nil {
				http.Error(w, fmt.Sprintf("%s stage failed with error '%s'", cmd.Path, err.Error()), http.StatusInternalServerError)
				return
			}
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
