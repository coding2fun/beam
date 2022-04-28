package states

// Expression instance for evaluating conditions
type Expression struct {
	Variable   string
	MatchType  string
	MatchValue string
}

func NewExpression(variable, matchType, matchValue string) *Expression {
	return &Expression{Variable: variable, MatchType: matchType, MatchValue: matchValue}
}

// Condition interface contract
type Condition struct {
	Expressions   []Expression
	ConditionType string
	Next          string
}

func NewCondition(expressions []Expression, conditionType, next string) *Condition {
	return &Condition{Expressions: expressions, ConditionType: conditionType, Next: next}
}

type ConditionState struct {
	conditionList []*Condition
	previousState string
	nextState     string
	end           bool
}

func NewConditionState(conditionList []*Condition, previous, next string, end bool) *ConditionState {
	return &ConditionState{conditionList, previous, next, end}
}

func (cs *ConditionState) Name() string {
	return "Condition_No_Name"
}

func (cs *ConditionState) Type() string {
	return Cond
}

func (cs *ConditionState) Previous() string {
	return cs.previousState
}

func (cs *ConditionState) Next() string {
	return cs.nextState
}

func (cs *ConditionState) End() bool {
	return cs.end
}

func (cs *ConditionState) Conditions() []*Condition {
	return cs.conditionList
}
