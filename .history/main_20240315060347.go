package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"

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
	table     table.Model
	spinner   spinner.Model
	loading   bool
	responses int           // how many responses we've received
	sub       chan struct{} // channel to receive activity notifications
}

func initialModel() (spinners spinner.Model, loading bool) {
	s := spinner.NewModel()
	s.Spinner = spinner.Dot
	s.Style = lipgloss.NewStyle().Foreground(lipgloss.Color("205"))
	return s, true
}

// Init function for the Bubble Tea program
func (m model) Init() tea.Cmd {
	return tea.Batch(fetchEplData, listenForActivity(m.sub))
}

func (m model) View() string {
	if m.loading {
		s := fmt.Sprintf("\n %s Loading League Table \n\n Press any key to exit\n", m.spinner.View())
		return s
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
				// Show all the row details
				tea.Printf("Let's show the details for %s:\n", m.table.SelectedRow()[1], ""),
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
	case responseMsg:
		m.responses++ // record external activity
	}
	m.table, cmd = m.table.Update(msg)
	return m, cmd
}

func main() {
	spin, loaded := initialModel()
	m := model{
		spinner: spin,
		loading: loaded,
		sub:     make(chan struct{}), // Initialize the channel
	}
	if err := tea.NewProgram(m).Start(); err != nil {
		fmt.Println("Error running program:", err)
		os.Exit(1)
	}
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
			fmt.Println("Error getting data:", err)
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
	case <-time.After(5 * time.Second):
		// If the goroutine doesn't complete within 5 seconds, return an error message.
		return errorMsg{fmt.Errorf("timeout: failed to fetch data within 5 seconds")}
	}
}

// A message used to indicate that activity has occurred. In the real world (for
// example, chat) this would contain actual data.
type responseMsg struct{}

// Simulate a process that sends events at an irregular interval in real time.
// In this case, we'll send events on the channel at a random interval between
// 100 to 1000 milliseconds. As a command, Bubble Tea will run this
// asynchronously.
func listenForActivity(sub chan struct{}) tea.Cmd {
	return func() tea.Msg {
		// Simulate network delay by sleeping for a random duration
		time.Sleep(time.Millisecond * time.Duration(rand.Int63n(2000)+1000))
		sub <- struct{}{}
		return nil
	}
}
