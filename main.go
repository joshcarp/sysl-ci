// main.go
package main

import (
	"context"

	"github.com/joshcarp/sysl-ci/pkg/server"
)

func main() {
	server.LoadServices(context.Background())
}
