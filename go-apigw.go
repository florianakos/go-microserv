package main

import (
	_ "bytes"
	_ "crypto/tls"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
	_ "time"
)

// global variable to hold the DATASTORE address
// (TODO: change to something more dynamic)
const DATASTORE_ADDR string = "http://localhost:7001"

type MyError struct {
	Error string `json:"error"`
}

// helper function to return a JSON response
func sendJSONError(w http.ResponseWriter, code int, message string) {
	response, _ := json.Marshal(&MyError{Error: message})
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

// helper function to return a JSON response
func sendJSONOkay(w http.ResponseWriter, code int, payload string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write([]byte(payload))
}

// handler for the static mainpage with HTML/JS to display the stored values...
func mainPageHandler(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	vars := mux.Vars(r)

	if vars["file"] == "style.css" {
		w.Header().Set("Content-Type", "text/css; charset=utf-8")
		dat, err := ioutil.ReadFile("./static/style.css")
		if err != nil {
			fmt.Println(err)
		}
		io.WriteString(w, string(dat))

	} else if vars["file"] == "index.js" {
		w.Header().Set("Content-Type", "text/javascript; charset=utf-8")
		dat, err := ioutil.ReadFile("./static/index.js")
		if err != nil {
			fmt.Println(err)
		}
		io.WriteString(w, string(dat))

	} else /*if vars["file"] == "index.html"*/ {
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		dat, err := ioutil.ReadFile("./static/index.html")
		if err != nil {
			fmt.Println(err)
		}
		io.WriteString(w, string(dat))
	}
}

// handler function that acts as the API GW sending the HTTP GET to the Datastore container
func getSensorData(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	response, err := http.Get(DATASTORE_ADDR + r.RequestURI)

	if err != nil {
		fmt.Printf("ERROR: %s\n", err.Error())
		sendJSONError(w, http.StatusInternalServerError, "Error connecting to datastore!")
	} else {
		contents, err := ioutil.ReadAll(response.Body)
		if err != nil {
			fmt.Printf("ERROR: %s\n", err.Error())
			sendJSONError(w, http.StatusInternalServerError, "Error processing the datastore response!")
			response.Body.Close()
		}
		sendJSONOkay(w, http.StatusOK, string(contents))
	}
}

func main() {
	// check if the program was called with proper argument list
	if len(os.Args[1:]) != 1 {
		fmt.Println("Usage: go run %s [PORT]\n\n", os.Args[0])
		os.Exit(1)
	}

	// exit if the argument is not valid integer...
	_, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Printf("Starting server on port: %s\n\n", os.Args[1])

	// configure and start MUX router
	r := mux.NewRouter()
	r.HandleFunc("/{file}", mainPageHandler).Methods("GET")
	r.HandleFunc("/api/sensors", getSensorData).Methods("GET")
	r.HandleFunc("/api/sensors/{sensorid}", getSensorData).Methods("GET")

	if err := http.ListenAndServe(":"+os.Args[1], r); err != nil {
		log.Fatal(err)
	}
}
