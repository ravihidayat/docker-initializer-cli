package templates

var PernDockerCompose string = `
version: '3.7'
services:
  api:
   container_name: {{.projectName}}-api
   build:
    context: .
    dockerfile: Dockerfile.api
   volumes:
    - /api/node_modules
    - .:/api
   ports: 
     - "3000:3000"
   networks:  
    - {{.projectName}}-app
    - {{.projectName}}-backend
   restart: unless-stopped

  app:          
   container_name: {{.projectName}}-app
   build:
    context: .
	dockerfile: Dockerfile.app
   volumes:
    - /app/node_modules
	- .:/app
   ports:
    - "8000:3000"
   networks:
    - app
   restart: unless-stopped
   				
  postgres:
   image: postgres:{{.dbTag}}
   container_name: {{.projectName}}-db
   ports:
    - "5432:5432"
   env_file:
    - .env
   volumes:
    - postgres:/var/lib/postgresql/data
   networks:	
    - {{.projectName}}-backend
   restart: unless-stopped		 

volumes:
 postgres:
  name: {{.projectName}}-db  	
networks:
 app:
 backend:
   driver: bridge	
`
var ApiDockerfile string = `
FROM node:{{.nodeTag}}
WORKDIR /api

COPY package.json /api/package.json
COPY package-lock.json /api/package-lock.json

RUN npm install

CMD ["npm", "run", "start"]
`

var AppDockerfile string = `
FROM node:{{.nodeTag}}
EXPOSE 8000
WORKDIR /app

COPY package.json /app/package.json
COPY package-lock.json /app/package-lock.json

RUN npm install

CMD ["npm", "run", "start"]
`
