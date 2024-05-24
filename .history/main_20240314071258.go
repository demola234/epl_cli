package main

import tea "github.com/charmbracelet/bubbletea"


type 

type Task struct {
	title string
	status string

}

func main() {
	tea.NewProgram(Model{})
}

type Model struct {
	tea.Model
}

