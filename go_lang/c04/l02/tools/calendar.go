package tools

import (
	"encoding/json"
	"github.com/sashabaranov/go-openai"
	"github.com/sashabaranov/go-openai/jsonschema"
)

const CalendarType = "Calendar"

type CalendarData struct {
	Task string
	Date string
}

func processCalendarQuery(functionCall *openai.FunctionCall) (string, error) {
	args := functionCall.Arguments
	calendarJson := CalendarData{}
	err := json.Unmarshal([]byte(args), &calendarJson)
	if err != nil {
		return "", err
	}

	toolResponse := `{"tool":"Calendar","desc":"` + calendarJson.Task + `","date":"` + calendarJson.Date + `"}`
	return toolResponse, nil
}

func getCalendarFunctionDef() openai.FunctionDefinition {
	return openai.FunctionDefinition{
		Name:        CalendarType,
		Description: "Run this function when user suggested task with a date.",
		Parameters: jsonschema.Definition{
			Type: jsonschema.Object,
			Properties: map[string]jsonschema.Definition{
				"task": {
					Type:        jsonschema.String,
					Description: "Task description",
				},
				"date": {
					Type:        jsonschema.String,
					Description: "Date of the task in the format YYYY-MM-DD",
				},
			},
			Required: []string{"country"},
		},
	}
}
