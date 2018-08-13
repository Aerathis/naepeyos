package main

import (
	"context"
	"flag"
	"fmt"

	"github.com/aerathis/naepeyos/common"
	"github.com/aerathis/naepeyos/pkg/types"
)

var (
	serverPort = flag.Int("port", 8180, "Port the server will listen on")
	serverHost = flag.String("host", "localhost", "Host server will listen on")
)

// NaepeyosServer is a structure that parses incoming requests and parses them
// to necessary parts of the application.
type NaepeyosServer struct {
	runners []types.RunnerDescription
}

// RegisterRunner implements the RegisterRunner RPC method for the runner server
func (s *NaepeyosServer) RegisterRunner(ctx context.Context, description common.RunnerDescription) (*common.RegisterResponse, error) {

	return nil, nil
}

// DeregisterRunner implements the DeregisterRunner RPC method for the runner server
func (s *NaepeyosServer) DeregisterRunner(ctx context.Context, runnerID common.RunnerId) (*common.DeregisterResponse, error) {
	return nil, nil
}

func main() {
	flag.Parse()
	fmt.Println("Testing")
}
