package main

import (
	"fmt"
	"os"
	"sort"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type team struct {
	name          string
	matchesPlayed int
	wins          int
	draws         int
	goalsFor      int
	goalsAgainst  int
	points        int
}

type model struct {
	teams []team
}

var (
	titleStyle    = lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color("63")).PaddingBottom(1)
	itemStyle     = lipgloss.NewStyle().Foreground(lipgloss.Color("229"))
	altItemStyle  = lipgloss.NewStyle().Foreground(lipgloss.Color("245"))
	thinLineStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("241")).SetString("│").Padding(0, 1)
	headerStyle   = lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color("212"))
)

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
		case 
}

func (m model) View() string {
	var b strings.Builder

	// Table header
	fmt.Fprintf(&b, "%s\n", titleStyle.Render("Premier League Table"))
	header := fmt.Sprintf("%s %s %s %s %s %s %s %s",
		headerStyle.Render("Team"),
		headerStyle.Render("MP"),
		headerStyle.Render("W"),
		headerStyle.Render("D"),
		headerStyle.Render("GF"),
		headerStyle.Render("GA"),
		headerStyle.Render("GD"),
		headerStyle.Render("Pts"),
	)
	fmt.Fprintln(&b, header)

	// Sort teams by points
	sort.SliceStable(m.teams, func(i, j int) bool {
		return m.teams[i].points > m.teams[j].points
	})

	// Table body
	for i, t := range m.teams {
		style := itemStyle
		if i%2 == 1 {
			style = altItemStyle // Alternate color for every second item
		}

		gd := t.goalsFor - t.goalsAgainst // Calculate Goal Difference
		teamLine := fmt.Sprintf("%s %d %d %d %d %d %+d %d",
			style.Render(t.name),
			t.matchesPlayed,
			t.wins,
			t.draws,
			t.goalsFor,
			t.goalsAgainst,
			gd,
			t.points,
		)

		// Apply the thin line style i % 2 times
		if i%2 == 0 {
			fmt.Fprint(&b, thinLineStyle.Render(" "))
		}

		fmt.Fprintln(&b, teamLine)
	}

	return b.String()
}

func main() {
	initialModel := model{
		teams: []team{
			{"Liverpool", 38, 30, 3, 85, 33, 93},
			{"Manchester City", 38, 29, 6, 102, 35, 93},
			{"Chelsea", 38, 25, 9, 76, 36, 84},
			{"Manchester United", 38, 19, 15, 73, 56, 72},
			// Add more teams as needed...
		},
	}

	p := tea.NewProgram(initialModel)
	if err := p.Start(); err != nil {
		fmt.Printf("Alas, there's been an error: %v", err)
		os.Exit(1)
	}
}
