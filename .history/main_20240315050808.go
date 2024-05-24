package main

import (
	"fmt"
	"os"

	"github.com/charmbracelet/bubbles/table"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/demola234/api"
	"github.com/demola234/models"
)

var baseStyle = lipgloss.NewStyle().
	BorderStyle(lipgloss.NormalBorder()).
	BorderForeground(lipgloss.Color("240"))

type model struct {
	table    table.Model
	eplTable models.EplTable // Assuming this is correctly defined in your models package
}

// Init function for the Bubble Tea program
func (m model) Init() tea.Cmd {
	return nil
}

// View function renders the UI to the screen
func (m model) View() string {
	return baseStyle.Render(m.table.View()) + "\n"
}

// Update function processes input and updates the model
func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "esc":
			if m.table.Focused() {
				m.table.Blur()
			} else {
				m.table.Focus()
			}
		case "q", "ctrl+c":
			return m, tea.Quit
		case "enter":
			return m, tea.Batch(
				tea.Printf("Let's go to %s!", m.table.SelectedRow()[1]),
			)
		}
	}
	m.table, cmd = m.table.Update(msg)
	return m, cmd
}

func main() {
	eplTableData, err := api.GetData()
	if err != nil {
		fmt.Println("Error fetching data:", err)
		os.Exit(1)
	}

	columns := []table.Column{{title: "Rank"}, {title:"Team"},{title: "Points"}, {title:"Goals"}, {title:"GD"}, {title:"Matches"}}

	// Assuming eplTableData.Data.Table is a slice of models.EplTableRow
	rows := make([]table.Row, 0)
	for _, row := range eplTableData.Data.Table {
		rows = append(rows, table.Row{row.Rank, row.Name, row.Points, row.GoalsScored, row.GoalDiff, row.Matches})
	}

	t := table.New(
		table.WithColumns(columns),
		table.WithRows(rows),
		table.WithFocused(true),
		table.WithHeight(7),
	)

	s := table.DefaultStyles()
	s.Header = s.Header.
		BorderStyle(lipgloss.NormalBorder()).
		BorderForeground(lipgloss.Color("240")).
		BorderBottom(true).
		Bold(false)
	s.Selected = s.Selected.
		Foreground(lipgloss.Color("229")).
		Background(lipgloss.Color("57")).
		Bold(false)
	t.SetStyles(s)

	m := model{table: t}
	if _, err := tea.NewProgram(m).Run(); err != nil {
		fmt.Println("Error running program:", err)
		os.Exit(1)
	}
}
