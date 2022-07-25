package hello

import (
	"context"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/pachyderm/bazel-poc/protos"
	"google.golang.org/protobuf/testing/protocmp"
)

func TestEcho(t *testing.T) {
	ctx := context.Background()
	s := new(Server)
	res, err := s.Echo(ctx, &protos.EchoRequest{Text: "foo"})
	if err != nil {
		t.Fatal(err)
	}
	if diff := cmp.Diff(res, &protos.EchoReply{Text: "foo"}, protocmp.Transform()); diff != "" {
		t.Fatalf("diff:\n%s", diff)
	}
}
