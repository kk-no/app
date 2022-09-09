package main

import (
	"log"

	"github.com/kk-no/expapp/app/config"
	"github.com/kk-no/expapp/app/server"
)

func main() {
	if err := config.Load(); err != nil {
		log.Fatalf("load config failed: %s", err)
	}
	log.Fatal(run())
}

func run() error {
	conf := config.Conf
	s, err := server.NewGRPCServer()
	if err != nil {
		return err
	}
	return s.Serve(conf.Port)
}
