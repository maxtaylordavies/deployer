package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os/exec"
	"time"
)

type data struct {
	Repo string `json:"repo"`
}

func registerRoutes() http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("/deploy", func(w http.ResponseWriter, r *http.Request) {
		decoder := json.NewDecoder(r.Body)
		var data data

		err := decoder.Decode(&data)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		cmd := exec.Command("/bin/sh", "-c", fmt.Sprintf("git pull && sudo systemctl restart %s.service", data.Repo))
		cmd.Dir = "/home/pi/code/" + data.Repo

		go cmd.Output()
		w.WriteHeader(http.StatusOK)
		// if err != nil {
		// 	http.Error(w, "output: "+string(out)+", error: "+err.Error(), http.StatusInternalServerError)
		// 	return
		// }
	})

	return mux
}

func main() {
	server := http.Server{
		Addr:         ":9000",
		Handler:      registerRoutes(),
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 5 * time.Second,
		IdleTimeout:  600 * time.Second,
	}

	server.ListenAndServe()
}
