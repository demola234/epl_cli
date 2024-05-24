package main

import (
    "fmt"
    "os"
    "sort"

    tea "github.com/charmbracelet/bubbletea"
    "github.com/charmbracelet/lipgloss"
)



var (
    titleStyle = lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color("#FFA500")).Background(lipgloss.Color("#333")).Padding(0, 1)
    teamStyle  = lipgloss.NewStyle().Foreground(lipgloss.Color("#FFF")).Padding(0, 1)
    pointsStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("#FFD700")).Padding(0, 1)
)


func (m model) View() string {
    // Sort the teams slice based on points, descending
    sort.SliceStable(m.teams, func(i, j int) bool {
        return m.teams[i].points > m.teams[j].points
    })

    s := titleStyle.Render("Premier League Table\n\n")
    for i, team := range m.teams {
        line := fmt.Sprintf("%d. %s,
		