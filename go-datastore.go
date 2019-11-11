package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"net/http"
	"os"
	"strconv"
)

// Record is a struct that represents rows in the database table
type Record struct {
	Id           int     `json:"id"`
	Sensorid     string  `json:"sensorid"`
	Batterylevel int     `json:"batterylevel"`
	Humidity     int     `json:"humidity"`
	Light        int     `json:"light"`
	Motion       int     `json:"motion"`
	Temperature  float64 `json:"temperature"`
	Timestamp    string  `json:"timestamp"`
}

// MyErris is a struct representing custom error type
type MyError struct {
	Error string `json:"error"`
}

// global variable to hold the database reference
var db *sql.DB

// method on the Record struct to easily print it as string
func (r *Record) toString() string {
	return fmt.Sprintf("%v, %v, %v, %v, %v, %v, %v, %v\n", r.Id,
		r.Sensorid, r.Batterylevel, r.Humidity, r.Light,
		r.Motion, r.Temperature, r.Timestamp)
}

// helper function to return a JSON response
func sendJSONError(w http.ResponseWriter, code int, message string) {
	response, _ := json.Marshal(&MyError{Error: message})
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

// helper function to return a JSON response
func sendJSONOkay(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

// GetAllSensorData is a helper function to get Sensor data from all sensors...
func GetAllSensorData(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	// temp variable to hold the querystring
	var queryStr string

	// check w
	limit := r.URL.Query().Get("number")
	if limit == "" {

		queryStr = "SELECT * FROM datastore ORDER BY 1 DESC;"

	} else {

		_, err := strconv.Atoi(limit)
		if err != nil {
			fmt.Println(err.Error())
			sendJSONError(w, http.StatusBadRequest, "Invalid URL parameter (needs valid Integer)!")
			return
		}
		queryStr = fmt.Sprintf("SELECT * FROM datastore ORDER BY 1 DESC LIMIT %v;", limit)

	}

	var rows *sql.Rows
	rows, err := db.Query(queryStr)
	if err != nil {
		fmt.Println(err.Error())
		sendJSONError(w, http.StatusInternalServerError, "Internal Error")
		return
	}

	var dbRecords []Record
	for rows.Next() {
		var record Record
		err := rows.Scan(&record.Id, &record.Sensorid,
			&record.Batterylevel, &record.Humidity,
			&record.Light, &record.Motion,
			&record.Temperature, &record.Timestamp)
		if err != nil {
			fmt.Println(err.Error())
			sendJSONError(w, http.StatusInternalServerError, "Internal Error")
			return
		}
		dbRecords = append(dbRecords, record)
	}

	r.Close = true
	sendJSONOkay(w, http.StatusOK, dbRecords)
}

// GetSensor is a helper function to get data from single sensor
func GetSensor(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	// get sensor ID from MUX router
	sensorID := mux.Vars(r)["sensorid"]

	// temp string to hold the queryString
	var queryStr string

	limit := r.URL.Query().Get("number")
	if limit == "" {

		queryStr = fmt.Sprintf("SELECT * FROM datastore WHERE sensorid = '%v' ORDER BY 1 DESC;", sensorID)

	} else {

		_, err := strconv.Atoi(limit)
		if err != nil {
			fmt.Println(err.Error())
			sendJSONError(w, http.StatusBadRequest, "Invalid URL parameter (needs valid Integer)!")
			return
		}
		queryStr = fmt.Sprintf("SELECT * FROM datastore WHERE sensorid = '%v' ORDER BY 1 DESC LIMIT %v;", sensorID, limit)

	}

	var rows *sql.Rows
	rows, err := db.Query(queryStr)
	if err != nil {
		fmt.Println(err.Error())
		sendJSONError(w, http.StatusInternalServerError, "Internal Error")
		return
	}

	var dbRecords []Record
	for rows.Next() {
		var record Record
		err := rows.Scan(&record.Id, &record.Sensorid,
			&record.Batterylevel, &record.Humidity,
			&record.Light, &record.Motion,
			&record.Temperature, &record.Timestamp)
		if err != nil {
			fmt.Println(err.Error())
			sendJSONError(w, http.StatusInternalServerError, "Internal Error")
			return
		}
		dbRecords = append(dbRecords, record)
	}

	if len(dbRecords) == 0 {
		sendJSONError(w, http.StatusNotFound, "Sensor not found in the database!")
	} else {
		sendJSONOkay(w, http.StatusOK, dbRecords)
	}
}

// SaveSensorData helps to save data to the database...
func SaveSensorData(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	// temp variable to store the parsed JSON fields
	var t Record

	// unmarshal the HTTP request body which is JSON and contains the data to save to DB
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&t)
	if err != nil {
		sendJSONError(w, http.StatusInternalServerError, "Error parsing the JSON data!")
		return
	}

	stmt, err := db.Prepare("INSERT INTO datastore(sensorid, batterylevel, humidity, light, motion, temperature, timestamp) values(?, ?, ?, ?, ?, ?, ?);")
	if err != nil {
		sendJSONError(w, http.StatusInternalServerError, "Error preparing the SQL statement!")
		return
	}

	_, err = stmt.Exec(t.Sensorid, t.Batterylevel, t.Humidity,
		t.Light, t.Motion, t.Temperature, t.Timestamp)
	if err != nil {
		sendJSONError(w, http.StatusInternalServerError, "Error executing the SQL statement!")
		return
	}

	sendJSONOkay(w, http.StatusOK, struct{ Message string }{Message: "Insert successful!"})

}

func main() {
	// check if program was called correctly, with PORT in args
	if len(os.Args[1:]) != 1 {
		fmt.Println("Usage: go run %s [PORT]\n\n", os.Args[0])
		os.Exit(1)
	}
	// check if the supplied argument is valid INT
	_, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// open the SQLITE3 database
	db, _ = sql.Open("sqlite3", "./database.db")
	defer db.Close()

	// configure and start the REST API server
	PORT := os.Args[1]
	fmt.Printf("Starting server on port [%s]!\n", PORT)

	r := mux.NewRouter()
	r.HandleFunc("/api/sensors", GetAllSensorData).Methods("GET")
	r.HandleFunc("/api/sensors/{sensorid}", GetSensor).Methods("GET")
	r.HandleFunc("/api/sensors", SaveSensorData).Methods("POST")

	if err := http.ListenAndServe(":"+PORT, r); err != nil {
		log.Fatal(err)
	}
}
