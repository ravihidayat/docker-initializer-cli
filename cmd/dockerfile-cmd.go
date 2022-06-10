/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
)

// dockerfileCmd represents the dockerfile command
var dockerfileCmd = &cobra.Command{
	Use:   "dockerfile",
	Short: "Creates a Dockerfile.",
	Long: `Creates a new Dockerfile based upon the base
	configuration for a `,
	Example: `docker-initializer create dockerfile`,
	Run:     runCreateDockerfile,
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func init() {
	createCmd.AddCommand(dockerfileCmd)

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
	createDockerfile()
}

func createDockerfile() {
	filepath, _ := filepath.Abs("./Dockerfile")
	file, err := os.Create(filepath)
	check(err)
	defer file.Close()
}
