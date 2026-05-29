package tui

import (
	"fmt"
	"strconv"
	"time"
	"wther/internal/weather"

	"charm.land/bubbles/v2/table"
	tea "charm.land/bubbletea/v2"
	"charm.land/lipgloss/v2"
)

var tableBaseStyle = lipgloss.NewStyle().
	BorderStyle(lipgloss.NormalBorder()).
	BorderForeground(lipgloss.Color("240")) // grey 35

var (
	NormalRowStyle = lipgloss.NewStyle()

	CurrentHourRowStyle = lipgloss.NewStyle().
				Bold(true).
				Foreground(lipgloss.Color("78")) // aquamarine 3
)

func (m Model) View() tea.View {
	if m.err != nil {
		return tea.NewView(fmt.Sprintf("\nWe had some trouble: %v\n\n", m.err))
	}

	s := "wther - what's the weather looking like?\n"

	switch m.state {

	case StateLocationInput:
		s += "Changing locations:\n"
		s += m.locationInput.View() + "\n\n"
		s += "(press enter to submit, esc to cancel)\n"
	case StateTable:

		s += m.forecast.String() + "\n"

		s += tableBaseStyle.Render(m.table.View()) + "\n  " + m.table.HelpView() + "\n"

	}

	return tea.NewView(s)

}

func CreateTable(f weather.Forecast) (table.Model, error) {

	columns := []table.Column{
		{Title: "Day", Width: 8},
		{Title: "Hour", Width: 8},
		{Title: "Temp", Width: 10},
		{Title: "Condition", Width: 32},
	}

	currentHourIndex := 0

	// row generation
	var rows []table.Row
	var dateTimeLayout = "2006-01-02 15:04"
	var dateLayout = "Jan 2"
	var timeLayout = "15:04"

	for i, fDay := range f.ForecastObject.ForecastList {
		for j, fHour := range fDay.Hour {

			rowDateTime, err := time.Parse(dateTimeLayout, fHour.Time)
			if err != nil {
				return table.Model{}, err
			}
			rowDate := rowDateTime.Format(dateLayout)
			rowTime := rowDateTime.Format(timeLayout)

			style := NormalRowStyle

			if rowDateTime.Hour() == time.Now().Hour() && rowDateTime.Day() == time.Now().Day() {
				style = CurrentHourRowStyle
				currentHourIndex = i*24 + j
			}

			rows = append(rows, table.Row{
				style.Render(rowDate),
				style.Render(rowTime),
				style.Render(strconv.FormatFloat(fHour.TempC, 'f', -1, 64)),
				style.Render(fHour.Condition.Text)})
		}
	}

	t := table.New(
		table.WithColumns(columns),
		table.WithRows(rows),
		table.WithFocused(true),
		table.WithHeight(10),
		table.WithWidth(64),
		table.WithKeyMap(table.DefaultKeyMap()),
	)

	t.MoveDown(currentHourIndex)
	return t, nil
}
