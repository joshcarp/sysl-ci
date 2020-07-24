package server

import (
	"context"

	"github.com/joshcarp/sysl-ci/gen/simple"
)

func Get(ctx context.Context, req *simple.GetRequest, client simple.GetClient) (*simple.Welcome, error) {
	return &simple.Welcome{Content: "Hello, World"}, nil
}
