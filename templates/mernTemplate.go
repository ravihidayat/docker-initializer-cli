package templates

var MernDockerCompose string = `
version: '3.1'

services:
    node:
      build: 
        context: .
        dockerfile: Dockerfile
      ports:
          - '3000:3000'
      volumes: 
        - .:/app/mern-docker-app
    
    mongo:
      image: {.param}
      ports: 
        - '27017:27017'
      volumes:
          - ~/mongo-db-docker-data:/data/db
		  
`

var MernDockerfile string = `

FROM node:latest

WORKDIR /app/mern-docker-app

EXPOSE 3000

CMD [ "yarn","start-dev" ]
`
