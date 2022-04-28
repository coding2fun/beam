package states

type WaitState struct {
	name          string
	previousState string
	nextState     string
	end           bool
}

func NewWaitState(name, previous, next string, end bool) *WaitState {
	return &WaitState{name, previous, next, end}
}

func (ws *WaitState) Name() string {
	return ws.name
}

func (ws *WaitState) Type() string {
	return Wait
}

func (ws *WaitState) Previous() string {
	return ws.previousState
}

func (ws *WaitState) Next() string {
	return ws.nextState
}

func (ws *WaitState) End() bool {
	return ws.end
}
