package rebase

import (
	"testing"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/google/go-cmp/cmp"
	"gotest.tools/v3/assert"
)

func TestNoChange(t *testing.T) {
	m := prepareModel([]Commit{
		{command: CmdPick, hash: "00000000", title: "foo"},
	})

	m, cmd := applyKeyMsgs(m, []tea.KeyMsg{
		{Type: tea.KeyEnter},
	})
	assert.Equal(t, cmd(), tea.Quit())

	assertCommits(t, m, []Commit{
		{command: CmdPick, hash: "00000000", title: "foo"},
	})
}

func TestCancel(t *testing.T) {
	m := prepareModel([]Commit{
		{command: CmdPick, hash: "00000001", title: "1st"},
	})

	m, cmd := applyKeyMsgs(m, []tea.KeyMsg{
		{Type: tea.KeyEsc},
	})
	assert.Equal(t, cmd(), tea.Quit())

	// returns empty commits to abort rebasing
	assertCommits(t, m, []Commit{})
}

func TestCommandChange(t *testing.T) {
	m := prepareModel([]Commit{
		{command: CmdPick, hash: "00000001", title: "1st"},
		{command: CmdPick, hash: "00000002", title: "2nd"},
	})

	m, _ = applyKeyMsgs(m, []tea.KeyMsg{
		{Type: tea.KeyDown},
		{Type: tea.KeyRunes, Runes: []rune{'e'}},
	})

	assertCommits(t, m, []Commit{
		{command: CmdPick, hash: "00000001", title: "1st"},
		{command: CmdEdit, hash: "00000002", title: "2nd"},
	})
}

func prepareModel(commits []Commit) tea.Model {
	m := NewModel(commits)
	m.Init()
	return m
}

func applyKeyMsgs(m tea.Model, msgs []tea.KeyMsg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	for _, v := range msgs {
		m, cmd = m.Update(v)
	}
	return m, cmd
}

func assertCommits(t *testing.T, m tea.Model, commits []Commit) {
	model := m.(Model)
	assert.DeepEqual(t, model.commits, commits, cmp.AllowUnexported(Commit{}))
}
