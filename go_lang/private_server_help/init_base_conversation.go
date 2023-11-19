package private_server_help

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/google/uuid"
	"os"
)

type privateApiConversation struct {
	Id      int
	uuid    string
	History History
}

type History []Message

type Message struct {
	Sender string
	Body   string
}

func InitConversation(task string) (string, error) {
	host := os.Getenv("TASK_SERVER_URL")
	id, err := generateConversationId()
	if err != nil {
		return "", err
	}
	url := host + "/" + task + "/" + id
	return url, nil
}

func generateConversationId() (string, error) {
	uuidValue := uuid.New()
	var conversation privateApiConversation
	conversation.uuid = uuidValue.String()
	err := insertDataToMySQL(conversation)
	if err != nil {
		return "", err
	}
	return conversation.uuid, nil
}

func insertDataToMySQL(conversation privateApiConversation) error {
	fmt.Println("Open DB connection")
	db, err := sql.Open("mysql", os.Getenv("MYSQL_DB"))
	defer db.Close()
	if err != nil {
		fmt.Println("error:", err.Error())
		return err
	}
	sql := "INSERT INTO private_api_conversation(uuid) VALUES (?)"
	res, err := db.Exec(sql, conversation.uuid)

	if err != nil {
		return err
	}
	lastId, err := res.LastInsertId()

	if err != nil {
		return err
	}

	fmt.Printf("The last inserted row id: %d\n", lastId)
	return nil
}
