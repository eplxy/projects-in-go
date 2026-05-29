package tui

import (
	"wther/internal/weather"

	"charm.land/bubbles/v2/table"
	tea "charm.land/bubbletea/v2"
)

type FetchStatus int

const (
	None FetchStatus = iota
	Fetching
	Error
	Success
)

type Model struct {
	err         error
	fetchStatus FetchStatus // none, fetching, error, success

	forecast weather.Forecast
	// mode ViewMode

	queryOptions weather.ForecastQueryOptions
	// config AppConfig
	table table.Model
}

func NewModel(queryOptions weather.ForecastQueryOptions) Model {
	return Model{queryOptions: queryOptions}
}

func (m Model) Init() tea.Cmd {

	return func() tea.Msg {
		// wrapped to allow passing the queryOptions
		return weather.GetForecast(m.queryOptions)
	}
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
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
		table, err := CreateTable(msg.Body)
		if err != nil {
			m.err = err
		}
		m.table = table
	case weather.ForecastErrMsg:
		m.err = msg.Err
	}

	// check if applicable to table
	m.table, cmd = m.table.Update(msg)

	return m, cmd
}
