package tui

import (
	"wther/internal/weather"

	tea "charm.land/bubbletea/v2"
)

type FetchStatus int

const (
	None FetchStatus = iota
	Fetching
	Error
	Success
)

type model struct {
	err         error
	fetchStatus FetchStatus // none, fetching, error, success

	forecast weather.Forecast
	// mode ViewMode

	// query WeatherQuery
	// config AppConfig
}

func NewModel() model {
	return model{}
}

func (m model) Init() tea.Cmd {
	return weather.GetForecast
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {

	// Is it a key press?
	case tea.KeyPressMsg:

		// Cool, what was the actual key pressed?
		switch msg.String() {

		// These keys should exit the program.
		case "ctrl+c", "q":
			return m, tea.Quit
		}
	case weather.ForecastFetchedMsg:
		m.forecast = msg.Body
	case weather.ForecastErrMsg:
		m.err = msg.Err
	}

	return m, nil
}
