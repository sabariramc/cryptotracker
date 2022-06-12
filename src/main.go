package main

import (
	"log"
	"net/http"

	"thinklink/src/app"
)

func main() {
	s, err := app.GetDefaultApp()
	if err != nil {
		panic(err)
	}
	log.Fatal(http.ListenAndServe(s.GetPort(), s))
}
