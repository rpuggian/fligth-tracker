# fligth-tracker
fligth tracker tracks a person path of flights 

This Service Expose an API to sort Filghts in order to track his path. 

## How to run it ? 

You can run this project via command line using: 

```
go run main.go
```
The default port is setted to `:8080`

## How to use the api?

> POST  /track 

This endpoint is responsable to sort a list of paths of flights.

Sample request: 
``` 
curl --location --request POST 'localhost:8080/track' \
--header 'Content-Type: application/json' \
--data-raw '[["ATL", "EWR"], ["SFO", "ATL"]]' 
```