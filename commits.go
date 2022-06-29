package rebase

type Commits []Commit

func (cs *Commits) MoveUp(idx int) {
	if idx <= 0 {
		return
	}
	(*cs)[idx], (*cs)[idx-1] = (*cs)[idx-1], (*cs)[idx]
}

func (cs *Commits) MoveDown(idx int) {
	if idx >= len(*cs)-1 {
		return
	}
	(*cs)[idx], (*cs)[idx+1] = (*cs)[idx+1], (*cs)[idx]
}
