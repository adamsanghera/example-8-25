package main

import (
	"os"

	"github.com/adamsanghera/example-8-25/pkg/web"
)

func main() {
	cfg := &web.Config{
		Addr: os.Getenv("HOST"),
	}

	webSrv, err := web.New(cfg)
	if err != nil {
		panic(err)
	}

	err = webSrv.Start()
	if err != nil {
		panic(err)
	}
}
