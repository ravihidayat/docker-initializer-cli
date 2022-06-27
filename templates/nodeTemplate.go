package templates

var NodeDockerfile string = `
FROM postgres:{{.tag}}
{{.env}}
WORKDIR {{.workdir}}
RUN npm install
COPY {{.relPath}} .
EXPOSE 3000
CMD npm start
`
