package greeter

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/adamsanghera/go-rpc-demo/greeter/greeterpb"
)

// Greet generates a response based on the local time of the caller.
func (g *Greeter) Greet(ctx context.Context, in *greeterpb.GreetRequest) (*greeterpb.GreetResponse, error) {
	name, err := os.Hostname()
	if err != nil {
		return nil, err // isn't it nice that we can return a language-native error as a separate return value?
	}

	t := time.Unix(in.OpeningGreeting.Time.UnixTime, 0) // isn't unix time great?

	log.Printf("Received a greeting %s, representing %s, which took %v to get here",
		in.OpeningGreeting.Greeting, in.OpeningGreeting.GreeterName, time.Since(t)) // isn't it nice to have string serialization for free?

	return &greeterpb.GreetResponse{
		ResponseGreeting: &greeterpb.Greeting{
			Greeting: fmt.Sprintf(
				"Hello, from %s, it seems that it took you %v to reach me",
				g.location, time.Since(t)),
			GreeterName: name,
			Time:        &greeterpb.Time{UnixTime: time.Now().In(time.Local).Unix()},
		},
	}, nil
}
