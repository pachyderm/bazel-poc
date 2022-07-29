package main

import (
	"github.com/jrockway/opinionated-server/server"
	"github.com/pachyderm/bazel-poc/go/pkg/hello"
	"github.com/pachyderm/bazel-poc/protos"
	"go.uber.org/zap"
	"google.golang.org/grpc"
)

type Config struct {
	Message string `long:"message" env:"MESSAGE" description:"The message of the day." default:"Hello, world!"`
}

func main() {
	server.AppName = "echo-go"

	var cfg Config
	server.AddFlagGroup("config", &cfg)
	server.Setup()

	hs := &hello.Server{
		Message: cfg.Message,
	}
	server.AddService(func(s *grpc.Server) {
		protos.RegisterHelloServer(s, hs)
	})

	zap.L().Info("message of the day", zap.String("message", hs.Message))
	server.ListenAndServe()
}
