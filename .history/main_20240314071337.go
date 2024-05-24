package main

import tea "github.com/charmbracelet/bubbletea"


type status int


const (
	todo status = iota
	inProgress
	done
)

type Task struct {
	title string
	status status
	

}

func main() {
	tea.NewProgram(Model{})
}

type Model struct {
	tea.Model
}

