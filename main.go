package main

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"

	_ "github.com/lib/pq"
)
type User struct {
	Name string `json:"name"`
}
func main() {
 db,err := sql.Open(
	"postgres", "host=localhost port=5433 user=postgres password=BEZH0909 dbname=postgres sslmode=disable",
 )
 if err != nil{
	log.Fatal(err)
 }
 defer db.Close()

 if err := db.Ping(); err !=nil{
	log.Fatal("It didn't connect", err)
 }
 http.HandleFunc("POST /userss", func(w http.ResponseWriter, r *http.Request) {
	var user User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil{
		http.Error(w, "JSON is incorrect!", http.StatusBadRequest)
		return
	}
	if user.Name == " "{
		http.Error(w, "The name can be empty!", http.StatusBadRequest)
		return 
	}
     var id int
	 err := db.QueryRow(
		"INSERT INTO userss (name) VALUES ($1) RETURNING id",user.Name,
     ).Scan(&id)
	 if err != nil{
		http.Error(w, "хатои базаи маълумот!", http.StatusInternalServerError)
		log.Println(err)
		return
	 }
	 w.Header().Set("Content-Type", "application/json")
	 w.WriteHeader(http.StatusCreated)
	 json.NewEncoder(w).Encode(map[string]interface{}{
		"massage": "The name is registred succesfully!",
		"id": id,
		"name": user.Name,
	 })
 })
 log.Println("Server started http://localhost:8080")
 log.Fatal(http.ListenAndServe(":8080", nil ))
}