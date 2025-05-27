package http_server

import (
	"context"
	"flag"
	"log"
	"net/http"

	pb "colinahealth_emr-api/internal/adapters/framework/left/company/grpc/pb"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

// http gin server
type HTTP_API struct {
}

// NewHTTPServer creates a new HTTP Gin 
func New() *HTTP_API {
	return &HTTP_API{}
}

type Company struct {
	ID				int			`json:"id"`
	UUID			string		`json:"uuid"`
	Name			string		`json:"name"`
	Contact			string		`json:"contact"`
	Website			string		`json:"website"`
	Email			string		`json:"email"`
	Country			string		`json:"country"`
	State			string		`json:"state"`
	Zip				string		`json:"zip"`
}

var companies = []Company{}

var addr_gRPC = flag.String("addr_gRPC", ":50051", "the address to connect to")

func (ha HTTP_API) SetupRouter() *gin.Engine {

	router := gin.Default()
	router.Use(cors.Default())

	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Welcome to Colina Health EMR APIs",
		})
	})

	router.GET("/companies", getCompanies)
	router.GET("/company/:id", getCompany)
	router.POST("/company", setCompany)	
	router.PATCH("/company/:id", editCompany)
	router.DELETE("/company/:id", removeCompany)

	return router
}

func getCompanies(c *gin.Context) {
	
	conn, err := grpc.NewClient(*addr_gRPC, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	//client := pb.NewCompanyServiceClient(conn)

	c.JSON(200, gin.H{
		"message": "Read all Companies API not yet Implemented!",
	})
}

func getCompany(c *gin.Context) {
	
	conn, err := grpc.NewClient(*addr_gRPC, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	//client := pb.NewCompanyServiceClient(conn)

	c.JSON(200, gin.H{
		"message": "Read a Company API not yet Implemented!",
	})
}

func setCompany(c *gin.Context) {
	var nc Company

	if err := c.BindJSON(&nc); err != nil {
		return
	}

	// Log the newCompany JSON to the console
	log.Println("newCompany=", nc)

	conn, err := grpc.NewClient(*addr_gRPC, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	client := pb.NewCompanyServiceClient(conn)

	req := &pb.RequestVariables{Name: nc.Name, Contact: nc.Contact, Website: nc.Website, Email: nc.Email, Country: nc.Country, State: nc.State, Zip: nc.Zip}
	resp, err := client.AddNewCompany(context.Background(), req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	companies = append(companies, nc)

	c.JSON(http.StatusOK, gin.H{"message": resp.GetCompany()})
}

func editCompany(c *gin.Context) {
	
	conn, err := grpc.NewClient(*addr_gRPC, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	//client := pb.NewCompanyServiceClient(conn)

	c.JSON(200, gin.H{
		"message": "Updating a Company API not yet Implemented!",
	})
}

func removeCompany(c *gin.Context) {
	
	conn, err := grpc.NewClient(*addr_gRPC, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	//client := pb.NewCompanyServiceClient(conn)

	c.JSON(200, gin.H{
		"message": "Removing a Company API not yet Implemented!",
	})
}