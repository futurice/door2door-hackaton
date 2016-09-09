package api

import (
	"net/http"
	"encoding/json"
)

type Greetings struct {
	Message string `json:"message"`
	From string `json:"from"`
}

func Greeter(res http.ResponseWriter, req *http.Request) {

	json, _ := json.Marshal(Greetings{
		Message: "Hello",
		From: "World",
	})
	res.Header().Set("Content-Type", "application/json; charset=UTF-8")
	res.Write(json)
}