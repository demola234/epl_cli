package main

import (
	"github.com/charmbracelet/bubbles/table"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

var baseStyle = lipgloss.NewStyle().
	BorderStyle(lipgloss.NormalBorder()).
	BorderForeground(lipgloss.Color("240"))

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) View() string {
	return baseStyle.Render(m.table.View()) + "\n"
}

type model struct {
	table table.Model
}

func main() {

}
