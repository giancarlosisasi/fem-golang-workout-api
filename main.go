package main

import (
	"fm-api-project/internal/app"
	"fmt"
	"net/http"
	"time"
)

func main() {
	app, err := app.NewApplication()
	if err != nil {
		// it's not recommended to use panic, there are other better ways
		panic(err)
	}

	app.Logger.Println("We are running our app")

	http.HandleFunc("/health", HealthCheck)
	server := &http.Server{
		Addr:         ":8080",
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	err = server.ListenAndServe()
	if err != nil {
		app.Logger.Fatal(err)
	}

}

func HealthCheck(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Status is available")
}
