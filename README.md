# Fast Quiz: Golang API & Cobra CLI

## Server
![](./images/server1.png)
### With Docker
````
cd server
docker build -t server .
docker run -it -p 9444:9444 server
````
### With docker-compose
````
docker-compose build
docker network create ytimoumi
docker-compose up
````
### Without Docker
````
cd server
go mod download
go run main.go
````
### Host & port
````
Started server on : 
http://localhost:9444/

Test with Curl:
curl --location --request GET 'localhost:9444/v1/questions/1'

Result:
[{"question":"What is the capital city of Tunisia ?","answers":["Sfax","Tunis","Kairouan"]},{"question":"How many countries are there in the region of Europe ?","answers":["42","43","44"]},{"question":"How many elements are in the periodic table ?","answers":["118","117","119"]}]
````

## Client
![](./images/cli1.png)

![](./images/cli2.png)
````
cd client
go mod download
go run main.go answer --id=1
go run main.go answer --id=2
````
**Note**: There are 2 question with ID 1 and 2 for testing.