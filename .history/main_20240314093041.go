package main

import (
    "fmt"
    "os"
    "sort"

    tea "github.com/charmbracelet/bubbletea"
    "github.com/charmbracelet/lipgloss"
)

var (
    titleStyle  = lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color("#FFA500")).Background(lipgloss.Color("#333")).Padding(0, 1)
    teamStyle   = lipgloss.NewStyle().Foreground(lipgloss.Color("#FFF")).Padding(0, 1)
    pointsStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("#FFD700")).Padding(0, 1)
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
    sort.SliceStable(m.teams, func(i, j int) bool {
        return m.teams[i].points > m.teams[j].points
    })

    s := titleStyle.Render("Premier League Table\n\n")
    for i, team := range m.teams {
        line := fmt.Sprintf("%d. %s, %s\n", i+1, teamStyle.Render(team.name), pointsStyle.Render(fmt.Sprintf("%d pts", team.points)))
        s += line
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
