package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

type responseObject struct {
	UserId  int		`json:"userId"`
	ID 		int		`json:"id"`
	Title	string	`json:"title"`
	Body	string	`json:"body"`
}

func main() {
	response, err := http.Get("https://jsonplaceholder.typicode.com/posts")

	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}
	
	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}
	var responses []responseObject
	json.Unmarshal(responseData, &responses)
	for k := range responses {
		fmt.Printf( "Title: '%s' id: '%v'\n", responses[k].Title, responses[k].ID );
	}
}
