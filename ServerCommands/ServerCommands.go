package ServerCommands

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

const (
	GET_WORDS = "getWords"
)

type result struct {
	id            int
	word          string
	transcription string
	translation   string
}

const (
	DB_USER     = "test"
	DB_PASSWORD = "123"
	DB_NAME     = "postgres"
)

func Commands(command string) []result {

	connStr := "user=test password=123 dbname=postgres"
	db, errorDb := sql.Open("postgres", connStr)
	if errorDb != nil {
		panic(errorDb)
	}

	results := []result{}
	if command == GET_WORDS {
		rows, err := db.Query("SELECT * FROM english_words ORDER BY RANDOM() LIMIT 5")
		if err != nil {
			//fmt.Println(errorSelect)
			panic(err)
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
