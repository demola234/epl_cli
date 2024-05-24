package main

import (
	"fmt"
	"os"
	"sort"

	tea "github.com/charmbracelet/bubbletea"
)

type team struct {
	name   string
	points int
}

type model struct {
	teams []team
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	return m, nil
}

func (m model) View() string {
	// Sort the teams slice based on points, descending
	sort.SliceStable(m.teams, func(i, j int) bool {
		return m.teams[i].points > m.teams[j].points
	})

	s := "Premier League Table\n\n"
	for i, team := range m.teams {
		s += fmt.Sprintf("%d. %s, %d pts\n", i+1, team.name, team.points)
	}
	return s

}

func main() {
	initialModel := model{
		teams: []team{
			{"Liverpool", 82},
			{"Manchester City", 80},
			{"Chelsea", 75},
			{"Manchester United", 70},
			// Add more teams as desired...
		},
	}

	p := tea.NewProgram(initialModel)
	if err := p.Start(); err != nil {
		fmt.Printf("Alas, there's been an error: %v", err)
		os.Exit(1)
	}
}
