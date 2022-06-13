package rebase

import (
	"fmt"
	"io"

	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

var docStyle = lipgloss.NewStyle().Margin(1, 2)

var itemStyle = list.NewDefaultItemStyles()

type itemDelegate struct{}

func (d itemDelegate) Render(w io.Writer, m list.Model, index int, item list.Item) {
	c, ok := item.(Commit)
	if !ok {
		return
	}

	l := fmt.Sprintf("%c %s %s\n", c.Command(), c.Hash(), c.Title())
	if index == m.Index() {
		fmt.Fprintf(w, itemStyle.SelectedTitle.Render(l))
	} else {
		fmt.Fprintf(w, itemStyle.NormalTitle.Render(l))
	}
}

func (d itemDelegate) Height() int  { return 1 }
func (d itemDelegate) Spacing() int { return 0 }

func (d itemDelegate) Update(msg tea.Msg, m *list.Model) tea.Cmd { return nil }

type Model struct {
	list list.Model
}

func (m Model) Init() tea.Cmd {
	return nil
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch keypress := msg.String(); keypress {
		case "ctrl+c":
			return m, tea.Quit
		case "e":
			idx := m.list.Index()
			if c, ok := m.list.SelectedItem().(Commit); ok {
				c.SetCommand('e')
				m.list.SetItem(idx, c)
			}
		case "enter":
			return m, tea.Quit
		}
	case tea.WindowSizeMsg:
		h, v := docStyle.GetFrameSize()
		m.list.SetSize(msg.Width-h, msg.Height-v)
	}

	var cmd tea.Cmd
	m.list, cmd = m.list.Update(msg)
	return m, cmd
}

func (m Model) View() string {
	return docStyle.Render(m.list.View())
}

func NewModel() Model {
	var items []list.Item
	for _, v := range commits {
		items = append(items, v)
	}

	list := list.New(items, itemDelegate{}, 0, 0)
	list.SetShowTitle(false)
	list.SetShowPagination(false)

	return Model{
		list: list,
	}
}
