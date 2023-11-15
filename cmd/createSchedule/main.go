package main

import (
	"github.com/kusipay/api-go-messenger/adapter/input"
	"github.com/kusipay/api-go-messenger/middleware"
	"github.com/mefellows/vesper"
)

func main() {
	v := vesper.New(input.CreateScheduleHandler).
		Use(middleware.EventLog())

	v.Start()
}
