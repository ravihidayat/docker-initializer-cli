package templates

var NodeDockerfile string = `
FROM node:{{.tag}}
{{.env}}
WORKDIR {{.workdir}}
RUN npm install
COPY {{.relPath}} .
EXPOSE 3000
CMD npm start
`
