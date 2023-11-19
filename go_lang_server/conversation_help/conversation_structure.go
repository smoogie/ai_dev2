package conversation_help

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
)

type PrivateApiConversation struct {
	Id      int
	Uuid    string
	History History
}

type History []Message

type Message struct {
	Sender string
	Body   string
}

const SenderUser = "user"
const SenderAssistant = "assistant"

func (h History) Value() (driver.Value, error) {
	return json.Marshal(h) // return json marshalled value
}

func (h *History) Scan(v interface{}) error {
	switch tv := v.(type) {
	case []uint8:
		return json.Unmarshal([]byte(tv), &h) // can't remember the specifics, but this may be needed
	}
	return errors.New("unsupported type")
}
