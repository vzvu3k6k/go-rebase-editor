package rebase

import (
	"os"

	"github.com/charmbracelet/bubbles/table"
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

var _ tea.Model = (*Model)(nil)

func NewModel(commits Commits) Model {
	w, h, err := term.GetSize(int(os.Stdout.Fd()))
	if err != nil {
		w = 80
		h = 24
	}

	m := Model{
		commits: commits,
		table:   buildTable(w, h),
	}
	m.applyCommits()

	return m
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
			if m.table.Cursor() > 0 {
				m.commits.MoveUp(m.table.Cursor())
				m.applyCommits()
				m.table.MoveUp(1)
			}
			return m, nil
		case "right":
			if m.table.Cursor() < len(m.table.Rows())-1 {
				m.commits.MoveDown(m.table.Cursor())
				m.applyCommits()
				m.table.MoveDown(1)
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
		m.table = buildTable(msg.Width, msg.Height)
		m.applyCommits()
		m.table.UpdateViewport()
	}

	m.table, _ = m.table.Update(msg)
	return m, nil
}

func (m Model) View() string {
	return m.table.View()
}

func (m *Model) applyCommits() {
	rows := make([]table.Row, len(m.commits))
	for i, v := range m.commits {
		rows[i] = v.Render()
	}
	m.table.SetRows(rows)
}

const (
	commandWidth = 8
	idWidth      = 9
	headerHeight = 1
)

func buildTable(w, h int) table.Model {
	return table.New(
		table.WithColumns([]table.Column{
			{Title: "Command", Width: commandWidth},
			{Title: "ID", Width: idWidth},
			{Title: "Title", Width: w - commandWidth - idWidth},
		}),
		table.WithFocused(true),
		table.WithHeight(h-headerHeight),
	)
}
