package main

import (
	"github.com/golang/glog"
	"github.com/alok87/k8s-operator/pkg/agent"

	"github.com/spf13/cobra"
)

// RootCmd represents the base command when called without any subcommands
var RootCmd = &cobra.Command{
	Use:   "node-agent",
	Short: "Node agent to figure out the node health",
	Long:  `Node agent to figure out the node health`,
	Run: func(cmd *cobra.Command, args []string) {
		glog.Info("Starting node-agent")
		agent.Start()
	},
}

// Execute adds all child commands to the root command sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := RootCmd.Execute(); err != nil {
		glog.Fatalf("Agent exited with erorr: %v", err)
	}
}

func init() {
	RootCmd.AddCommand(VersionCmd)
}

func main() {
	RootCmd.Execute()
}
