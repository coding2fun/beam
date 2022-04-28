package states

type TaskState struct {
	name          string
	previousState string
	nextState     string
	end           bool
}

func NewTaskState(name, previous, next string, end bool) *TaskState {
	return &TaskState{name, previous, next, end}
}

func (ts *TaskState) Name() string {
	return ts.name
}

func (ts *TaskState) Type() string {
	return Task
}

func (ts *TaskState) Previous() string {
	return ts.previousState
}

func (ts *TaskState) Next() string {
	return ts.nextState
}

func (ts *TaskState) End() bool {
	return ts.end
}
