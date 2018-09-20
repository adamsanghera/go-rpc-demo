package greeter

import (
	"context"
	"errors"
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"

	"github.com/adamsanghera/go-rpc-demo/greeter/greeterpb"
	"google.golang.org/grpc"

	"github.com/spf13/cobra"
)

// sendCmd represents the start command
var sendCmd = &cobra.Command{
	Use:   "send [ipaddr:port] [your location name]",
	Short: "Sends a greeting to a greeter rpc daemon",
	Long:  `Sends a greeting to a greeter rpc daemon, located at the address the argument, advertising for a given location`,
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) != 2 {
			return errors.New("\n  Incorrect number of arguments.  See `greeter start -h`")
		}
		if !strings.Contains(args[0], ":") {
			return errors.New("\n  Please include both an IPv4 Address and a port number")
		}

		if strings.Split(args[0], ":")[0] == "localhost" {
			return nil
		}

		if strings.Count(args[0], ".") != 3 {
			return errors.New("\n  Please provide a valid IPv4 Address and a port number")
		}

		for idx, num := range strings.Split(args[0], ".") {
			if idx == 3 {
				num = strings.Split(num, ":")[0]
			}

			n, err := strconv.Atoi(num)
			if err != nil {
				return err
			}

			if n < 0 || n > 255 {
				return errors.New("\n  Please provide a valid IPv4 Address and a port number")
			}
		}

		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		conn, err := grpc.Dial(args[0], grpc.WithInsecure())
		if err != nil {
			log.Println(err)
			return
		}

		gc := greeterpb.NewGreeterClient(conn)

		resp, err := gc.Greet(context.Background(), &greeterpb.GreetRequest{
			OpeningGreeting: &greeterpb.Greeting{
				GreeterName: args[1],
				Greeting:    fmt.Sprintf("Howdy! I represent %s", args[1]),
				Time:        &greeterpb.Time{UnixTime: time.Now().In(time.Local).Unix()},
			},
		})

		if err != nil {
			log.Println(err)
			return
		}

		respT := time.Unix(resp.ResponseGreeting.Time.UnixTime, 0)

		log.Printf(`Received response: 
	'%s' from %s, which took %v`, resp.ResponseGreeting.Greeting, resp.ResponseGreeting.GreeterName, time.Since(respT))

	},
}

func init() {
	RootCmd.AddCommand(sendCmd)
}
