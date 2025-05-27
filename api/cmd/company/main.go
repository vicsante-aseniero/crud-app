package main

import (
	"fmt"
	"log"

	// application
	"colinahealth_emr-api/internal/application/api"
	"colinahealth_emr-api/internal/application/core/domain/company"

	// adapters
	"colinahealth_emr-api/internal/adapters/framework/right/db"
	gRPC "colinahealth_emr-api/internal/adapters/framework/left/company/grpc"
)

func main () {
	var err error

	// Init and open database
	PG_HOST := "localhost"
	PG_USER := "colina"
	PG_PORT := 5432
	PG_DATABASE := "colina"
	PG_PASSWORD := "jajnav5@"
	dbaseDriver := "postgres"
	dsourceName := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s connect_timeout=10 sslmode=disable", PG_HOST, PG_PORT, PG_USER, PG_PASSWORD, PG_DATABASE)

	dbAdapter, err := db.NewAdapter(dbaseDriver, dsourceName)
	if err != nil {
		log.Fatalf("failed to initiate dbase connection: %v", err)
	}
	defer dbAdapter.CloseDbConnection()

	// core
	core := company.New()

	// application
	applicationAPI := api.NewApplication(dbAdapter, core)

	// adapter
	gRPCAdapter := gRPC.NewAdapter(applicationAPI, gRPC.PbUCS_Server{})
	gRPCAdapter.Run()
}