package rebase

import (
	"strings"
	"testing"

	"github.com/google/go-cmp/cmp"
	"gotest.tools/v3/assert"
)

func TestParseCommits(t *testing.T) {
	todo := `pick deadbeef Hello, world
pick facefeed Goes on

# Rebase 917936c onto 171a6e6 (2 commands)`

	commits := ParseCommits(strings.NewReader(todo))

	assert.DeepEqual(t, commits, Commits{
		{
			command: CmdPick,
			hash:    "deadbeef",
			title:   "Hello, world",
		},
		{
			command: CmdPick,
			hash:    "facefeed",
			title:   "Goes on",
		},
	}, cmp.AllowUnexported(Commit{}))
}
