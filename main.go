package main

import (
	"log"
	"net/http"
)

type server struct {
	addr string
}

func (s *server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello from server!"))
}

func main() {
	svr := &server{
		addr: ":8080",
	}

	log.Println("Server is running")

	http.ListenAndServe(svr.addr, svr)
}
