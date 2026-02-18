package main

import (
	"flag"

	_ "github.com/eriksejr/promsaint/logging"
	"github.com/eriksejr/promsaint/server"
	log "github.com/sirupsen/logrus"
)

var (
	Version   string
	BuildTime string
)

func main() {
	flag.Parse()

	log.WithFields(log.Fields{
		"version": Version,
		"build":   BuildTime,
	}).Warn("Starting Promsaint!")
	s := server.NewPromsaint()
	s.Start()
}
