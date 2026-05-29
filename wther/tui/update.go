package tui

import (
	"wther/internal/weather"

	"charm.land/bubbles/v2/textinput"
	tea "charm.land/bubbletea/v2"
)

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	var cmds []tea.Cmd
	switch msg := msg.(type) {

	// Is it a key press?
	case tea.KeyPressMsg:

		// Cool, what was the actual key pressed?
		switch msg.String() {

		// These keys should exit the program.
		case "ctrl+c":
			return m, tea.Quit
		case "q":
			if m.state != StateLocationInput {
				return m, tea.Quit
			}

		case "l":

			if m.state != StateLocationInput {
				m.state = StateLocationInput
				m.locationInput.Focus()
				m.locationInput.Placeholder = m.queryOptions.Location
				m.locationInput.SetValue("")
				return m, textinput.Blink
			}

		case "enter":
			if m.state == StateLocationInput && m.locationInput.Value() != "" {
				m.queryOptions.Location = m.locationInput.Value()
				m.locationInput.Blur()
				m.state = StateTable

				return m, func() tea.Msg {
					return weather.GetForecast(m.queryOptions)

				}
			}

		case "esc":
			if m.state == StateLocationInput {
				m.locationInput.Blur()
				m.state = StateTable
				return m, nil
			}

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

	switch m.state {
	case StateLocationInput:
		m.locationInput, cmd = m.locationInput.Update(msg)
		cmds = append(cmds, cmd)
	case StateTable:
		m.table, cmd = m.table.Update(msg)
		cmds = append(cmds, cmd)
	}

	return m, tea.Batch(cmds...)
}
