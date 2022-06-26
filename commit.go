package rebase

import (
	"fmt"
	"io"

	table "github.com/calyptia/go-bubble-table"
)

type Commit struct {
	command Cmd
	title   string
	hash    string
}

func (c *Commit) SetCommand(cmd Cmd) {
	c.command = cmd
}

func (c Commit) Render(w io.Writer, model table.Model, index int) {
	s := string(c.command)
	if index == model.Cursor() {
		s = model.Styles.SelectedRow.Render(s)
	}
	fmt.Fprintln(w, s)
}

type Cmd rune

const (
	CmdPick   Cmd = 'p'
	CmdReword     = 'r'
	CmdEdit       = 'e'
	CmdSquash     = 's'
	CmdFixup      = 'f'
	CmdDrop       = 'd'
)
