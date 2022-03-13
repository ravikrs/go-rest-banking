package app

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"time"
)

func Start() {
	r := mux.NewRouter()
	r.HandleFunc("/api/time", timeHandler)
	err := http.ListenAndServe("localhost:8080", r)
	log.Fatal(err)
}

func timeHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	tz := r.URL.Query().Get("tz")
	//tz := mux.Vars(r)["tz"]
	if tz == "" {
		tz = "UTC"
	}
	location, err := time.LoadLocation(tz)
	w.Header().Set("Content-Type", "application/json")
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		resp := make(map[string]string)
		resp["message"] = "invalid timezone"
		jsonResp, err := json.Marshal(resp)
		if err != nil {
			log.Fatalf("Error happened in JSON marshal. Err: %s", err)
		}
		w.Write(jsonResp)
		return
	}

	apiTime := ApiTime{CurrentTime: time.Now().In(location).Format("2006-01-02 15:04:05 -0700 MST")}
	json.NewEncoder(w).Encode(apiTime)
}

type ApiTime struct {
	CurrentTime string `json:"current_time"`
}
