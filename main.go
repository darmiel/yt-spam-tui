/*
package main

import (
	"fmt"
	"github.com/charmbracelet/bubbles/progress"
	"github.com/charmbracelet/bubbles/textinput"
	"github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"log"
	"time"
)

type model struct {
	i       textinput.Model
	p       *progress.Model
	percent float64
}

var initialModel model

type tickMsg time.Time

func init() {
	i := textinput.NewModel()
	i.CursorStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("63"))
	i.Width = 48
	i.Placeholder = "Daniel"
	i.CursorEnd()
	i.Focus()

	p, err := progress.NewModel(progress.WithScaledGradient("#FF7CCB", "#FDFF8C"))
	if err != nil {
		panic(err)
	}

	initialModel = model{i, p, 0.0}
}

func (model) Init() tea.Cmd {
	return textinput.Blink
}

func tickCmd() tea.Cmd {
	return tea.Tick(time.Second/20, func(t time.Time) tea.Msg {
		return tickMsg(t)
	})
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyCtrlC, tea.KeyEscape, tea.KeyEnter:
			return m, tea.Quit
		}
		break

	case tea.WindowSizeMsg:
		m.p.Width = msg.Width - 4*2 - 4
		return m, nil

	case tickMsg:
		m.percent += 1.0 / 20 * 2
		if m.percent > 1.0 {
			m.percent = 1.0
			return m, tea.Quit
		}
		return m, tickCmd()
	}
	var cmd tea.Cmd
	m.i, cmd = m.i.Update(msg)
	return m, cmd
}

func (m model) View() string {
	prc := float64(len(m.i.Value())) / 26.0
	if prc > 1.0 {
		prc = 1
	}
	return fmt.Sprintf("What's your name?\n\n%s\n\n  > [%s] <\n\n%s",
		m.i.View(),
		m.p.View(prc),
		"(esc to quit)")
}

func main() {
	prog := tea.NewProgram(initialModel)
	if err := prog.Start(); err != nil {
		log.Fatalln(err)
		return
	}
}
 */