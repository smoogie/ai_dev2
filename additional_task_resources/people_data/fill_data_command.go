package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/urfave/cli/v2"
	"io/ioutil"
	"os"
)

type Person struct {
	Id                               int
	Imie                             string
	Nazwisko                         string
	Wiek                             int
	O_mnie                           string
	Ulubiona_postac_z_kapitana_bomby string
	Ulubiony_serial                  string
	Ulubiony_film                    string
	Ulubiony_kolor                   string
}

func Command(c *cli.Context) error {
	fmt.Println("Reading json file")
	jsonFile, err := os.Open("people.json")
	if err != nil {
		fmt.Printf("error: %s", err.Error())
		return err
	}
	byteValue, _ := ioutil.ReadAll(jsonFile)
	var people []Person
	err = json.Unmarshal(byteValue, &people)
	if err != nil {
		fmt.Println("error:", err.Error())
		return err
	}

	fmt.Println("Open DB connection")
	db, err := sql.Open("mysql", os.Getenv("MYSQL_DB"))
	defer db.Close()
	if err != nil {
		fmt.Println("error:", err.Error())
		return err
	}

	fmt.Println("Starting adding data")
	for i := 0; i < len(people); i++ {
		person := people[i]
		person.Id = i + 1
		err = insertDataToMySQL(db, person)
		if err != nil {
			fmt.Printf("error: %s", err.Error())
			return err
		}
	}

	return nil
}

func insertDataToMySQL(db *sql.DB, person Person) error {
	sql := "INSERT INTO people(imie, nazwisko, wiek, o_mnie, ulubiona_postac_z_kapitana_bomby, ulubiony_serial, ulubiony_film, ulubiony_kolor) VALUES (?, ?, ?, ?, ?, ?, ?, ?)"
	res, err := db.Exec(sql, person.Imie, person.Nazwisko, person.Wiek, person.O_mnie, person.Ulubiona_postac_z_kapitana_bomby, person.Ulubiony_serial, person.Ulubiony_film, person.Ulubiony_kolor)

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
