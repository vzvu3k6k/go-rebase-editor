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
	s := fmt.Sprintf("%s\t%s\t%s", c.command, c.hash, c.title)
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

func (cmd Cmd) String() string {
	switch cmd {
	case CmdPick:
		return "pick"
	case CmdReword:
		return "reword"
	case CmdEdit:
		return "edit"
	case CmdSquash:
		return "squash"
	case CmdFixup:
		return "fixup"
	case CmdDrop:
		return "drop"
	default:
		return string(cmd)
	}
}
