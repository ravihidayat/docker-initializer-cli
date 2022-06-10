/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"os"
	"path/filepath"

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

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// dockerComposeCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// dockerComposeCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func runCreateDockerCompose(cmd *cobra.Command, args []string) {
	createDockerCompose()
}

func createDockerCompose() {
	filepath, _ := filepath.Abs("./docker-compose.yml")
	file, err := os.Create(filepath)
	check(err)
	defer file.Close()
}
