package rebase

import (
	"bufio"
	"io"
	"strings"
)

func ParseCommits(r io.Reader) Commits {
	var commits Commits

	s := bufio.NewScanner(r)
	for s.Scan() {
		line := s.Text()
		if line == "" || line[0] == '#' {
			break
		}
		commits = append(commits, parseCommit(line))
	}

	return commits
}

func parseCommit(s string) Commit {
	a := strings.SplitN(s, " ", 3)

	// For now, assumes that command is always "pick"
	return Commit{command: CmdPick, hash: a[1], title: a[2]}
}
