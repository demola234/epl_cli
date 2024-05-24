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
	description string

}


// Implement List.Item Interface
func (t Task) FilterValue() string {
	return t.title
}
func main() {
	tea.NewProgram(Model{})
}

type Model struct {
	tea.Model
}

