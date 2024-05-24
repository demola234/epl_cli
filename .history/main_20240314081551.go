package main

import (
	"fmt"
	"os"

	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type status int

const divisor = 4
const (
	todo status = iota
	inProgress
	done
)

// Custom List Item
type Task struct {
	title       string
	status      status
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

// Model
type Model struct {
	focused status
	lists   []list.Model
	err     error
	loaded  bool
}

func New() *Model {
	return &Model{}

}

// TODO: CALL THIS ON tea.Window.Draw
func (m *Model) initLists(width, height int) {
	defaultList := list.New([]list.Item{}, list.NewDefaultDelegate(), width/divisor, height)
	defaultList
	m.lists = []list.Model{defaultList, defaultList, defaultList}

	m.lists[todo].Title = "Tasks"
	m.lists[todo].SetItems([]list.Item{
		Task{title: "Task 1", status: todo, description: "This is a task"},
		Task{title: "Task 2", status: inProgress, description: "This is a task"},
		Task{title: "Task 3", status: done, description: "This is a task"},
	})

	m.lists[inProgress].Title = "In Progress"
	m.lists[inProgress].SetItems([]list.Item{
		Task{title: "Stay Cool!!1", status: todo, description: "Happy Birthday"},
	})

	m.lists[done].Title = "Done"
	m.lists[done].SetItems([]list.Item{
		Task{title: "Stay Cool!!1", status: todo, description: "Happy Birthday"},
	})
}

// Make Model a bubble.Model
func (m *Model) Init() tea.Cmd {
	return nil
}

func (m *Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		if !m.loaded {
			m.loaded = true
		}
		m.initLists(msg.Width, msg.Height)
	}
	var cmd tea.Cmd
	m.lists[m.focused], cmd = m.lists[m.focused].Update(msg)
	return m, cmd
}

func (m *Model) View() string {
	if m.loaded {
		return lipgloss.JoinHorizontal(lipgloss.Left,
			m.lists[todo].View(),
			m.lists[inProgress].View(),
			m.lists[done].View(),
		)
	} else {
		return "Loading.."
	}

}

func main() {
	m := New()
	p := tea.NewProgram(m)
	if err := p.Start(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

}
