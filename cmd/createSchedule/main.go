package main

import (
	"github.com/kusipay/api-go-messenger/adapter/input/createschedule"
	"github.com/kusipay/api-go-messenger/middleware"
	"github.com/mefellows/vesper"
)

func main() {
	v := vesper.New(createschedule.Handler).
		Use(middleware.EventLog())

	v.Start()
}
