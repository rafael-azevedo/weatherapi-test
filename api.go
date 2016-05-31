package main

// WeatherData is a struct for the weather data the comes from the api calls
type WeatherData struct {
	Name string `json:"name"`
	Main struct {
		Kelvin float64 `json:"temp"`
	} `json:"main"`
}

// GoogleAPI struct hold google api key
type GoogleAPI struct {
	apikey string
}

// OpenWeatherMap struct to fufill weatherProvider interace
type OpenWeatherMap struct {
	apiKey string
}

// WeatherUnderground struct to fulfill weatherProvider interface
type WeatherUnderground struct {
	apiKey string
}

// ForecastIO struct to fulfill weatherProvider interface
type ForecastIO struct {
	apiKey string
}

// WeatherProvider interface can be fufilled with any struct that can fulfill the temperature func
type WeatherProvider interface {
	Temperature(city string) (float64, error) // in Kelvin, naturally
}

// MultiWeatherProvider slice can hold multiple weatherProvider structs
type MultiWeatherProvider []WeatherProvider
