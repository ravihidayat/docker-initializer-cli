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

// dockerfileCmd represents the dockerfile command
var dockerfileCmd = &cobra.Command{
	Use:   "dockerfile",
	Short: "Create a Dockerfile in the current folder.",
	Long: `The dockerfile command creates a new Dockerfile based upon
	the arguments passed. Only node and postgres currently available. 
	`,
	Example: `  
	1. docker-initializer create dockerfile 
	--image=node --env=NODE_ENV=PRODUCTION,APP=appName --tag=alpine --workdir=/app --relPath=.
	2. docker-initializer create dockerfile --image=node --tag=latest --relPath=.
	3. docker-initializer create dockerfile --image=node
	4. docker-initializer create dockerfile --image=postgres --env=POSTGRES_PASSWORD=secret
	5. docker-initializer create dockerfile --image=postgres
	`,
	Run: runCreateDockerfile,
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func init() {
	createCmd.AddCommand(dockerfileCmd)

	dockerfileCmd.PersistentFlags().String("image", "", "An image name set as the base image for the Dockerfile e.g node, postgres")
	dockerfileCmd.PersistentFlags().String("env", "", "A variable-value pair set as the environment variables e.g POSTGRES_PASSWORD=secret")
	dockerfileCmd.PersistentFlags().String("tag", "", "The tag of the corresponding image e.g alpine, latest")
	dockerfileCmd.PersistentFlags().String("workdir", "", "The working directory path e.g src/app")
	dockerfileCmd.PersistentFlags().String("relPath", "", "The relative app of the app from the Dockerfile location e.g .")
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// dockerfileCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// dockerfileCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle"

}

// func preRunCreateDockerfile(cmd *cobra.Command, args []string) {

// }

func runCreateDockerfile(cmd *cobra.Command, args []string) {
	image, _ := cmd.Flags().GetString("image")
	env, _ := cmd.Flags().GetString("env")
	tag, _ := cmd.Flags().GetString("tag")
	workdir, _ := cmd.Flags().GetString("workdir")
	relPath, _ := cmd.Flags().GetString("relPath")

	if len(image) == 0 && len(args) < 1 {
		cmd.Help()
		os.Exit(0)
	}

	createDockerfile(image, env, tag, workdir, relPath, args)
}

func createDockerfile(image string, env string, tag string, workdir string, relPath string, args []string) {
	filepath, _ := filepath.Abs("./Dockerfile")
	file, err := os.Create(filepath)
	check(err)
	defer file.Close()

	if tag == "" {
		tag = "latest"
	}

	if relPath == "" {
		relPath = "."
	}

	if workdir == "" {
		workdir = "/app"
	}

	var template string

	switch image {

	case "node":

		template = templates.NodeDockerfile
		template = strings.Replace(template, "{{.tag}}", tag, 1)

		if env != "" {
			envSplit := (envSplit(env))
			template = strings.Replace(template, "{{.env}}", envSplit, 1)
		} else if env == "" {
			template = strings.Replace(template, "{{.env}}", "ENV NODE_ENV production", 1)
		}

		template = strings.Replace(template, "{{.relPath}}", relPath, 1)
		template = strings.Replace(template, "{{.workdir}}", workdir, 1)

	case "postgres":

		template = templates.PostgresDockerfile
		template = strings.Replace(template, "{{.tag}}", tag, 1)

		if env != "" {
			envSplit := (envSplit(env))
			template = strings.Replace(template, "{{.env}}", envSplit, 1)
		} else if env == "" {
			template = strings.Replace(template, "{{.env}}", "ENV POSTGRES_PASSWORD secret", 1)
		}
	}

	_, err2 := file.WriteString(template)
	check(err2)
}

func envSplit(env string) string {
	envSplit := strings.Split(env, ",")

	var envs string = ``

	for i := 0; i < len(envSplit); i++ {
		envs += "ENV " + strings.Replace(envSplit[i], "=", " ", 1) + "\n"
	}

	return envs
}
