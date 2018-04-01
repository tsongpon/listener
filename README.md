# Listener
[![Go Report Card](https://goreportcard.com/badge/github.com/tsongpon/listener)](https://goreportcard.com/report/github.com/tsongpon/listener)

 Listen for facebook user's activities and save to database

 **subscribed user's event** : Only basic information, first name, last name, profile picture, gender and age range.
 
## Project dependencies
- Go - programming language
- Dep - package management
- Mongodb - data storage
- Docker - deployment and integration test

## Test (integration)
mongodb docker container will be created and tests will run against it, after finish test container will be removed

    go test -v

## Run
required environment variable

    export REDPLANET_DB_HOST=YOUR_DB_HOST_IP
    export TOKEN=YOUR_FACEBOOK_WEBHOOK_TOKEN
    
run service (native)

    dep ensure
    go build -o bin/listener .
    ./bin/listener

run service (docker-compose)

    docker-compose build
    docker-compose up

service will be running on port `5000`

## Infrastructure

![enter image description here](https://github.com/tsongpon/listener/blob/master/diagram/infra.png?raw=true)

## System Architecture

![enter image description here](https://github.com/tsongpon/listener/blob/master/diagram/architecture.png?raw=true)

## API(s)
**get user activities**

    GET http://localhost:5000/useractivities
    
query parameter supported:  
	- `userid` : filter by userId   
	- `field` : filter by update field    
	- `size` : limit response size, default value is 5  
	- `start` : specify start offset of response (response order by time)   
