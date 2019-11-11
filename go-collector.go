package main

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"
)

// global constant to store the SSIO auth URL
const SSIO_URL string = "http://130.240.134.128:9000/oauth2/token"

// global constant to store the DATASTORE remote address
const DATASTORE_ADDR string = "http://localhost:7001"

// global constant to control collector time interval
const collector_interval time.Duration = 5

// global variable to control sleep time
const token_refresh_interval time.Duration = 60

// global accesstoken for the SSiO platform
var access_token string = "NULL"

// Struct to represent entries in the database table
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

// struct to parse SSIO auth request JSON
type SSIOToken struct {
	AccessToken  string `json:"access_token"`
	TokenType    string `json:"token_type"`
	ExpiresIn    int    `json:"expires_in"`
	RefreshToken string `json:"refresh_token"`
}

// Struct to parse SSIO JSON data response
type SSIOResponse struct {
	ContextResponses []struct {
		ContextElement struct {
			Type       string `json:"type"`
			IsPattern  string `json:"isPattern"`
			ID         string `json:"id"`
			Attributes []struct {
				Name      string `json:"name"`
				Type      string `json:"type"`
				Value     string `json:"value"`
				Metadatas []struct {
					Name  string `json:"name"`
					Type  string `json:"type"`
					Value string `json:"value"`
				} `json:"metadatas,omitempty"`
			} `json:"attributes"`
		} `json:"contextElement"`
		StatusCode struct {
			Code         string `json:"code"`
			ReasonPhrase string `json:"reasonPhrase"`
		} `json:"statusCode"`
	} `json:"contextResponses"`
}

// refresh access_token indefinitely, every 1 minute
func refreshAccessToken() {

	for {
		// prepare request Headers, Body and Transport
		var payload = []byte(`grant_type=password&username=user01ssr%40ssr.se&password=password&client_id=8f2e5c99050b48348a5badfe68b55a67&client_secret=48f18ebe9120476a8e852f157dcf9ff5&undefined=`)
		req, err := http.NewRequest("POST", SSIO_URL, bytes.NewBuffer(payload))
		req.Header.Set("Authorization", "Basic OGYyZTVjOTkwNTBiNDgzNDhhNWJhZGZlNjhiNTVhNjc6NDhmMThlYmU5MTIwNDc2YThlODUyZjE1N2RjZjlmZjU=")
		req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
		req.Header.Set("cache-control", "no-cache")
		req.Header.Set("Postman-Token", "71d23475-d813-4ecb-9352-1abc0637adac")
		tr := &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		}

		// send the request and check if there were any errors
		client := &http.Client{Transport: tr}
		resp, err := client.Do(req)
		if err != nil {
			fmt.Println(err)
			fmt.Println("UNMARSHAL SSIO")
			continue
		}

		// read the response body and then close it to free up memory
		body, _ := ioutil.ReadAll(resp.Body)
		defer resp.Body.Close()

		// un-marshal the JSON and store the fields in a struct to get the token...
		var newCredentials SSIOToken
		err = json.Unmarshal(body, &newCredentials)
		if err != nil {
			fmt.Println(err.Error())
			fmt.Println("UNMARSHAL SSIO")
			continue
		}

		// update the global variable
		access_token = newCredentials.AccessToken

		// sleep for a minute before starting again...
		time.Sleep(token_refresh_interval * time.Second)
		// fmt.Printf("TickTock from the token_refresh goroutine (token=%s)\n", access_token)
	}
}

// contact the SSIO GW for new data from sensors
func querySSiO(sensorID string) Record {

	// prepare the request content and HTTP headers
	ssioURL := "https://130.240.134.128:3000/v1/queryContext/"
	var payload = []byte(`{"entities": [{"id": "` +
		sensorID +
		`", "type": "LORA_Sensor","isPattern": false}]}`)

	req, err := http.NewRequest("POST", ssioURL, bytes.NewBuffer(payload))
	req.Header.Set("X-Auth-Token", access_token)
	req.Header.Set("Content-Type", "application/json")
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}

	// send the request to SSIO and check if there was any error
	client := &http.Client{Transport: tr}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return Record{}
	}

	// read the content of the response Body
	body, _ := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()

	// temp variable for the JSON un-marshaling
	var ssiodata SSIOResponse
	err = json.Unmarshal(body, &ssiodata)
	if err != nil {
		fmt.Println(err.Error())
		return Record{}
	}
	if &ssiodata.ContextResponses == nil {
		// return empty struct
		return Record{}
	}

	// temp variable to hold values from JSON response
	var sensorData Record
	sensorData.Sensorid = sensorID

	// iterate over the various attributes in the JSON response
	for _, val := range ssiodata.ContextResponses[0].ContextElement.Attributes {

		if val.Name == "Battery level" {
			v, _ := strconv.Atoi(val.Value)
			sensorData.Batterylevel = v
		} else if val.Name == "Humidity" {
			v, _ := strconv.Atoi(val.Value)
			sensorData.Humidity = v
		} else if val.Name == "Light" {
			v, _ := strconv.Atoi(val.Value)
			sensorData.Light = v
		} else if val.Name == "Motion" {
			v, _ := strconv.Atoi(val.Value)
			sensorData.Motion = v
		} else if val.Name == "Temperature" {
			v, _ := strconv.ParseFloat(val.Value, 64)
			sensorData.Temperature = v
		} else if val.Name == "Timestamp" {
			sensorData.Timestamp = val.Value
		}
	}
	// return the struct that was filled up with useful data points
	return sensorData
}

func saveToDatastore(dataToStore Record) {
	// marshal the struct into a JSON byte array!
	jsonStr, err := json.Marshal(dataToStore)
	if err != nil {
		fmt.Println(err)
	}

	// prepare an HTTP POST request with Header set to "app/json"
	req, err := http.NewRequest("POST", DATASTORE_ADDR+"/api/sensors", bytes.NewBuffer(jsonStr))
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}

	// capture the response and see if there were any errors!
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}

	// check the returned STATUS CODE of the HTTP POST.
	if resp.StatusCode == 200 {
		fmt.Println("HTTP POST to DATASTORE successful!")
	} else {
		body, _ := ioutil.ReadAll(resp.Body)
		fmt.Printf("HTTP POST to DATASTORE failed!\n-StatusCode: \t%d\n-Message:\t%v\n", resp.StatusCode, string(body))
	}
}

func main() {

	fmt.Println("Starting background token update method ...")
	go refreshAccessToken()

	time.Sleep(5 * time.Second)

	fmt.Println("Starting data collection routine ...")
	for {
		saveToDatastore(querySSiO("a81758fffe031a79"))
		saveToDatastore(querySSiO("a81758fffe031a81"))
		saveToDatastore(querySSiO("a81758fffe031a83"))
		saveToDatastore(querySSiO("a81758fffe031a82"))

		time.Sleep(collector_interval * time.Second)
	}
}
