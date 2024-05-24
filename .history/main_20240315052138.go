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

func initialModel() model {
	s := spinner.New()
	s.Spinner = spinner.Dot
	s.Style = lipgloss.NewStyle().Foreground(lipgloss.Color("205"))
	return model{spinner: s, loading: true} // Start in loading state
}

// Init function for the Bubble Tea program
func (m model) Init() tea.Cmd {

	return m.spinner.Tick
}

func (m model) View() string {
	if m.loading {
		return m.spinner.View() // Show the spinner while loading
	}
	// Once loading is complete, show the table
	return baseStyle.Render(m.table.View()) + "\n"
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg.(type) {
	case tea.KeyMsg:
		// Handle key presses here. You might want to allow quitting with 'q' even when loading.
	case tea.WindowSizeMsg:
		// Handle window size changes.
	}

	if m.loading {
		var cmds []tea.Cmd
		m.spinner, cmd := m.spinner.Update(msg)
		cmds = append(cmds, cmd)

		// Check if loading is done. This could be done by checking a condition or receiving a custom message type that you emit after loading data.
		// For illustration, let's say we have a custom message type `dataLoadedMsg`:
		switch msg.(type) {
		case dataLoadedMsg:
			m.loading = false
			// Process the loaded data here
		}

		return m, tea.Batch(cmds...)
	} else {
		// Normal update logic here.
		var cmd tea.Cmd
		m.table, cmd = m.table.Update(msg)
		return m, cmd
	}
}

func main() {
	m := initialModel()

	go func() {
		eplTableData, err := api.GetData() // Simulate data loading
		if err != nil {
			// Handle error
			return
		}
		// Assume processing of `eplTableData` here

		// Send a custom message (dataLoadedMsg) to notify the program that data loading is complete
		program.Send(dataLoadedMsg{data: processedData})
	}()

	columns := []table.Column{
		{Title: "Rank", Width: 10}, // Example of setting a column width
		{Title: "Team", Width: 20},
		{Title: "Points", Width: 15},
		{Title: "Goals", Width: 15},
		{Title: "GD", Width: 10},
		{Title: "Matches", Width: 15},
	}

	// Assuming eplTableData.Data.Table is a slice of models.EplTableRow
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

	m := model{table: t}
	if _, err := tea.NewProgram(m).Run(); err != nil {
		fmt.Println("Error running program:", err)
		os.Exit(1)
	}
}
