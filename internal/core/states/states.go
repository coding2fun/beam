package states

// State Type
const (
	Task = "Task"
	Cond = "Condition"
	Wait = "Wait"
	Par  = "Parallel"
)

// Condition Type
const (
	Simple = "Simple"
	And    = "And"
	Or     = "Or"
)

// Match Type
const (
	StringEquals    = "StringEquals"
	StringNotEquals = "StringNotEquals"
)

// State interface contract
type State interface {
	Name() string
	Type() string
	Previous() string
	Next() string
	End() bool
}
