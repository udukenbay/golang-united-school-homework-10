package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/gorilla/mux"
)

/**
Please note Start functions is a placeholder for you to start your own solution.
Feel free to drop gorilla.mux if you want and use any other solution available.

main function reads host/port from env just for an example, flavor it following your taste
*/

// Start /** Starts the web server listener on given host and port.
func Start(host string, port int) {
	router := mux.NewRouter()

	log.Println(fmt.Printf("Starting API server on %s:%d\n", host, port))
	if err := http.ListenAndServe(fmt.Sprintf("%s:%d", host, port), router); err != nil {
		log.Fatal(err)
	}

	router.HandleFunc("/bad", handleBad).Methods(http.MethodGet)
	router.HandleFunc("/name/{user}", handleName).Methods(http.MethodGet)
	router.HandleFunc("/data", handlePost).Methods(http.MethodPost)

	log.Fatalln(http.ListenAndServe(":8081", router))
}

func handlePost(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		fmt.Fprintf(w, "invalid_http_method")
		return
	}

	r.ParseForm()
	fmt.Fprintf(w, "I got message:\n"+r.Form.Get("name"))
}

func handleBad(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusBadGateway)
}

func handleName(w http.ResponseWriter, r *http.Request) {
	name := "mister X"
	if p, ok := mux.Vars(r)["user"]; ok {
		name = p
	}
	fmt.Fprintf(w, "Hello, %s!", name)
}

//main /** starts program, gets HOST:PORT param and calls Start func.
func main() {
	host := os.Getenv("HOST")
	port, err := strconv.Atoi(os.Getenv("PORT"))
	if err != nil {
		port = 8081
	}
	Start(host, port)
}
