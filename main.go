package main

import (
	"fmt"
	"io"
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

		for _, user := range users {
			fmt.Fprintln(w, user)
		}

		return
	}

	if r.Method == "POST" && r.URL.Path == "/users" {
		body, err := io.ReadAll(r.Body)

		if err != nil {
			log.Println(err)
			return
		}

		name := string(body)

		if name == "" {
			http.Error(w, "name is missing", http.StatusBadRequest)
			return
		}

		newId := uuid.New().String()

		newUser := User{
			ID:   newId,
			Name: name,
		}

		users[newId] = newUser

		w.WriteHeader(http.StatusCreated)

		fmt.Fprintln(w, "User created")

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

			fmt.Fprintln(w, user.ID, user.Name)

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
