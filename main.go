package main

import (
	"context"
	"os"
	"os/signal"
	"github.com/montcao/gadget/cmd"
)

func main() {

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
	defer stop()
	if err := cmd.Root.ExecuteContext(ctx); err != nil {
		os.Exit(1)
	}
}
