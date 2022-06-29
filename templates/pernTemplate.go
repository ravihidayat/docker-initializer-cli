package templates

var Pern string = `
version: '3.7'
services:
    api:
		container_name: {{.projectName}}-api
		build:
		context: ./app
		dockerfile: Dockerfile.api
		volumes:
		 - /api/node_modules
		 - ./api:/api
		ports:
		 - "3000:3000"
		networks:
		 - {{.projectName}}-app
		 - {{.projectName}}-backend
		restart: unless-stopped
	app:
		container_name: {{.projectName}}-app
		build:
		  context: ./app
		  dockerfile: Dockerfile.app
		volumes:
		  - /app/node_modules
          - ./app:/app
		ports:
		  - "8000:3000"
		networks:
		  - {{.projectName}}-app
		restart: unless-stopped

	postgres:
    	image: postgres:{{.tag}}
    	container_name: {{.projectName}}-db
    	restart: always
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
