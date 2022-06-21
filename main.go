package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type User struct {
	ID     int
	Nama   string
	Divisi string
}

var user = map[int]User{}

// GET, Get All User
func getAllUser(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(user)
}

//GET, Ger by Id
func getUserById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	queryId, err := strconv.Atoi(vars["id"])
	if err != nil {
		panic(err.Error())
	}
	for _, item := range user {
		if item.ID == queryId {
			json.NewEncoder(w).Encode(item)
		}
	}
}

//POST, Create
func createUser(w http.ResponseWriter, r *http.Request) {
	var NewUser User
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&NewUser)
	if err != nil {
		panic(err.Error())
	}
	// log.Println(NewUser) // for debug
	// log.Println(user[1].ID)
	user[NewUser.ID] = NewUser
	log.Println(user)
	json.NewEncoder(w).Encode(user)
}

//PUT, Update
func updateUserById(w http.ResponseWriter, r *http.Request) {
	var NewUser User
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&NewUser)
	if err != nil {
		panic(err.Error())
	}

	vars := mux.Vars(r)
	queryId, err := strconv.Atoi(vars["id"])
	if err != nil {
		panic(err.Error())
	}
	NewUser.ID = queryId
	user[queryId] = NewUser
	json.NewEncoder(w).Encode(user)
}

func deleteUserById(w http.ResponseWriter, r *http.Request) {
	// var NewUser User
	vars := mux.Vars(r)
	queryId, err := strconv.Atoi(vars["id"])
	if err != nil {
		panic(err.Error())
	}
	// for _, item := range user {
	// 	if item.ID == queryId {
	// 		json.NewEncoder(w).Encode(item)
	// 	}
	// }

	delete(user, queryId)
	json.NewEncoder(w).Encode(user)
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/user", getAllUser).Methods("GET")
	r.HandleFunc("/user/{id}", getUserById).Methods("GET")
	r.HandleFunc("/user", createUser).Methods("POST")
	r.HandleFunc("/user/{id}", updateUserById).Methods("PUT")
	r.HandleFunc("/user/{id}", deleteUserById).Methods("DELETE")
	http.Handle("/", r)
	http.ListenAndServe(":8080", r)
}
