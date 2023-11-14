package tools

import (
	"encoding/json"
	"github.com/sashabaranov/go-openai"
	"github.com/sashabaranov/go-openai/jsonschema"
)

const ToDoType = "ToDo"

type ToDoData struct {
	Task string
}

func processToDoQuery(functionCall *openai.FunctionCall) (string, error) {
	args := functionCall.Arguments
	toDoJson := ToDoData{}
	err := json.Unmarshal([]byte(args), &toDoJson)
	if err != nil {
		return "", err
	}

	toolResponse := `{"tool":"Calendar","desc":"` + toDoJson.Task + `"}`
	return toolResponse, nil
}

func getToDoFunctionDef() openai.FunctionDefinition {
	return openai.FunctionDefinition{
		Name:        ToDoType,
		Description: "Run this function when user suggested task for todo list without a date.",
		Parameters: jsonschema.Definition{
			Type: jsonschema.Object,
			Properties: map[string]jsonschema.Definition{
				"task": {
					Type:        jsonschema.String,
					Description: "Task description",
				},
			},
			Required: []string{"country"},
		},
	}
}
