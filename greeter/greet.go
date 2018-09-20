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

	log.Println(in.OpeningGreeting.String()) // isn't it nice to have string serialization for free?

	hourDifference := time.Now().In(time.Local).Sub(t).Round(time.Hour)

	return &greeterpb.GreetResponse{
		ResponseGreeting: &greeterpb.Greeting{
			Greeting: fmt.Sprintf(
				"Hello, from %s, it seems that your time zone is %v hours off from mine.",
				g.location, hourDifference.Hours()),
			GreeterName: name,
			Time:        &greeterpb.Time{UnixTime: time.Now().In(time.Local).Unix()},
		},
	}, nil
}
