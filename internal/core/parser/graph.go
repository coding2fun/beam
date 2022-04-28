package parser

import "coding2fun.in/beam/internal/core/states"

type WorkflowGraph struct {
	WorkflowName   string
	IsAsync        bool
	StartAt        string
	ResultVariable string
	StatesMap      map[string]states.State
}

type WorkflowBuilder interface {
	WorkflowName(string) WorkflowBuilder
	IsAsync(bool) WorkflowBuilder
	StartState(string) WorkflowBuilder
	ResultVariable(string) WorkflowBuilder
	States(map[string]states.State) WorkflowBuilder
	Build() WorkflowGraph
}

type wBuilder struct {
	workflowName   string
	isAsync        bool
	startAt        string
	resultVariable string
	statesMap      map[string]states.State
}

func (wb *wBuilder) WorkflowName(name string) WorkflowBuilder {
	wb.workflowName = name
	return wb
}

func (wb *wBuilder) IsAsync(async bool) WorkflowBuilder {
	wb.isAsync = async
	return wb
}

func (wb *wBuilder) StartState(startAt string) WorkflowBuilder {
	wb.startAt = startAt
	return wb
}

func (wb *wBuilder) ResultVariable(resultVar string) WorkflowBuilder {
	wb.resultVariable = resultVar
	return wb
}

func (wb *wBuilder) States(statesMap map[string]states.State) WorkflowBuilder {
	wb.statesMap = statesMap
	return wb
}

func (wb *wBuilder) Build() WorkflowGraph {
	return WorkflowGraph{
		WorkflowName:   wb.workflowName,
		IsAsync:        wb.isAsync,
		StartAt:        wb.startAt,
		ResultVariable: wb.resultVariable,
		StatesMap:      wb.statesMap,
	}
}

func NewWorkflowBuilder() WorkflowBuilder {
	return &wBuilder{}
}
