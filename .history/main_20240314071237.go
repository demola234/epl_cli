package main

import tea "github.com/charmbracelet/bubbletea"


type Task struct {
	title 
}

func main() {
	tea.NewProgram(Model{})
}

type Model struct {
	tea.Model
}

