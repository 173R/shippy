package main

import (
	"context"
	pb "github.com/173R/shippy/service-consignment/proto/consignment"
	vesselProto "github.com/173R/shippy/service-vessel/proto/vessel"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
	"os"
)

const (
	defaultHost = "datastore:27017"
)

const port = ":50051"
const vesselAddress = "vessel:50052"

func main() {
	db_uri := os.Getenv("DB_HOST")
	if db_uri == "" {
		db_uri = defaultHost
	}

	client, err := CreateClient(context.Background(), db_uri, 0)
	if err != nil {
		log.Panic(err)
	}

	defer client.Disconnect(context.Background())

	consignmentCollection := client.
		Database("shippy").
		Collection("consignment")

	consignmentRepo := &MongoRepository{consignmentCollection}

	listener, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()

	//Создание клиента для сервера vessels
	conn, err := grpc.Dial(vesselAddress, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Did not connect: %v", err)
	}
	defer conn.Close()

	vesselClient := vesselProto.NewVesselServiceClient(conn)

	pb.RegisterShippingServiceServer(grpcServer, &handler{
		consignmentRepo: consignmentRepo,
		vesselClient:    vesselClient,
	})

	reflection.Register(grpcServer)

	log.Println("Running on port:", port)
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
