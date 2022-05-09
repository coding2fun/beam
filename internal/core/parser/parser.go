package parser

import (
	"io/ioutil"
	"os"

	"github.com/rs/zerolog/log"
	"github.com/tidwall/gjson"
)

const json = `{"name":{"first":"Janet","last":"Prichard"},"age":47}`

// Parser contract, Generates [WorkflowGraph] from workflow definition
type Parser interface {
	CreateWorkflowGraph(file *os.File) (*WorkflowGraph, error)
}

type jsonParser struct {
}

func NewJsonParser() Parser {
	return &jsonParser{}
}

func (jp *jsonParser) CreateWorkflowGraph(file *os.File) (*WorkflowGraph, error) {
	log.Info().Msg("in jsonParser: CreateWorkflowGraph")

	workflowDefinition, err := ioutil.ReadAll(file)
	if err != nil {

	}

	jsonContent := string(workflowDefinition)
	jsonNode := gjson.Parse(jsonContent)

	validateFields(jsonNode)

	value := gjson.Get(json, "name.last")
	println(value.String())
	return nil, nil
}

func validateFields(jsonNode gjson.Result) {
	log.Info().Msg("")
}
