# Very Simple Microservices architecture in Go

Simple microservices project in GO for M7024E @ LTU

![Basic architecture](https://github.com/florianakos/go-microserv/blob/master/static/Screenshot%20from%202019-01-05%2023-15-12.png "Basic Architecture")

The microservices are:
* API Gateway
* Collector
* Datastore

For simplicity the API Gateway also serves the static main-page with the Javascript embedded to call the API GW endpoints.

### Port 7000 is for the API GW which can handle
```
* /index.html
* /api/sensors [GET/POST]
* /api/sensors/{sensor-id} [GET]
```
An optional ?number=X URL parameter can be passsed to limit the number of datapoints returned.

### Port 7001 is for the Datastore which for now stores the data in a simple SQLITE3 local db
```
* /api/sensors
* /api/sensors/{sensor-id}
```