package main

import (
	"context"

	"github.com/kusipay/api-go-messenger/adapter/output"
	"github.com/kusipay/api-go-messenger/middleware"
	"github.com/mefellows/vesper"
)

func Handler(ctx context.Context) (string, error) {
	logger := output.NewLogger()

	logger.Info("Handler |", "Hello, World!")

	return "Hello, World!", nil
}

func main() {
	v := vesper.New(Handler).Use(middleware.EventLog())

	v.Start()
}
