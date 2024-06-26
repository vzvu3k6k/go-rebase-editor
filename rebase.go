package rebase

import (
	"errors"
	"os"

	tea "github.com/charmbracelet/bubbletea"
)

func Run(args []string) error {
	if len(args) < 1 {
		return errors.New("no input file specified")
	}
	inputPath := args[0]

	commits, err := readCommits(inputPath)
	if err != nil {
		return err
	}

	commits, err = editCommits(commits)
	if err != nil {
		return err
	}

	if err := writeCommits(inputPath, commits); err != nil {
		return err
	}

	return nil
}

func readCommits(filepath string) (Commits, error) {
	file, err := os.Open(filepath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	return ParseCommits(file), nil
}

func editCommits(commits Commits) ([]Commit, error) {
	var m tea.Model = NewModel(commits)
	p := tea.NewProgram(m, tea.WithAltScreen())
	m, err := p.Run()
	if err != nil {
		return nil, err
	}
	return m.(Model).commits, nil
}

func writeCommits(filepath string, commits Commits) error {
	file, err := os.OpenFile(filepath, os.O_WRONLY|os.O_TRUNC, os.ModePerm)
	if err != nil {
		return err
	}

	for _, c := range commits {
		if _, err := file.WriteString(c.String() + "\n"); err != nil {
			return err
		}
	}

	return nil
}
