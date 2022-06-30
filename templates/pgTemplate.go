package templates

var PostgresDockerfile string = `
FROM postgres:{{.tag}}
{{.env}}
`
