package main

import (
	"github.com/jrockway/opinionated-server/server"
	"github.com/pachyderm/bazel-poc/protos"
	"google.golang.org/grpc"
)

func main() {
	server.AppName = "echo-go"
	server.AddService(func(s *grpc.Server) {
		protos.RegisterHelloServer(s, &protos.UnimplementedHelloServer{})
	})
	server.Setup()
	server.ListenAndServe()
}
