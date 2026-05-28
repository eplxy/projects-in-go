package weather

import (
	"fmt"
	"math"
	"strings"
)

type Forecast struct {
	Location       Location `json:"location"`
	Current        Current  `json:"current"`
	ForecastObject struct {
		ForecastList []ForecastDay `json:"forecastday"`
	} `json:"forecast"`
}

type Location struct {
	Name    string `json:"name"`
	Region  string `json:"region"`
	Country string `json:"country"`
}

type Condition struct {
	Text string `json:"text"`
	Icon string `json:"icon"`
	Code int    `json:"code"`
}

type WeatherReading struct {
	TempC        float64   `json:"temp_c"`
	TempF        float64   `json:"temp_f"`
	IsDay        int       `json:"is_day"`
	Condition    Condition `json:"condition"`
	WindMph      float64   `json:"wind_mph"`
	WindKph      float64   `json:"wind_kph"`
	WindDegree   int       `json:"wind_degree"`
	WindDir      string    `json:"wind_dir"`
	PressureMb   float64   `json:"pressure_mb"`
	PressureIn   float64   `json:"pressure_in"`
	PrecipMm     float64   `json:"precip_mm"`
	PrecipIn     float64   `json:"precip_in"`
	Humidity     int       `json:"humidity"`
	Cloud        int       `json:"cloud"`
	FeelslikeC   float64   `json:"feelslike_c"`
	FeelslikeF   float64   `json:"feelslike_f"`
	WindchillC   float64   `json:"windchill_c"`
	WindchillF   float64   `json:"windchill_f"`
	HeatindexC   float64   `json:"heatindex_c"`
	HeatindexF   float64   `json:"heatindex_f"`
	WillItRain   int       `json:"will_it_rain"`
	ChanceOfRain int       `json:"chance_of_rain"`
	WillItSnow   int       `json:"will_it_snow"`
	ChanceOfSnow int       `json:"chance_of_snow"`
	GustMph      float64   `json:"gust_mph"`
	GustKph      float64   `json:"gust_kph"`
	Uv           float64   `json:"uv"`
}

type Current struct {
	LastUpdatedEpoch int    `json:"last_updated_epoch"`
	LastUpdated      string `json:"last_updated"`
	WeatherReading
}

type ForecastDay struct {
	Date      string `json:"date"`
	DateEpoch int    `json:"date_epoch"`
	Day       struct {
		MaxtempC          float64   `json:"maxtemp_c"`
		MaxtempF          float64   `json:"maxtemp_f"`
		MintempC          float64   `json:"mintemp_c"`
		MintempF          float64   `json:"mintemp_f"`
		AvgtempC          float64   `json:"avgtemp_c"`
		AvgtempF          float64   `json:"avgtemp_f"`
		MaxwindMph        float64   `json:"maxwind_mph"`
		MaxwindKph        float64   `json:"maxwind_kph"`
		TotalprecipMm     float64   `json:"totalprecip_mm"`
		TotalprecipIn     float64   `json:"totalprecip_in"`
		TotalsnowCm       float64   `json:"totalsnow_cm"`
		Avghumidity       int       `json:"avghumidity"`
		DailyWillItRain   int       `json:"daily_will_it_rain"`
		DailyChanceOfRain int       `json:"daily_chance_of_rain"`
		DailyWillItSnow   int       `json:"daily_will_it_snow"`
		DailyChanceOfSnow int       `json:"daily_chance_of_snow"`
		Condition         Condition `json:"condition"`
		Uv                float64   `json:"uv"`
	} `json:"day"`
	Hour []Hour `json:"hour"`
}

type Hour struct {
	TimeEpoch int    `json:"time_epoch"`
	Time      string `json:"time"`
	WeatherReading
	SnowCm float64 `json:"snow_cm"`
}

type ForecastQueryOptions struct {
}

func (f Forecast) String() string {

	if f.Current.Condition.Text == "" {
		return "No weather data is currently available."
	}

	s := fmt.Sprintf("Currently in %s, it's %d and %s", f.Location.Name, int(math.Round(f.Current.TempC)), strings.ToLower(f.Current.Condition.Text))

	return s

}
