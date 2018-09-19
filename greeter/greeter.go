package greeter

import (
	"errors"
	"net"

	"github.com/adamsanghera/go-rpc-demo/greeter/greeterpb"
	"google.golang.org/grpc"
)

// Greeter is a service implementing the service defined in service.proto
type Greeter struct {
	location   string
	listenPort string
}

// NewGreeter creates and returns a new greeter, given a greeter.config ojbect
func NewGreeter(cfg *Config) (*Greeter, error) {
	if cfg.Location == "poly" {
		return nil, errors.New("Pick a better location") // easter egg :)
	}

	return &Greeter{
		location:   cfg.Location,
		listenPort: cfg.Port,
	}, nil
}

// Run registers a greeter object as a gRPC service, and returns an error if it fails.
func (g *Greeter) Run() error {
	srv := grpc.NewServer()
	greeterpb.RegisterGreeterServer(srv, g)

	lis, err := net.Listen("tcp", ":"+g.listenPort)
	if err != nil {
		return err
	}

	return srv.Serve(lis)
}
