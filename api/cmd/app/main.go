package main

import (
	// http gin server
	"colinahealth_emr-api/cmd/app/http_server"
)

func main() {
	// HTTP API Gin Server
	ha := http_server.New()
	router := ha.SetupRouter()
	router.Run()
}