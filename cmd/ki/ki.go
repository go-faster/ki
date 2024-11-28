package main

import (
	"context"
	"fmt"
	"os"

	"github.com/go-faster/sdk/zctx"
	"go.uber.org/zap"

	"github.com/go-faster/ki/k8s/release"
)

func main() {
	ctx := context.Background()
	lg, err := zap.NewDevelopment()
	if err != nil {
		panic(err)
	}
	ctx = zctx.Base(ctx, lg)

	v, err := release.Stable(ctx)
	if err != nil {
		lg.Error("Fetch release", zap.Error(err))
		os.Exit(2)
		return
	}

	fmt.Println("Got version:", v)
}
