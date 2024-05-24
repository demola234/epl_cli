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
    name         string
    matchesPlayed int
    wins         int
    draws        int
    goalsFor     int
    goalsAgainst int
    points       int
}

type model struct {
    teams []team
}



var (
    titleStyle = lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color("63")).PaddingBottom(1)
    itemStyle  = lipgloss.NewStyle().Foreground(lipgloss.Color("229"))
    altItemStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("245"))
    thinLineStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("241")).SetString("│").Padding(0, 1)
    headerStyle = lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color("212"))
)


