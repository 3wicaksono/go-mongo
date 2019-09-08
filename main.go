package main

import (
	"go-mongo/commands"
	"go-mongo/infrastructures"

	log "github.com/sirupsen/logrus"
)

func main() {
	// initial configurations
	log.SetFormatter(&log.JSONFormatter{})
	infrastructures.SetConfig()
	command := commands.NewCommandEngine()
	command.Run()
}
