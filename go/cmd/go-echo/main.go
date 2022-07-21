package main

import (
	"github.com/jrockway/opinionated-server/server"
)

func main() {
	server.AppName = "echo-go"
	server.Setup()
	server.ListenAndServe()
}
