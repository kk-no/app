package main

import (
	"context"
	"log"

	"github.com/kk-no/expapp/gw/config"
	"github.com/kk-no/expapp/gw/server"
)

func main() {
	ctx := context.Background()

	if err := config.Load(); err != nil {
		log.Fatal(err)
	}
	log.Fatal(run(ctx))
}

func run(ctx context.Context) error {
	conf := config.Conf
	s, err := server.NewHTTPServer(ctx)
	if err != nil {
		return err
	}
	return s.Serve(conf.Port)
}
