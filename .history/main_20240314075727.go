package main

import (
	"fmt"
	"os"

	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
)

type status int

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
	list list.Model
	err  error
}

func New() *Model {
	return &Model{}

}

// TODO: CALL THIS ON tea.Window.Draw
func (m *Model) initLists(width, height int) {
	m.list = list.New([]list.Item{}, list.NewDefaultDelegate(), width, height)
	m.list.Title = "Tasks"
	m.list.SetItems([]list.Item{
		Task{title: "Task 1", status: todo, description: "This is a task"},
		Task{title: "Task 2", status: inProgress, description: "This is a task"},
		Task{title: "Task 3", status: done, description: "This is a task"},
	})
}

// Make Model a bubble.Model
func (m *Model) Init() tea.Cmd {
	return nil
}

func (m *Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.initList(msg.Width, msg.Height)
	}
	var cmd tea.Cmd
	m.list, cmd = m.list.Update(msg)
	return m, cmd
}

func (m *Model) View() string {
	return m.list.View()
}

func main() {
	m := New()
	p := tea.NewProgram(m)
	if err := p.Start(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

}
