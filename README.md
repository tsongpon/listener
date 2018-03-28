# Listener    
 Listen for facebook user's activities and save to database    
    
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
