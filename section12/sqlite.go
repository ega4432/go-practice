package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

var DbConnection *sql.DB

type Person struct {
	Name string
	Age  int
}

func main() {
	DbConnection, _ := sql.Open("sqlite3", "./section12/example.sql")
	defer DbConnection.Close()
	cmd := `CREATE TABLE IF NOT EXISTS person(
				name STRING,
				age INT)`
	_, err := DbConnection.Exec(cmd)
	if err != nil {
		log.Fatalln(err)
	}

	// Insert
	cmd = "INSERT INTO person (name, age) VALUES (?, ?)"
	_, err = DbConnection.Exec(cmd, "Mike", 24)
	if err != nil {
		log.Fatalln(err)
	}

	// Update
	cmd = "UPDATE person SET age = ? WHERE name = ?"
	_, err = DbConnection.Exec(cmd, 25, "Mike")
	if err != nil {
		log.Fatal(err)
	}

	// Select
	cmd = "SELECT * FROM person"
	rows, _ := DbConnection.Query(cmd)
	defer rows.Close()
	var pp []Person
	for rows.Next() {
		var p Person
		err := rows.Scan(&p.Name, &p.Age)
		if err != nil {
			log.Println(err)
		}
		pp = append(pp, p)
	}
	err = rows.Err()
	if err != nil {
		log.Fatalln(err)
	}
	for _, p := range pp {
		fmt.Println(p.Name, p.Age)
	}

	cmd = "SELECT * FROM person WHERE age = ?"
	row := DbConnection.QueryRow(cmd, 2000)
	var p Person
	err = row.Scan(&p.Name, &p.Age)
	if err != nil {
		if err == sql.ErrNoRows {
			log.Println("No row")
		} else {
			log.Println(err)
		}
	}
	fmt.Println(p.Name, p.Age)

	// Delete
	cmd = "DELETE FROM person WHERE name = ?"
	_, err = DbConnection.Exec(cmd, "Nancy")
	if err != nil {
		log.Fatalln(err)
	}

	// SQL injection
	tableName := "person"
	// ex : tableName := "person; INSERT INTO person (name, age) VALUES ('Mr.X', 100);"
	cmd = fmt.Sprintf("SELECT * FROM %s", tableName)
	rows, _ := DbConnection.Query(cmd)
	defer rows.Close()
	var pp []Person
	for rows.Next() {
		var p Person
		err := rows.Scan(&p.Name, &p.Age)
		if err != nil {
			log.Println(err)
		}
		pp = append(pp, p)
	}
	err = rows.Err()
	if err != nil {
		log.Fatalln(err)
	}
	for _, p := range pp {
		fmt.Println(p.Name, p.Age)
	}
}
