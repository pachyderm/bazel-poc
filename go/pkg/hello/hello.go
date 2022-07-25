package hello

import (
	"context"

	"github.com/pachyderm/bazel-poc/protos"
)

type Server struct {
	protos.UnimplementedHelloServer
}

func (s *Server) Echo(ctx context.Context, in *protos.EchoRequest) (*protos.EchoReply, error) {
	return &protos.EchoReply{
		Text: in.GetText(),
	}, nil
}
