package state

import (
	"time"
)

type ExecutionStatusType int

const (
	NotStarted ExecutionStatusType = iota
	Executing
	Succeeded
	Failed
)

type ExecutionStatusHistory struct {
	DateTime *time.Time
	Type     ExecutionStatusType
}

type StateExecution struct {
	Definition      *State
	Input           string
	Output          string
	StartedOn       *time.Time
	EndedOn         *time.Time
	ExecutionStatus ExecutionStatusType
	History         []*ExecutionStatusHistory
}
