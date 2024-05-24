package main

import tea "github.com/charmbracelet/bubbletea"


type Task struct {
	title string
	
}

func main() {
	tea.NewProgram(Model{})
}

type Model struct {
	tea.Model
}

