package main

import tea "github.com/charmbracelet/bubbletea"

func main() {
	tea.NewProgram(Model{})
}

type Model struct {
	tea.Model
}

