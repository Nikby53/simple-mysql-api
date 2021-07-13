package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

type text struct {
	Film_id     int    `json:"film___id"`
	Title       string `json:"title"`
	Description string `json:"description"`
}

func main() {

	http.HandleFunc("/films", getFilms)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func getFilms(w http.ResponseWriter, r *http.Request) {
	db, err := sql.Open("mysql", "root:12325066Qq@/sakila")

	if err != nil {
		panic(err)
	}
	defer db.Close()
	rows, err := db.Query("SELECT * FROM sakila.film_text;")
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	texts := []text{}

	for rows.Next() {
		p := text{}
		err := rows.Scan(&p.Film_id, &p.Title, &p.Description)
		if err != nil {
			panic(err)
		}
		texts = append(texts, p)
	}
	vivod, err := json.MarshalIndent(&texts, "\t", "")
	if err != nil {
		panic(err)
	}
	fmt.Fprintf(w, string(vivod))
}
