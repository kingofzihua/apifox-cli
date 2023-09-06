package main

import (
	"log"

	"github.com/spf13/cobra"

	"github.com/kingofzihua/apifox-cli/internal/importc"
)

func init() {
	// import
	rootCmd.AddCommand(importc.ImportCmd)
}

var rootCmd = &cobra.Command{
	Use:   "apifox-cli",
	Short: "apifox-cli is a command-line tool for apifox",
	Long: `apifox-cli is a command-line tool for apifox
@see https://apifox.com`,
	Run: func(cmd *cobra.Command, args []string) {
		err := cmd.Help()
		if err != nil {
			log.Fatal(err)
		}
	},
	Version: Version,
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}
