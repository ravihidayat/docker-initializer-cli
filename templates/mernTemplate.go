package templates

var MernDockerCompose string = `
version: '3.7'

services:
    node:
      build: 
        context: .
        dockerfile: Dockerfile
      ports:
          - '3000:3000'
      volumes: 
        - .:/app
    
    mongo:
      image: mongo:{{.dbTag}}
      env_file:
        - .env
      ports: 
        - '27017:27017'
      volumes:
          - ~/mongo-db-docker-data:/data/db
		  
`

var MernDockerfile string = `

FROM node:{{.nodeTag}}

WORKDIR /app

EXPOSE 3000

CMD [ "npm","start" ]
`
