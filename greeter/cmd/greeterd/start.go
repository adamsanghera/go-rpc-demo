package greeterd

import (
	"errors"
	"log"
	"strconv"

	"github.com/adamsanghera/go-rpc-demo/greeter"
	"github.com/spf13/cobra"
)

// startCmd represents the start command
var startCmd = &cobra.Command{
	Use:   "start [port number] [location name]",
	Short: "Starts the greeter rpc daemon",
	Long:  `Starts the greeter rpc daemon on a port specified in the argument, advertising for a given location`,
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) != 2 {
			return errors.New("\n  Incorrect number of arguments.  See `greeterd start -h`")
		}
		port, err := strconv.Atoi(args[0])
		if err != nil {
			return err
		}
		if port > 65535 || port < 0 {
			return errors.New("\n  Invalid port number")
		}
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		log.Printf("Starting up on port %s, representing %s\n", args[0], args[1])
		cfg := &greeter.Config{
			Port:     args[0],
			Location: args[1],
		}

		g, err := greeter.NewGreeter(cfg)
		if err != nil {
			log.Printf("Encountered error creating greeter: %s", err.Error())
		}

		err = g.Run()
		if err != nil {
			log.Printf("Encountered error while running greeter: %s", err.Error())
		}
	},
}

func init() {
	RootCmd.AddCommand(startCmd)

}
