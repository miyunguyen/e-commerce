package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/google/uuid"
)

type server struct {
	addr string
}

func (s *server) getUserList(w http.ResponseWriter, r *http.Request) {
	userList := []User{}

	for _, user := range users {
		userList = append(userList, user)
	}

	writeJSON(w, http.StatusOK, userList)
}

func (s *server) addUser(w http.ResponseWriter, r *http.Request) {
	req := CreateUserPayload{}
	err := json.NewDecoder(r.Body).Decode(&req)

	if err != nil {
		http.Error(w, "error reading json", http.StatusBadRequest)
		return
	}

	if req.Name == "" {
		http.Error(w, "name is missing", http.StatusBadRequest)
		return
	}

	newId := uuid.New().String()

	newUser := User{
		ID:   newId,
		Name: req.Name,
	}

	users[newId] = newUser

	writeJSON(w, http.StatusCreated, newUser)
}

func (s *server) getUser(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")

	user, err := users[id]

	if !err {
		fmt.Fprintln(w, "user not found", http.StatusNotFound)
		return
	}

	writeJSON(w, http.StatusOK, user)
}

func main() {
	svr := &server{
		addr: ":8080",
	}

	mux := http.NewServeMux()

	mux.HandleFunc("GET /users", svr.getUserList)
	mux.HandleFunc("POST /users", svr.addUser)
	mux.HandleFunc("GET /users/{id}", svr.getUser)

	log.Println("Server is running")

	http.ListenAndServe(svr.addr, mux)
}
