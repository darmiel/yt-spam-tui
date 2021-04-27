package main

import (
	"fmt"
	"github.com/charmbracelet/bubbletea"
	"github.com/emirpasic/gods/maps/treemap"
	"log"
)

type model struct {
	choices *treemap.Map
	cursor  int
}

var initialModel model

func init() {
	initialModel = model{
		choices: treemap.NewWithStringComparator(),
	}
	initialModel.choices.Put("Hallo", false)
	initialModel.choices.Put("das", false)
	initialModel.choices.Put("ist", false)
	initialModel.choices.Put("ein", false)
	initialModel.choices.Put("Test", false)
}

func (model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit
		case "up":
			if m.cursor > 0 {
				m.cursor--
			} else {
				m.cursor = m.choices.Size() - 1
			}
			break
		case "down":
			if m.cursor < m.choices.Size()-1 {
				m.cursor++
			} else {
				m.cursor = 0
			}
			break
		case "enter", " ":
			num := -1
			m.choices.Each(func(key interface{}, value interface{}) {
				check, ok := value.(bool)
				if !ok {
					return
				}
				num++
				if num != m.cursor {
					return
				}
				m.choices.Put(key, !check)
			})
			break
		}
		break
	}
	return m, nil
}

func (m model) View() string {
	s := "Select:\n\n"

	num := -1
	m.choices.Each(func(key interface{}, value interface{}) {
		choice, ok := key.(string)
		if !ok {
			return
		}
		check, ok := value.(bool)
		if !ok {
			return
		}
		num++
		cursor := "  "
		if m.cursor == num {
			cursor = " >"
		}

		checked := "[ ]"
		if check {
			checked = "[x]"
		}

		s += fmt.Sprintf("%s %s :: (%v) %s\n", cursor, checked, num, choice)
	})
	return s
}

func main() {
	prog := tea.NewProgram(initialModel)
	if err := prog.Start(); err != nil {
		log.Fatalln(err)
		return
	}
}
