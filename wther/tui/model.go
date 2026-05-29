package tui

import (
	"wther/internal/weather"

	"charm.land/bubbles/v2/table"
	"charm.land/bubbles/v2/textinput"
	tea "charm.land/bubbletea/v2"
)

type SessionState int

const (
	StateTable SessionState = iota
	StateLocationInput
)

type Model struct {
	err error

	forecast     weather.Forecast
	state        SessionState
	queryOptions weather.ForecastQueryOptions
	// config AppConfig
	table         table.Model
	locationInput textinput.Model
}

func NewModel(queryOptions weather.ForecastQueryOptions) Model {
	locationInput := textinput.New()
	locationInput.Placeholder = queryOptions.Location
	locationInput.CharLimit = 100
	locationInput.SetWidth(32)

	return Model{queryOptions: queryOptions, state: StateTable, locationInput: locationInput}
}

func (m Model) Init() tea.Cmd {

	return func() tea.Msg {
		// wrapped to allow passing the queryOptions
		return weather.GetForecast(m.queryOptions)
	}
}
