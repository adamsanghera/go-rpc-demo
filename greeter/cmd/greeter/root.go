package greeter

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var cfgFile string

// RootCmd represents the base command when called without any subcommands
var RootCmd = &cobra.Command{
	Use:   "greeter",
	Short: "A cli tool for sending greet rpc requests to a greeterd service",
	Long: `A cli tool for sending greet rpc requests to a greeterd service

Sends a request over the network to a distant greeter, and prints the reply.`,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
