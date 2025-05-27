package rpc_company

import (
	"flag"
	"log"
	"net"

	"colinahealth_emr-api/internal/ports"
	pb "colinahealth_emr-api/internal/adapters/framework/left/company/grpc/pb"

	"google.golang.org/grpc"
)

// Adapter implements the GRPCPort interface
type Adapter struct {
	api		ports.APICompanyPort
	pbsvr	*PbUCS_Server
}

// Company Service Server...
type PbUCS_Server struct {	
	pb.UnimplementedCompanyServiceServer
}

// NewAdapter creates a new Adapter
func NewAdapter(api ports.APICompanyPort, pbsvr PbUCS_Server) *Adapter {
	return &Adapter{api: api, pbsvr: &pbsvr}
}

// Run registers the CompanyServiceServer to a grpcServer and serves on
// the specified port
func (grpca Adapter) Run() {
	var err error
	var server_ADDR = flag.String("server_ADDR", ":50051", "the address to connect to")

	listen, err := net.Listen("tcp", *server_ADDR)
	if err != nil {
		log.Fatalf("failed to listen on port %s: %v", *server_ADDR, err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterCompanyServiceServer(grpcServer, new(PbUCS_Server))

	if err := grpcServer.Serve(listen); err != nil {
		log.Fatalf("failed to serve gRPC server over port %s: %v", *server_ADDR, err)
	}
}