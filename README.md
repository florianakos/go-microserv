# Very Simple Microservices architecture in Go

Simple microservices project in GO for M7024E @ LTU

![Basic architecture](https://github.com/florianakos/go-microserv/blob/master/static/Screenshot%20from%202019-01-05%2023-15-12.png "Basic Architecture")

The microservices are:
* API Gateway
* Collector
* Datastore

For simplicity the API Gateway also serves the static main-page with the Javascript embedded to call the API GW endpoints.

### The API GW
 
It uses port 7000 and handles the below endpoints
```
* /index.html                [GET]
* /api/sensors               [GET]
* /api/sensors/{sensor-id}   [GET]
```
An optional ?number=X URL parameter can be passsed to limit the number of datapoints returned.

### The DataStore

It uses port 7001 and handles the below endpoints
```
* /api/sensors               [GET/POST]
* /api/sensors/{sensor-id}   [GET]
```
The data for now is stored in a simple SQLite3 lcoal database.

### The Collector

It does not need a port because it is simply implementing a functionality as described on the architecture.
It has two hard-coded timers defined. One defines how often to refresh the SSiO auth token and another defines how often to poll and store new Data from SSiO.


### Plans for improvement:
* migrate away from SQLite3 to PostgreSQL
* fully implement it via docker containers
* deploy in a docker swarm