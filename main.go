package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"html/template"
	"log"
	"net/http"
	"time"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/", welcome)
	router.HandleFunc("/signin", signin)
	router.HandleFunc("/createUser", createUser)
	router.HandleFunc("/signup", signup)

	http.ListenAndServe(":8090", router)
}

func signup(writer http.ResponseWriter, request *http.Request) {
	tmpl := template.Must(template.ParseFiles("register.html"))
	tmpl.Execute(writer, nil)
}

func createUser(writer http.ResponseWriter, request *http.Request) {
	username := request.FormValue("exampleInputEmail1")
	password := request.FormValue("exampleInputPassword1")
	db, _ := sql.Open("mysql", "root:welcome@(127.0.0.1:3306)/information?parseTime=true")
	result, _ := db.Exec(`INSERT INTO users (username, password, created_at) VALUES (?, ?, ?)`, username, password, time.Now())
	if result == nil {
		log.Fatal("Error")
		fmt.Fprintf(writer, "Oops! try again later.")
	} else {
		fmt.Fprintf(writer, "Congratulations "+username+" You are successfully regsitered.")
	}

}

func signin(writer http.ResponseWriter, request *http.Request) {
	username := request.FormValue("exampleInputEmail1")
	password := request.FormValue("exampleInputPassword1")
	db, _ := sql.Open("mysql", "root:welcome@(127.0.0.1:3306)/information?parseTime=true")

	passwordFromDB := ""
	query := `SELECT password FROM users WHERE username = ?`
	err := db.QueryRow(query, username).Scan(&passwordFromDB)
	print(err)
	if password == passwordFromDB {
		fmt.Fprintf(writer, "Congratulations "+username+" You are successfully signed in.")
	} else {
		fmt.Fprintf(writer, "Oops! Username and password did not match.")
	}

}

func welcome(writer http.ResponseWriter, request *http.Request) {
	tmpl := template.Must(template.ParseFiles("index.html"))
	tmpl.Execute(writer, nil)
}
