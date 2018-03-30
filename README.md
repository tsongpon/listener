# Listener    
 Listen for facebook user's activities and save to database
 
 **subscribed user's event** : about, birthday, email, first_name, gender, last_name, local, 
 pic_big_https, pic_big_with_logo, pic_https, pic_small_https, pic_small_with_logo, 
 pic_square_https, pic_square_with_logo, pic_with_logo, quote, religion, website
    
## Tech stack 
- Go for programming language    
- Mongodb for storage  
- Docker for deployment (running on AWS Elastic Beanstalk)  
  
## Run
required environment variable

    export REDPLANET_DB_HOST=YOUR_DB_HOST_IP
    export TOKEN=YOUR_FACEBOOK_WEBHOOK_TOKEN

run service 

    docker-compose build
    docker-compose up

## Infrastructure

![enter image description here](https://github.com/tsongpon/listener/blob/master/diagram/infra.png?raw=true)

## System Architecture

![enter image description here](https://github.com/tsongpon/listener/blob/master/diagram/architecture.png?raw=true)
