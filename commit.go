package rebase

type Commit struct {
	command Cmd
	title   string
	hash    string
}

func (c Commit) Description() string { return c.Hash() }
func (c Commit) Command() Cmd        { return c.command }
func (c Commit) Title() string       { return c.title }
func (c Commit) Hash() string        { return c.hash }
func (c Commit) FilterValue() string { return c.title }

func (c *Commit) SetCommand(cmd Cmd) {
	c.command = cmd
}

func (c Commit) Render(w io.Writer, model table.Model, index int) {
	s := string(c.Command())
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

var commits = []Commit{
	{
		command: 'p',
		title:   "initial commit",
		hash:    "deadbeef",
	},
	{
		command: 'p',
		title:   "2nd commit",
		hash:    "deadbeef",
	},
}
