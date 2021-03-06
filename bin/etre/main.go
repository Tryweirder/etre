// Copyright 2017-2020, Square, Inc.

package main

import (
	"flag"
	"log"

	"github.com/square/etre/app"
	"github.com/square/etre/server"
)

var configFile string

func init() {
	flag.StringVar(&configFile, "config", "", "Config file")
}

func main() {
	flag.Parse()
	s := server.NewServer(app.Defaults())
	if err := s.Boot(configFile); err != nil {
		log.Fatalf("Error starting: %s", err)
	}
	if err := s.Run(); err != nil {
		log.Fatalf("Error running: %s", err)
	}
	log.Println("Etre has stopped")
}
