package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/urfave/cli/v2"
	"os"
)

const initSql = `CREATE TABLE articles (
    id INT UNSIGNED AUTO_INCREMENT NOT NULL,
    title text DEFAULT NULL,
    url text DEFAULT NULL,
    info text DEFAULT NULL,
    c_date date DEFAULT NULL,
    uuid text DEFAULT NULL,
    PRIMARY KEY(id)
)
DEFAULT CHARACTER SET utf8mb4
COLLATE utf8mb4_unicode_ci
ENGINE = InnoDB;`

func CreateMySQLTable(c *cli.Context) error {
	fmt.Println("Open DB connection")
	db, err := sql.Open("mysql", os.Getenv("MYSQL_DB"))
	defer db.Close()
	if err != nil {
		fmt.Println("error: %s", err.Error())
		return err
	}

	_, err = db.Exec(initSql)
	if err != nil {
		fmt.Println("error: %s", err.Error())
		return err
	}
	fmt.Println("New table created")
	return nil
}
