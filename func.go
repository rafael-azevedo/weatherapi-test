package main

import (
	"encoding/json"
	"log"
	"net/http"
)

// Hello prints hello
func Hello(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello!"))
}

// Temperature returns temperature of city in kelvin for OpenWeatherMap Struct
func (w OpenWeatherMap) Temperature(city string) (float64, error) {
	resp, err := http.Get("http://api.openweathermap.org/data/2.5/weather?APPID=" + w.apiKey + "&q=" + city)
	if err != nil {
		return 0, err
	}

	defer resp.Body.Close()

	var d struct {
		Main struct {
			Kelvin float64 `json:"temp"`
		} `json:"main"`
	}

	if err := json.NewDecoder(resp.Body).Decode(&d); err != nil {
		return 0, err
	}

	log.Printf("openWeatherMap: %s: %.2f", city, d.Main.Kelvin)
	return d.Main.Kelvin, nil
}

// Temperature returns temperature of city in kelvin for WeatherUnderground Struct
func (w WeatherUnderground) Temperature(city string) (float64, error) {
	resp, err := http.Get("http://api.wunderground.com/api/" + w.apiKey + "/conditions/q/" + city + ".json")
	if err != nil {
		return 0, err
	}

	defer resp.Body.Close()

	var d struct {
		Observation struct {
			Celsius float64 `json:"temp_c"`
		} `json:"current_observation"`
	}

	if err := json.NewDecoder(resp.Body).Decode(&d); err != nil {
		return 0, err
	}

	kelvin := d.Observation.Celsius + 273.15
	log.Printf("weatherUnderground: %s: %.2f", city, kelvin)
	return kelvin, nil
}

// Temperature Returns temp for MultiWeatherProvider individual provider using goroutines
func (w MultiWeatherProvider) Temperature(city string) (float64, error) {
	temps := make(chan float64, len(w))
	errs := make(chan error, len(w))

	for _, provider := range w {
		go func(p WeatherProvider) {
			k, err := p.Temperature(city)
			if err != nil {
				errs <- err
				return
			}
			temps <- k
		}(provider)
	}
	sum := 0.0

	for i := 0; i < len(w); i++ {
		select {
		case temp := <-temps:
			sum += temp
		case err := <-errs:
			return 0, err
		}
	}
	return sum / float64(len(w)), nil
}

func (g GoogleApi) (float64, float64) {
    var log float64
    var lat float64
    return log, lat
}
// Temperature returns temperature of city in kelvin for multiple providers depreciated
//func Temperature(city string, providers ...WeatherProvider) (float64, error) {
//	sum := 0.0
//
//	for _, provider := range providers {
//		k, err := provider.Temperature(city)
//		if err != nil {
//			return 0, err
//		}
//
//		sum += k
//	}
//
//	return sum / float64(len(providers)), nil
//}
