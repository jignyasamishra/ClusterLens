/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"os"

	"github.com/jignyasamishra/ClusterLens/cmd/scan"
	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "clusterlens",
	Short: "Kubernetes Cluster Scanner: A quick summary generator Resources",
	Long:  ``,
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {

	rootCmd.AddCommand(scan.ScanCmd)

	rootCmd.CompletionOptions.HiddenDefaultCmd = true
}
