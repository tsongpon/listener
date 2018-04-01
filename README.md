# Listener
 Listen for facebook user's activities and save to database

 **subscribed user's event** : about, birthday, email, first_name, gender, last_name, local,
 pic_big_https, pic_big_with_logo, pic_https, pic_small_https, pic_small_with_logo,
 pic_square_https, pic_square_with_logo, pic_with_logo, quote, religion, website
 
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