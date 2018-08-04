package state

type StateType int

const (
	Task StateType = iota
)

type State struct {
	Name string
	Type StateType
	Next *State
}
