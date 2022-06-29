package rebase

import (
	"os"

	table "github.com/calyptia/go-bubble-table"
	tea "github.com/charmbracelet/bubbletea"
	"golang.org/x/term"
)

var keyToCmd = map[string]Cmd{
	"e": CmdEdit,
	"r": CmdReword,
	"p": CmdPick,
	"s": CmdSquash,
	"f": CmdFixup,
	"d": CmdDrop,
}

type Model struct {
	commits Commits
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
			m.commits = Commits{}
			return m, tea.Quit
		case "left":
			if !m.table.CursorIsAtTop() {
				m.commits.MoveUp(m.table.Cursor())
				m.applyCommits()
				m.table.GoUp()
			}
			return m, nil
		case "right":
			if !m.table.CursorIsAtBottom() {
				m.commits.MoveDown(m.table.Cursor())
				m.applyCommits()
				m.table.GoDown()
			}
			return m, nil
		default:
			if cmd, ok := keyToCmd[keypress]; ok {
				m.commits[m.table.Cursor()].SetCommand(cmd)
				m.applyCommits()
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

func (m *Model) applyCommits() {
	rows := make([]table.Row, len(m.commits))
	for i, v := range m.commits {
		rows[i] = v
	}
	m.table.SetRows(rows)
}

func NewModel(commits Commits) Model {
	w, h, err := term.GetSize(int(os.Stdout.Fd()))
	if err != nil {
		w = 80
		h = 24
	}
	tbl := table.New([]string{"Command", "ID", "Title"}, w, h)

	m := Model{
		commits: commits,
		table:   tbl,
	}
	m.applyCommits()

	return m
}
