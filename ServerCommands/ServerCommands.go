package ServerCommands

import (
	"database/sql"
	"fmt"
)

const (
	GET_WORDS = "getWords"
)

type result struct {
	id int
	word string
	transcription string
	translation string
}

const (
	DB_USER = "test"
	DB_PASSWORD = "123"
	DB_NAME = "postgres"
)

func Commands(command string) []result {

	dbinfo := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable",
		DB_USER, DB_PASSWORD, DB_NAME)
	db, errorDb := sql.Open("postgres", dbinfo)
	if errorDb != nil {
		fmt.Println(nil)
	}

	results := []result{}
	if command == GET_WORDS {
		rows, errorSelect := db.Query("SELECT * FROM english_words ORDER BY RANDOM() LIMIT 5")

		if errorSelect != nil {
			//fmt.Println(errorSelect)
			panic(errorSelect)
		}

		for rows.Next() {
			r := result{}
			err := rows.Scan(&r.id, &r.word, &r.transcription, &r.translation)
			if err != nil {
				fmt.Println(err)
				continue
			}
			results = append(results, r)
		}

		for _, r := range results {
			fmt.Println(&r.id, &r.word, &r.transcription, &r.translation)
		}
	}
	defer db.Close()
	return results
}
