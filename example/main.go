package main

import (
	"encoding/json"
	"fmt"
	"log"

	odt "github.com/efimovalex/odt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

var schema = `DROP TABLE IF EXISTS time_table; CREATE TABLE time_table (time TIME NULL, date DATE NULL);`

type dateTypes struct {
	Time odt.Time `sql:"time" json:"time"`
	Date odt.Date `sql:"date" json:"date"`
}

var db *sqlx.DB

func init() {
	var err error
	// Open up your database connection.
	db, err = sqlx.Connect("mysql", "username:password@tcp(127.0.0.1:3306)/test?multiStatements=true")
	if err != nil {
		log.Fatalln(err)

		return
	}

	db.MustExec(schema)
}

func main() {
	var dt dateTypes
	var encodedJSON = `{"time":"23:12:01", "date": "2008-01-03"}`

	if err := json.Unmarshal([]byte(encodedJSON), &dt); err != nil {
		log.Println(err)

		return
	}

	// time and date columns are of type TIME and ,respectively, DATE in mysql
	_, err := db.NamedExec(`INSERT INTO time_table (time,date) VALUES (:time,:date)`, &dt)
	if err != nil {
		log.Println(err)

		return
	}

	var readDTs []dateTypes

	err = db.Select(&readDTs, "SELECT * FROM time_table")
	if err != nil {
		fmt.Println(err)

		return
	}

	result, err := json.Marshal(readDTs)
	if err != nil {
		log.Println(err)

		return
	}

	// [{"time":"23:12:01","date":"2008-01-03"}]
	fmt.Println(string(result))
}
