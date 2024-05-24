package main

import (
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

func (m *Model) initList() {
	m.list = list.New([]list.Item{}, list.NewDefaultDelegate())

	
}

func main() {

}
