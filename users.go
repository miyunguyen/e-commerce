package main

type User struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type CreateUserPayload struct {
	Name string `json:"name"`
}

var users = map[string]User{
	"b545a1fd-b20b-40c4-b378-6ea56c75a6c3": {
		ID:   "b545a1fd-b20b-40c4-b378-6ea56c75a6c3",
		Name: "Ahoy",
	},
	"17dd3e0c-5908-4d6f-873f-aee7fc6b6633": {
		ID:   "17dd3e0c-5908-4d6f-873f-aee7fc6b6633",
		Name: "Aza",
	},
}
