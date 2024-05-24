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



