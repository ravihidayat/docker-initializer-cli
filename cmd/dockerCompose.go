/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"os"
	"path/filepath"
	"strings"

	"github.com/ravihidayat/docker-initializer-cli/templates"
	"github.com/spf13/cobra"
)

// dockerComposeCmd represents the dockerCompose command
var dockerComposeCmd = &cobra.Command{
	Use:   "docker-compose",
	Short: "Creates a docker-compose.yml.",
	Long:  ``,
	Run:   runCreateDockerCompose,
}

func init() {
	createCmd.AddCommand(dockerComposeCmd)

	dockerComposeCmd.PersistentFlags().String("projectName", "", "The project name")
	dockerComposeCmd.PersistentFlags().String("stack", "", "The stack name preferred to be built upon e.g mern, pern")
	dockerComposeCmd.PersistentFlags().String("nodeTag", "", "The image tag of Node e.g latest, alpine")
	dockerComposeCmd.PersistentFlags().String("env", "", "A variable-value pair set as the environment variables e.g POSTGRES_PASSWORD=secret")
	dockerComposeCmd.PersistentFlags().String("dbTag", "", "The tag for mongo and postgres selected e.g latest, alpine")
	dockerComposeCmd.PersistentFlags().String("dbUsername", "", "The name of the database user e.g root")
	dockerComposeCmd.PersistentFlags().String("dbPassword", "", "The password set for the particular user e.g secret")
	dockerComposeCmd.PersistentFlags().String("dbName", "", "The database name e.g my-db")
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// dockerComposeCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// dockerComposeCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func runCreateDockerCompose(cmd *cobra.Command, args []string) {
	projectName, _ := cmd.Flags().GetString("projectName")
	stack, _ := cmd.Flags().GetString("stack")
	nodeTag, _ := cmd.Flags().GetString("nodeTag")
	env, _ := cmd.Flags().GetString("env")
	dbUsername, _ := cmd.Flags().GetString("dbUsername")
	dbPassword, _ := cmd.Flags().GetString("dbPassword")
	dbName, _ := cmd.Flags().GetString("dbName")
	dbTag, _ := cmd.Flags().GetString("dbTag")

	if len(projectName) == 0 {
		cmd.Help()
		os.Exit(0)
	}

	if len(stack) == 0 && len(args) < 1 {
		cmd.Help()
		os.Exit(0)
	}

	createDockerCompose(projectName, stack, nodeTag, env, dbUsername, dbPassword, dbName, dbTag)
}

func createDockerCompose(projectName string, stack string, nodeTag string, env string, dbUsername string, dbPassword string, dbName string, dbTag string) {

	if nodeTag == "" {
		nodeTag = "latest"
	}

	if dbUsername == "" {
		dbUsername = "root"
	}

	if dbPassword == "" {
		dbPassword = "secret"
	}

	if dbName == "" {
		dbName = projectName + "-db"
	}

	var dockerComposeTemplate string
	var envs string

	switch stack {
	case "mern":
		dockerComposeTemplate = templates.MernDockerCompose
		dockerComposeTemplate = strings.Replace(dockerComposeTemplate, "{{.dbTag}}", dbTag, 1)
		env, _ := filepath.Abs(".env")
		envFile, err := os.Create(env)
		check(err)
		defer file.Close()

		envs = `
		MONGO_INITDB_ROOT_USERNAME={{.dbUsername}}
        MONGO_INITDB_ROOT_PASSWORD={{.dbPassword}}
		`

		if dbUsername != "" {
			envs = strings.Replace(envs, "{{.dbUsername}}", dbUsername, 1)
		} else if dbUsername == "" {
			envs = strings.Replace(envs, "{{.dbUsername}}", "root", 1)
		}

		if dbUsername != "" {
			envs = strings.Replace(envs, "{{.dbPassword}}", dbPassword, 1)
		} else if dbUsername == "" {
			envs = strings.Replace(envs, "{{.dbPassword}}", "secret", 1)
		}

		_, err3 := envFile.WriteString(envs)
		check(err3)

	case "pern":
		dockerComposeTemplate = templates.PernDockerCompose
		dockerComposeTemplate = strings.Replace(dockerComposeTemplate, "{{.projectName}}", projectName, 8)
	}

	_, err2 := file.WriteString(dockerComposeTemplate)
	check(err2)
}
