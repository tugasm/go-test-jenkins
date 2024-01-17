package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, Dockerized Go App!")
	})

	http.ListenAndServe(":8080", nil)
}

// https://github.com/tugasm/go-test-jenkins
// Sample commands:
// Install the latest Weekly version: brew install jenkins
// Start the Jenkins service: brew services start jenkins --> start
// Restart the Jenkins service: brew services restart jenkins --> restart
// Update the Jenkins version: brew upgrade jenkins
