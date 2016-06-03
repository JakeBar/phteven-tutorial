package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/phteven", phtevenHandler)
	http.ListenAndServe(":8080", nil)
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "This is the home page!")
}

func phtevenHandler(w http.ResponseWriter, r *http.Request) {
	text := r.URL.Query().Get("text")
	body, _ := getPhtevenResponseBody(text)
	fmt.Fprintf(w, "%s", body)
}

func getPhtevenResponseBody(text string) ([]byte, error) {
	//declare variables
	var url string = "http://api.phteven.io/translate"
	dog := Dog{"Phteven", 3, text}

	requestBody := bytes.NewBufferString("text=%s" + dog.say())

	response, err := http.Post(url, "application/x-www-form-urlencoded", requestBody)

	if err != nil {
		fmt.Printf("Error sending request: %v\n", err)
		return []byte(""), err
	}

	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Printf("Error sending body: %v\n", err)
		return []byte(""), err
	}

	return body, nil
}
