package main

import (
	"context"
	"encoding/json"

	"github.com/kusipay/api-go-messenger/adapter/output"
	"github.com/kusipay/api-go-messenger/middleware"
	"github.com/mefellows/vesper"
)

func handler(ctx context.Context, event any) (string, error) {
	logger := output.NewLoggerRepository()

	bts, err := json.MarshalIndent(event, "", "  ")
	if err != nil {
		logger.Error("handler |", err.Error())
		return "", err
	}

	logger.Info("handler |", "CALLED WITH", string(bts))

	return "Hello World", nil
}

func main() {
	v := vesper.New(handler).Use(middleware.EventLog())

	v.Start()
}
