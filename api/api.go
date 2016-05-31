package api

var GoogleGeocodeURL string = "https://maps.googleapis.com/maps/api/geocode/json?"

// WeatherData is a struct for the weather data the comes from the api calls
type WeatherData struct {
	Name string `json:"name"`
	Main struct {
		Kelvin float64 `json:"temp"`
	} `json:"main"`
}

type GoogleGeocodeResponse struct {
	Results []struct {
		FormattedAddress string `json:"formatted_address"`
		Geometry         struct {
			Location struct {
				Lat float64 `json:"lat"`
				Lng float64 `json:"lng"`
			} `json:"location"`
		} `json:"geometry"`
	}
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
