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
		fmt.Println(r.URL)

		dir := "/home/pi/code/maxtaylordavi.es"

		cmd := exec.Command("/bin/sh", "-c", "git pull && go build && sudo systemctl restart maxtaylordavi.es.service")
		cmd.Dir = dir

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
