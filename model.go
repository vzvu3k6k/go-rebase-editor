package rebase

import (
	"os"

	table "github.com/calyptia/go-bubble-table"
	tea "github.com/charmbracelet/bubbletea"
	"golang.org/x/term"
)

type Model struct {
	commits []Commit
	table   table.Model
}

func (m Model) Init() tea.Cmd {
	return nil
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch keypress := msg.String(); keypress {
		case "q", "enter":
			return m, tea.Quit
		case "esc", "ctrl+c":
			m.commits = []Commit{}
			return m, tea.Quit
		case "e":
			if c, ok := m.table.SelectedRow().(Commit); ok {
				c.SetCommand(CmdEdit)
				m.commits[m.table.Cursor()] = c

				rows := make([]table.Row, len(m.commits))
				for i, v := range m.commits {
					rows[i] = v
				}

				m.table.SetRows(rows)
				return m, nil
			}
		}
	case tea.WindowSizeMsg:
		m.table.SetSize(
			msg.Width,
			msg.Height,
		)
	}

	var cmd tea.Cmd
	m.table, cmd = m.table.Update(msg)
	return m, cmd
}

func (m Model) View() string {
	return m.table.View()
}

func NewModel(commits []Commit) Model {
	w, h, err := term.GetSize(int(os.Stdout.Fd()))
	if err != nil {
		w = 80
		h = 24
	}
	tbl := table.New([]string{"Command", "ID", "Title"}, w, h)

	rows := make([]table.Row, len(commits))
	for i, v := range commits {
		rows[i] = v
	}
	tbl.SetRows(rows)

	return Model{
		commits: commits,
		table:   tbl,
	}
}
