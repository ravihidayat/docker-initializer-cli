package templates

var PostgresDockerfile string = `
FROM postgres:{{.tag}}
{{.env}}
`

var PostgresDockerCompose string = `

`
