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

func (t Task) Title() string {
	return t.title
}

func (t Task) Description() string {
	return t.description
}




func main() {
	tea.NewProgram(Model{})
}

type Model struct {
	tea.Model
}

