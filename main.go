package main

import (
	"encoding/json"
	"net/http"
	"strings"
	"time"
)

func main() {
	ga := GoogleAPI{apikey: ""}
	mw := MultiWeatherProvider{
		OpenWeatherMap{apiKey: ""},
		WeatherUnderground{apiKey: ""},
	}

	http.HandleFunc("/weather/", func(w http.ResponseWriter, r *http.Request) {
		begin := time.Now()
		city := strings.SplitN(r.URL.Path, "/", 3)[2]

		temp, err := mw.Temperature(city)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		json.NewEncoder(w).Encode(map[string]interface{}{
			"city": city,
			"temp": temp,
			"took": time.Since(begin).String(),
		})
	})

	http.ListenAndServe(":8080", nil)
}
