// Package hello implements the Hello RPC service.
package hello

import (
	"context"

	"github.com/pachyderm/bazel-poc/protos"
)

type Server struct {
	protos.UnimplementedHelloServer
	Message string
}

func (s *Server) Echo(ctx context.Context, in *protos.EchoRequest) (*protos.EchoReply, error) {
	return &protos.EchoReply{
		Text: in.GetText(),
	}, nil
}

func (s *Server) MOTD(ctx context.Context, in *protos.MOTDRequest) (*protos.MOTDReply, error) {
	return &protos.MOTDReply{
		Msg: s.Message,
	}, nil
}
