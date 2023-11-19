package private_api_tasks_help

import "strings"

func BuildReply(answer string) string {
	transformedAnswer := strings.Replace(answer, `"`, `\"`, -1)
	transformedAnswer = strings.Replace(transformedAnswer, "\n", "", -1)
	reply := `{"reply":"` + transformedAnswer + `"}`
	return reply
}
