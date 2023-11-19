package conversation_help

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"os"
)

func FindConversationByUUID(uuid string) (PrivateApiConversation, error) {
	fmt.Println("Open DB connection")
	db, err := sql.Open("mysql", os.Getenv("MYSQL_DB"))
	defer db.Close()
	if err != nil {
		return PrivateApiConversation{}, err
	}
	conversation := PrivateApiConversation{}
	row := db.QueryRow("SELECT * FROM private_api_conversation WHERE uuid=?", uuid)
	err = row.Scan(&conversation.Id, &conversation.Uuid, &conversation.History)
	if err != nil {
		return PrivateApiConversation{}, err
	}
	return conversation, nil
}

func UpdateConversation(conversation PrivateApiConversation) error {
	fmt.Println("Open DB connection")
	db, err := sql.Open("mysql", os.Getenv("MYSQL_DB"))
	defer db.Close()
	if err != nil {
		return err
	}

	sql := "UPDATE private_api_conversation SET history=? WHERE uuid=?"
	res, err := db.Exec(sql, conversation.History, conversation.Uuid)

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
