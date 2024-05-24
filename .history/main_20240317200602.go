package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"time"

	"github.com/charmbracelet/bubbles/spinner"
	"github.com/charmbracelet/bubbles/table"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/demola234/api"
	"github.com/demola234/models"
	"github.com/demola234/ui"
	"github.com/spf13/cobra"
)

type model struct {
	table     table.Model
	fixtures  table.Model
	spinner   spinner.Model
	loading   bool
	sub       chan struct{}
	showTable bool
}

func initialModel() (spinners spinner.Model, loading bool) {
	s := spinner.New()
	s.Spinner = spinner.Jump
	s.Style = lipgloss.NewStyle().Foreground(lipgloss.Color("205"))
	s.Style.Render(spinner.Jump.FPS.String())

	return s, true
}

// Init function for the Bubble Tea program
func (m model) Init() tea.Cmd {
	return tea.Batch(fetchEplData)
}

func (m model) View() string {
	if m.loading {
		return fmt.Sprintf("\n %s Loading League Table \n\n Press q to exit\n", m.spinner.View())
	} else if m.showTable {
		return ui.BaseStyle.Render(m.table.View()) + "\n"
	} else {
		return ui.BaseStyle.Render(m.fixtures.View()) + "\n"
	}
}

type errorMsg struct {
	err error
}

type updatedTableMsg struct {
	t table.Model
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	switch msg := msg.(type) {
	case tea.KeyMsg:
		// Handle key presses
		switch msg.String() {
		case "esc":
			if m.showTable {
				m.table.Blur()
			} else {
				m.fixtures.Blur()
			}
		case "q", "ctrl+c":
			return m, tea.Quit
		case "enter":
			if m.showTable {
				// Handle club selection
				selectedClub := m.table.SelectedRow()[1] // Assuming the team name is in the second column
				// Fetch fixtures for the selected club
				fetchFixtures(selectedClub)
				m.showTable = false // Hide league table
			} else {
				m.showTable = true // Show league table
			}
		case "backspace":
			m.showTable = true // Show league table
		}
	case errorMsg:
		// Handle error
		fmt.Println("Error occurred:", msg.err)
		m.loading = false // Ensure loading state is set to false on error
	case updatedTableMsg:
		if m.showTable {
			m.table = msg.t
		} else {
			m.fixtures = msg.t // Set the fixtures table
		}
		m.loading = false // Ensure loading state is set to false after updating the table
	}
	// Update the appropriate table based on the current view
	if m.showTable {
		m.table, cmd = m.table.Update(msg)
	} else {
		m.fixtures, cmd = m.fixtures.Update(msg)
	}
	return m, cmd
}

func fetchEplData() tea.Msg {
	done := make(chan bool)
	defer close(done)

	// Channel to receive data from goroutine
	ch := make(chan models.EplTable)

	// Start a goroutine to fetch data
	go func() {
		eplTableData, err := api.GetData()
		if err != nil {
			fmt.Println("Error getting data:")
			// Signal completion before returning
			done <- true
			return
		}
		// Send fetched data
		ch <- eplTableData
	}()

	// Wait for data or timeout
	select {
	case eplTableData := <-ch:
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
	case <-done:
		// This means the goroutine has completed execution.
		// You can add handling for this scenario if needed.
		return nil
	case <-time.After(100 * time.Second):
		// If the goroutine doesn't complete within 5 seconds, return an error message.
		return errorMsg{fmt.Errorf("timeout: failed to fetch data within 5 seconds")}
	}
}
func fetchFixtures(teamName string) tea.Cmd {
	return func() tea.Msg {
		done := make(chan bool)
		defer close(done)

		// Channel to receive data from goroutine
		ch := make(chan models.FixturesEntity)

		// Start a goroutine to fetch data
		go func() {
			fixtureData, err := api.GetFixtureData(teamName)
			tea.

			if err != nil {
				fmt.Println("Error getting data:", err)
				// Signal completion before returning
				done <- true
				return
			}
			// Send fetched data
			ch <- fixtureData
		}()

		// Wait for data or timeout
		select {
		case fixtureData := <-ch:
			columns := []table.Column{
				{Title: "Home Team", Width: 20}, // Example of setting a column width
				{Title: "Away Team", Width: 20},
				{Title: "Location", Width: 15},
				{Title: "Date", Width: 15},
				{Title: "Time", Width: 10},
				{Title: "League", Width: 15},
			}

			rows := make([]table.Row, 0)
			for _, row := range fixtureData.Data.Fixtures {
				rows = append(rows, table.Row{row.HomeName, row.AwayName, row.Location, row.Date, row.Time, row.Competition.Name})
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
		case <-done:
			// This means the goroutine has completed execution.
			// You can add handling for this scenario if needed.
			return nil
		case <-time.After(100 * time.Second):
			// If the goroutine doesn't complete within 5 seconds, return an error message.
			return errorMsg{fmt.Errorf("timeout: failed to fetch data within 5 seconds")}
		}
	}

}

func getFootballAscii() string {
	// Read the file contents into a byte slice
	data, err := ioutil.ReadFile("football_ascii.txt")
	if err != nil {
		log.Fatalf("Error reading file: %v", err)
	}

	// Convert the byte slice to a string
	content := string(data)

	// Print the content
	return content
}

func fetchCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "fetch",
		Short: "Fetch EPL data",
		Long:  "Fetches the latest English Premier League data and displays it.",
		Run: func(cmd *cobra.Command, args []string) {
			spin, loaded := initialModel()
			m := model{
				spinner: spin,
				loading: loaded,
				sub:     make(chan struct{}), // Initialize the channel
			}
			err := func() error {
				_, err := tea.NewProgram(m).Run()
				return err
			}()

			if err != nil {
				fmt.Println("Error running program:", err)
				os.Exit(1)
			}
		},
	}
}

func rootCmd() *cobra.Command {
	cmd := &cobra.Command{Use: "eplcli"}
	cmd.AddCommand(fetchCmd())
	return cmd
}

func main() {
	cmd := rootCmd()
	if err := cmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
