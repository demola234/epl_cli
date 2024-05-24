package main

import (
	"fmt"
	"os"

	"github.com/charmbracelet/bubbles/spinner"
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
	spinner  spinner.Model
	eplTable models.EplTable
	loading  bool
}

func initialModel() (spinners spinner.Model, loading bool) {
	s := spinner.NewModel()
	s.Spinner = spinner.Dot
	s.Style = lipgloss.NewStyle().Foreground(lipgloss.Color("205"))
	return s, true
}

// Init function for the Bubble Tea program
func (m model) Init() tea.Cmd {
	return fetchEplData
}

func (m model) View() string {
	if m.loading {
		return m.spinner.View() + baseStyl("Fetching Data...") + "\n"
	}
	return baseStyle.Render(m.table.View()) + "\n"
}

type errorMsg struct {
	err error
}

type updatedTableMsg struct {
	t table.Model
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
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
	case errorMsg:
		// Handle error
		fmt.Println("Error occurred:", msg.err)
		m.loading = false // Ensure loading state is set to false on error
	case updatedTableMsg:
		// Update table
		m.table = msg.t
		m.loading = false // Ensure loading state is set to false after updating the table
	}
	return m, nil
}

func main() {
	spin, loaded := initialModel()
	m := model{spinner: spin, loading: loaded}
	if err := tea.NewProgram(m).Start(); err != nil {
		fmt.Println("Error running program:", err)
		os.Exit(1)
	}
}

func fetchEplData() tea.Msg {
	eplTableData, err := api.GetData()
	if err != nil {
		fmt.Println("Error getting data:", err)
		return errorMsg{err}
	}

	columns := []table.Column{
		{Title: "Rank", Width: 10}, // Example of setting a column width
		{Title: "Team", Width: 20},
		{Title: "Points", Width: 15},
		{Title: "Goals", Width: 15},
		{Title: "GD", Width: 10},
		{Title: "Matches", Width: 15},
	}

	rows := make([]table.Row, 0)
	for _, row := range eplTableData.Data.Table {
		rows = append(rows, table.Row{row.Rank, row.Name, row.Points, row.GoalsScored, row.GoalDiff, row.Matches})
	}

	t := table.New(
		table.WithColumns(columns),
		table.WithRows(rows),
		table.WithFocused(true),
		table.WithHeight(20),
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

	return updatedTableMsg{t}
}
