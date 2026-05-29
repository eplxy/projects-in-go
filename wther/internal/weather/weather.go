package weather

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"os"
	"strconv"
	"wther/internal/constants"

	"net/http"
	"net/url"
	"time"

	tea "charm.land/bubbletea/v2"
)

func GetForecast(queryOptions ForecastQueryOptions) tea.Msg {
	forecast, err := fetchForecast(queryOptions)
	if err != nil {
		return ForecastErrMsg{Err: err}
	}

	return ForecastFetchedMsg{Body: forecast}

}

func fetchForecast(queryOptions ForecastQueryOptions) (Forecast, error) {

	c := &http.Client{Timeout: 10 * time.Second}

	if _, ok := os.LookupEnv("WEATHER_API_KEY"); !ok {
		return Forecast{}, errors.New("weather api key is not set")
	}

	apiUrl := constants.API_BASE_URL

	params := url.Values{}
	params.Set("q", queryOptions.Location)
	params.Set("days", strconv.Itoa(queryOptions.Days))
	params.Set("key", os.Getenv("WEATHER_API_KEY"))
	params.Set("aqi", "no")
	params.Set("alerts", "no")
	apiUrl += "/forecast.json?" + params.Encode()

	res, err := c.Get(apiUrl)

	if err != nil {
		return Forecast{}, err
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(res.Body)
		return Forecast{}, fmt.Errorf("weather api returned %s: %s", res.Status, string(body))
	}

	var forecast Forecast
	if err := json.NewDecoder(res.Body).Decode(&forecast); err != nil {
		return Forecast{}, err
	}

	return forecast, nil
}

type ForecastErrMsg struct{ Err error }

type ForecastFetchedMsg struct {
	Body Forecast
}
