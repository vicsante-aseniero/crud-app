package main

import (
	"flag"
	"context"
	"log"
	"net/http"

	pb "colinahealth_emr-api/cmd/app/rpc/pb"

	"github.com/gin-gonic/gin"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var addr_Http = flag.String("addr_Http", ":9090", "HTTP Gin Server address to connect to")
var addr_gRPC = flag.String("addr_gRPC", ":40051", "gRPC Server address to connect to")

func main() {
	//conn, err := grpc.Dial("localhost:40051", grpc.WithInsecure())
	conn, err := grpc.NewClient(*addr_gRPC, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	client := pb.NewMyServiceClient(conn)

	r := gin.Default()

	r.GET("/hello", func(c *gin.Context) {

		name := c.Query("name")
		req := &pb.MyRequest{Name: name}
		resp, err := client.MyMethod(context.Background(), req)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.String(http.StatusOK, "Welcome %s to ColinaHealthEMR", resp.GetMessage())
	})

	r.GET("/hello/:name", func(c *gin.Context) {
		name := c.Param("name")
		req := &pb.MyRequest{Name: name}
		resp, err := client.MyMethod(context.Background(), req)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"message": resp.GetMessage()})
	})

	r.Run(*addr_Http)
}