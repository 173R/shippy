package main

import (
	"context"
	"encoding/json"
	pb "github.com/173R/shippy/service-consignment/proto/consignment"
	"google.golang.org/grpc"
	"log"
	"os"
	"time"
)

const (
	//consignmentServiceAddr = "127.0.0.1:50051"
	consignmentServiceAddr = "consignment:50051"
	defaultFilename        = "consignment.json"
)

// Парсим коносамент который хранится в json тут на клиенте
func parseFile(file string) (*pb.Consignment, error) {
	var consignment *pb.Consignment
	data, err := os.ReadFile(file)
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(data, &consignment); err != nil {
		return nil, err
	}

	return consignment, err
}

func main() {
	//Коннектимся к серверу
	conn, err := grpc.Dial(consignmentServiceAddr, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Did not connect: %v", err)
	}
	defer conn.Close()

	client := pb.NewShippingServiceClient(conn)
	file := defaultFilename
	if len(os.Args) > 1 {
		file = os.Args[1]
	}

	time.Sleep(time.Second * 1)

	consignment, err := parseFile(file)
	if err != nil {
		log.Fatalf("Could not parse file: %v", err)
	}

	res, err := client.CreateConsignment(context.Background(), consignment)
	if err != nil {
		log.Fatalf("Could not greet: %v", err)
	}
	log.Printf("Created: %t", res.Created)

	getAll, err := client.GetConsignments(
		context.Background(), &pb.GetRequest{},
	)
	if err != nil {
		log.Fatalf("Could not list consignments: %v", err)
	}

	for _, v := range getAll.Consignments {
		log.Println(v)
	}
}
