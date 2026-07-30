package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	_ "github.com/lib/pq"
)

type Student struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Age   int    `json:"age"`
	Grade string `json:"grade"`
}

var db *sql.DB

func main() {
	var err error
	connStr := "host=localhost port=5432 user=postgres password=nintutanu1@ dbname=studentdb sslmode=disable"
	db, err = sql.Open("postgres", connStr)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer db.Close()

	createTable()

	http.HandleFunc("/students", handleStudents)

	fmt.Println("Server running on port 8080")
	http.ListenAndServe(":8080", nil)
}

func createTable() {
	query := `CREATE TABLE IF NOT EXISTS students (
		id SERIAL PRIMARY KEY,
		name TEXT,
		age INTEGER,
		grade TEXT
	)`
	db.Exec(query)
}

func handleStudents(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		getStudents(w)
		return
	}

	if r.Method == "POST" {
		createStudent(w, r)
		return
	}

	if r.Method == "PUT" {
		id, _ := strconv.Atoi(r.URL.Query().Get("id"))
		updateStudent(w, r, id)
		return
	}

	if r.Method == "DELETE" {
		id, _ := strconv.Atoi(r.URL.Query().Get("id"))
		deleteStudent(w, id)
		return
	}
}

func getStudents(w http.ResponseWriter) {
	rows, err := db.Query("SELECT id, name, age, grade FROM students")
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	defer rows.Close()

	var students []Student
	for rows.Next() {
		var s Student
		rows.Scan(&s.ID, &s.Name, &s.Age, &s.Grade)
		students = append(students, s)
	}

	json.NewEncoder(w).Encode(students)
}

func createStudent(w http.ResponseWriter, r *http.Request) {
	var s Student
	json.NewDecoder(r.Body).Decode(&s)

	db.QueryRow("INSERT INTO students (name, age, grade) VALUES ($1, $2, $3) RETURNING id", s.Name, s.Age, s.Grade).Scan(&s.ID)

	json.NewEncoder(w).Encode(s)
}

func updateStudent(w http.ResponseWriter, r *http.Request, id int) {
	var s Student
	json.NewDecoder(r.Body).Decode(&s)

	db.Exec("UPDATE students SET name=$1, age=$2, grade=$3 WHERE id=$4", s.Name, s.Age, s.Grade, id)

	fmt.Fprintf(w, "Student %d updated", id)
}

func deleteStudent(w http.ResponseWriter, id int) {
	db.Exec("DELETE FROM students WHERE id=$1", id)

	fmt.Fprintf(w, "Student %d deleted", id)
}
