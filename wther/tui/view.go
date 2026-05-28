package tui

import (
	"fmt"

	tea "charm.land/bubbletea/v2"
)

func (m model) View() tea.View {
	if m.err != nil {
		return tea.NewView(fmt.Sprintf("\nWe had some trouble: %v\n\n", m.err))
	}

	s := "wther - what's the weather looking like?\n\n"

	s += m.forecast.String()

	return tea.NewView(s)

}
