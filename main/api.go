package main

import (
	"../serverCommand"
	"../struct"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

const (
	GET_WORDS    = "\bgetWords"
	REGISTRATION = "registration"
	CONN         = "user=postgres dbname=postgres sslmode=disable"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/words", getWords).Methods("POST")
	r.HandleFunc("/registration", regUser).Methods("POST")
	r.HandleFunc("/login", logUser).Methods("POST")
	r.HandleFunc("/change", changeTable).Methods("POST")
	log.Fatal(http.ListenAndServe(":8080", r))
}

func regUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var err error
	var reg _struct.RegistrationRequest
	var hasRegistered _struct.BoolAnswer
	err = json.NewDecoder(r.Body).Decode(&reg)
	if err != nil {
		fmt.Println(err)
	}
	hasRegistered = serverCommand.Registration(reg.Login, reg.Password, reg.Name)
	err = json.NewEncoder(w).Encode(hasRegistered)
	fmt.Println(hasRegistered)
	if err != nil {
		fmt.Println(err)
	}
}

func logUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var err error
	var login _struct.LoginRequest
	var hasLogined _struct.BoolAnswer
	err = json.NewDecoder(r.Body).Decode(&login)
	if err != nil {
		fmt.Println(err)
	}
	hasLogined = serverCommand.Login(login.Login, login.Password)
	err = json.NewEncoder(w).Encode(hasLogined)
	fmt.Println(hasLogined)
	if err != nil {
		fmt.Println(err)
	}
}

func getWords(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var words []_struct.English
	var word _struct.Words
	json.NewDecoder(r.Body).Decode(&word)
	words = serverCommand.GetWords(word.CountOfWords, word.Login)
	fmt.Println(words)
	err := json.NewEncoder(w).Encode(words)
	if err != nil {
		fmt.Println(err)
	}
}

func changeTable(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var err error
	var login _struct.LoginRequest
	var res _struct.BoolAnswer
	err = json.NewDecoder(r.Body).Decode(&login)
	if err != nil {
		fmt.Println(res)
	}
	res = serverCommand.ChangeTable(login.Login, login.Password)
	fmt.Println(res)
	err = json.NewEncoder(w).Encode(res)
	if err != nil {
		fmt.Println(err)
	}
}
