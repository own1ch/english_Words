package serverCommand

import (
	"../logs"
	"../struct"
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"strconv"
	"time"
)

const (
	GET_WORDS    = "\bgetWords"
	REGISTRATION = "registration"
	CONN         = "user=postgres dbname=postgres sslmode=disable"
)

const (
	DB_USER     = "test"
	DB_PASSWORD = "123"
	DB_NAME     = "postgres"
)

func GetStruct() []_struct.English {
	return []_struct.English{}
}

func GetWords(countOfWords string, login string) []_struct.English {
	db := ConnToDatabase()
	results := []_struct.English{}
	var word []_struct.English
	var scanId []int
	var userId int
	var userTable bool
	idUserQuery, err := db.Query("SELECT id, table_type FROM users WHERE login = $1", login)
	for idUserQuery.Next() {
		err := idUserQuery.Scan(&userId, &userTable)
		if err != nil {
			logs.UpdateFile(time.Now().Format("15:04:23") + err.Error())
		}
	}

	var tab string
	if userTable {
		tab = "id < 18145"
	} else {
		tab = "id > 18144"
	}

	rows, err := db.Query(
		"SELECT word, transcription, translation, id FROM word_user wu "+
			"FULL OUTER JOIN english_words ew on wu.word_id = ew.id "+
			"WHERE wu.user_id = $1 and ew.id IS NULL OR wu.word_id IS NULL and "+tab+" ORDER BY RANDOM() LIMIT $2", userId, countOfWords)
	if err != nil {
		logs.UpdateFile(time.Now().Format("15:04:23") + err.Error())
	}

	for rows.Next() {
		r := _struct.English{}
		err := rows.Scan(&r.Word, &r.Transcription, &r.Translate, &r.ID)
		if err != nil {
			logs.UpdateFile(time.Now().Format("15:04:23") + err.Error())
			continue
		}
		example := parseHtml(r.Word)
		if example[3] != "" {
			word = append(word, _struct.English{ID: r.ID, Word: r.Word, Transcription: r.Transcription, Translate: r.Translate,
				Example1: example[0], TranslateExample1: example[1], Example2: example[2], TranslateExample2: example[3]})
		} else if example[1] != "" {
			word = append(word, _struct.English{ID: r.ID, Word: r.Word, Transcription: r.Transcription, Translate: r.Translate,
				Example1: example[0], TranslateExample1: example[1]})
		} else {
			word = append(word, _struct.English{ID: r.ID, Word: r.Word, Transcription: r.Transcription, Translate: r.Translate})
		}
		id, _ := strconv.Atoi(r.ID)
		scanId = append(scanId, id)
	}

	for _, r := range results {
		fmt.Println(r)
	}

	addReferences(*db, userId, scanId, tab)

	defer db.Close()
	return word
}

func addReferences(db sql.DB, userId int, scanId []int, tab string) {
	for _, wordId := range scanId {
		_, err := db.Exec("INSERT INTO word_user(word_id, user_id) VALUES ($1, $2)", wordId, userId)
		if err != nil {
			logs.UpdateFile(time.Now().Format("15:04:23") + err.Error())
		}
	}
}

func Registration(login string, password string, name string) _struct.BoolAnswer {
	var regAns _struct.BoolAnswer
	db := ConnToDatabase()
	_, err := db.Exec("INSERT INTO users(login, password, name) VALUES ($1, $2, $3)",
		login, password, name)
	if err != nil {
		logs.UpdateFile(time.Now().Format("15:04:23") + err.Error())
		regAns.BoolAnswer = false
		return regAns
	} else {
		regAns.BoolAnswer = true
		return regAns
	}
}

func Login(login string, password string) _struct.BoolAnswer {
	var logAns _struct.BoolAnswer
	db := ConnToDatabase()
	rows, err := db.Query("SELECT login FROM users WHERE login=$1 AND password=$2",
		login, password)
	if err != nil {
		logs.UpdateFile(time.Now().Format("15:04:23") + err.Error())
	}
	if rows.Next() {
		logAns.BoolAnswer = true
		return logAns
	} else {
		logAns.BoolAnswer = false
		return logAns
	}
}

func ChangeTable(login string, password string) _struct.BoolAnswer {
	var tableType _struct.BoolAnswer
	db := ConnToDatabase()
	rows, err := db.Query("SELECT table_type FROM users WHERE login=$1 AND password=$2",
		login, password)
	if rows.Next() {
		err = rows.Scan(&tableType.BoolAnswer)
		if err != nil {
			logs.UpdateFile(time.Now().Format("15:04:23") + err.Error())
			tableType.BoolAnswer = false
			return tableType
		}
		if tableType.BoolAnswer {
			_, err = db.Exec("UPDATE users SET table_type = false WHERE login=$1 AND password=$2",
				login, password)
			if err != nil {
				logs.UpdateFile(time.Now().Format("15:04:23") + err.Error())
				tableType.BoolAnswer = true
				return tableType
			}
		} else {
			_, err = db.Exec("UPDATE users SET table_type = true WHERE login=$1 AND password=$2",
				login, password)
			if err != nil {
				logs.UpdateFile(time.Now().Format("15:04:23") + err.Error())
				tableType.BoolAnswer = true
				return tableType
			}
		}
	}
	return tableType
}

func ConnToDatabase() *sql.DB {
	db, errorDb := sql.Open("postgres", CONN)
	if errorDb != nil {
		panic(errorDb)
	}
	return db
}
