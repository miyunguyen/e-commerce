package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/google/uuid"
)

type server struct {
	addr string
}

func (s *server) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	if r.Method == "GET" && r.URL.Path == "/users" {
		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-type", "application/json")

		userList := []User{}

		for _, user := range users {
			userList = append(userList, user)
		}
		json.NewEncoder(w).Encode(userList)

		return
	}

	if r.Method == "POST" && r.URL.Path == "/users" {

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

		w.WriteHeader(http.StatusCreated)
		w.Header().Set("Content-type", "application/json")

		json.NewEncoder(w).Encode(newUser)

		return
	}

	if r.Method == "GET" {
		parts := strings.Split(r.URL.Path, "/")

		if len(parts) == 3 && parts[1] == "users" {
			id := parts[2]

			user, err := users[id]

			if !err {
				fmt.Fprintln(w, "user not found", http.StatusNotFound)
				return
			}

			w.Header().Set("Content-type", "application/json")

			json.NewEncoder(w).Encode(user)

			return
		}
	}

	http.Error(w, "Not found", http.StatusNotFound)
}

func main() {
	svr := &server{
		addr: ":8080",
	}

	log.Println("Server is running")

	http.ListenAndServe(svr.addr, svr)
}
