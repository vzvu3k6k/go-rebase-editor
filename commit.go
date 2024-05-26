package rebase

import (
	"fmt"
)

type Commit struct {
	command Cmd
	title   string
	hash    string
}

func (c *Commit) SetCommand(cmd Cmd) {
	c.command = cmd
}

func (c *Commit) Render() []string {
	return []string{
		c.command.String(),
		c.hash,
		c.title,
	}
}

func (c *Commit) String() string {
	return fmt.Sprintf("%s %s %s", c.command, c.hash, c.title)
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
