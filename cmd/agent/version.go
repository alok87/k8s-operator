package main

import (
	"fmt"
	"os"

	"github.com/alok87/k8s-operator/pkg/version"
	"github.com/spf13/cobra"
)

// VersionCmd represents the subcommand for root command: node-agent
var VersionCmd = &cobra.Command{
	Use:   "version",
	Short: "Show the node-agent version information",
	Long:  `Show the node-agent version information`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf(version.Format())
		os.Exit(0)
	},
}
