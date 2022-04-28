package states

type Branch struct {
	StateMap map[string]State
	Name     string
	StartAt  string
}

func NewBranch(stateMap map[string]State, name, startAt string) *Branch {
	return &Branch{stateMap, name, startAt}
}

type ParallelState struct {
	branches      []*Branch
	name          string
	previousState string
	nextState     string
	end           bool
}

func NewParallelState(branches []*Branch, name, previous, next string, end bool) *ParallelState {
	return &ParallelState{branches, name, previous, next, end}
}

func (ps *ParallelState) Name() string {
	return ps.name
}

func (ps *ParallelState) Type() string {
	return Par
}

func (ps *ParallelState) Previous() string {
	return ps.previousState
}

func (ps *ParallelState) Next() string {
	return ps.nextState
}

func (ps *ParallelState) End() bool {
	return ps.end
}

func (ps *ParallelState) Branches() []*Branch {
	return ps.branches
}
