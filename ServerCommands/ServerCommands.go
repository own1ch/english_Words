package ServerCommands

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"strings"
)

const (
	GET_WORDS    = "\bgetWords"
	REGISTRATION = "registration"
	CONN         = "user=postgres dbname=postgres sslmode=disable"
)

type result struct {
	id            int
	Word          string
	Transcription string
	Translation   string
}

const (
	DB_USER     = "test"
	DB_PASSWORD = "123"
	DB_NAME     = "postgres"
)

func GetStruct() []result {
	return []result{}
}

func GetWords(countOfWords string, login string) []string {
	db := ConnToDatabase()
	results := []result{}
	var info []string
	var scanId []int
	var idUser int
	idUserQuery, err := db.Query("SELECT id FROM users WHERE login = $1", login)
	for idUserQuery.Next() {
		err := idUserQuery.Scan(&idUser)
		if err != nil {
			panic(err)
		}
	}

	rows, err := db.Query(
		"SELECT word, transcription, translation, id FROM word_user wu "+
			"FULL OUTER JOIN english_words ew on wu.word_id = ew.id "+
			"WHERE wu.user_id = $1 and ew.id IS NULL OR wu.word_id IS NULL ORDER BY RANDOM() LIMIT $2", idUser, countOfWords)
	if err != nil {
		panic(err)
	}

	for rows.Next() {
		r := result{}
		err := rows.Scan(&r.Word, &r.Transcription, &r.Translation, &r.id)
		if err != nil {
			fmt.Println(err)
			continue
		}
		info = append(info, r.Word)
		info = append(info, r.Transcription)
		info = append(info, r.Translation)
		scanId = append(scanId, r.id)
	}

	for _, r := range results {
		fmt.Println(r)
	}

	AddReferences(*db, login, scanId)

	defer db.Close()
	return info
}

func AddReferences(db sql.DB, userId int, scanId []int) {
	for _, wordId := range scanId {
		_, err := db.Exec("INSERT INTO word_user(word_id, user_id) VALUES ($1, $2)", wordId, userId)
		if err != nil {
			panic(err)
		}
	}
}

func Registration(netData string) bool {
	db := ConnToDatabase()
	info := strings.Split(netData, "|")
	_, err := db.Exec("INSERT INTO users(login, password, name) VALUES ($1, $2, $3)",
		info[0], info[1], info[2])
	if err != nil {
		fmt.Println(err)
		return false
	} else {
		return true
	}
}

func Login(netData string) bool {
	db := ConnToDatabase()
	info := strings.Split(netData, "|")
	rows, err := db.Query("SELECT login FROM users WHERE login=$1 AND password=$2",
		info[0], info[1])
	if err != nil {
		panic(err)
	}
	if rows != nil {
		return true
	} else {
		return false
	}
}

func ConnToDatabase() *sql.DB {
	db, errorDb := sql.Open("postgres", CONN)
	if errorDb != nil {
		panic(errorDb)
	}
	return db
}
