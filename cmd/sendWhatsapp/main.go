package main

import (
	"github.com/kusipay/api-go-messenger/adapter/input/sendwhatsapp"
	"github.com/kusipay/api-go-messenger/middleware"
	"github.com/mefellows/vesper"
)

func main() {
	v := vesper.New(sendwhatsapp.Handler).
		Use(middleware.EventLog())

	v.Start()
}
