package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func main() {
	port := "8080"

	if p := os.Getenv("APP_PORT"); p != "" {
		port = p
	}

	router := mux.NewRouter()
	router.HandleFunc("/", HelloServer)
	router.HandleFunc("/whoami", WhoAmI)

	log.Println("Application started at port " + port)

	h1 := handlers.CombinedLoggingHandler(os.Stdout, router)
	h2 := handlers.CompressHandler(h1)

	log.Fatal(http.ListenAndServe("0.0.0.0:"+port, h2))
}

// HelloServer return hello world text
func HelloServer(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World!")
}

// WhoAmI return my name is some_name
func WhoAmI(w http.ResponseWriter, r *http.Request) {
	name := "prakash"

	if n := os.Getenv("NAME"); n != "" {
		name = n
	}

	fmt.Fprintf(w, "my name is "+name)
}
