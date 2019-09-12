package sql

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql" // Needs the specific driver
	"log"
	"os"
)

var hasGottenData = false

var username, password, database string

func getUser() string {
	if !hasGottenData {
		getData()
	}
	return username
}
func getDatabase() string {
	if !hasGottenData {
		getData()
	}
	return database
}
func getPassword() string {
	if !hasGottenData {
		getData()
	}
	return password
}
func getData() {
	username = os.Getenv("DB_USER")
	password = os.Getenv("DB_PASSWORD")
	database = os.Getenv("DB_DATABASE")
	hasGottenData = true
}

/*CreateDatabase Creates a new sql database in proper format*/
func CreateDatabase() {
	db, err := sql.Open("mysql",
		getUser()+":"+getPassword()+"@/")

	defer db.Close()

	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println("Opened sql")
	}
	_, err = db.Exec("CREATE DATABASE " + getDatabase())
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println("Created database")
	}
	_, err = db.Exec("USE " + getDatabase())
	if err != nil {
		panic(err)
	} else {
		fmt.Println("Using database")
	}
	_, err = db.Exec("CREATE TABLE blogposts (id INT NOT NULL AUTO_INCREMENT PRIMARY KEY, title TEXT, author TEXT, post TEXT, timestamp BIGINT)")
	if err != nil {
		panic(err)
	} else {
		fmt.Println("Created blog table")
	}
	_, err = db.Exec("CREATE TABLE tags (id INT NOT NULL AUTO_INCREMENT PRIMARY KEY, tag TEXT)")
	if err != nil {
		panic(err)
	} else {
		fmt.Println("Created tag table")
	}
	_, err = db.Exec("CREATE TABLE blogtags (tag INT, post INT)")
	if err != nil {
		panic(err)
	} else {
		fmt.Println("Created blogtag table")
	}

}

/*RemoveDatabase removes the sql database*/
func RemoveDatabase() {
	db, err := sql.Open("mysql",
		getUser()+":"+getPassword()+"@/")
	defer db.Close()
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println("Opened sql")
	}
	_, err = db.Exec("DROP DATABASE " + getDatabase())
	if err != nil {
		panic(err)
	} else {
		fmt.Println("dropped database")
	}
}
