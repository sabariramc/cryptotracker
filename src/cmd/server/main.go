package main

import (
	"log"
	"net/http"

	"cryptotracker/src/app"
)

func main() {
	s, err := app.GetDefaultApp()
	if err != nil {
		panic(err)
	}
	// var wg sync.WaitGroup
	// wg.Add(1)
	// ctx, cancel := context.WithCancel(context.Background())
	// defer wg.Wait()
	// defer cancel()
	// go func() {
	// 	defer wg.Done()
	// 	// s.StartJob(ctx)
	// }()
	log.Fatal(http.ListenAndServe(s.GetPort(), s))
}
