package main

import (
	"context"
	"log"
	"net"
	"sync"

	pb "github.com/173R/shippy/service-consignment/proto/consignment"
	//vesselProto "github.com/173R/shippy/service-vessel/proto/vessel"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

const port = ":50051"

type RepositoryI interface {
	Create(*pb.Consignment) (*pb.Consignment, error)
	GetAll() []*pb.Consignment
}

type ConsignmentRepository struct {
	mu           sync.RWMutex
	consignments []*pb.Consignment
}

// Создание коносамента
func (repo *ConsignmentRepository) Create(consignment *pb.Consignment) (*pb.Consignment, error) {
	repo.mu.Lock()
	updated := append(repo.consignments, consignment)
	repo.consignments = updated
	repo.mu.Unlock()
	return consignment, nil
}

func (repo *ConsignmentRepository) GetAll() []*pb.Consignment {
	return repo.consignments
}

type service struct {
	repo RepositoryI
	//vesselClient vesselProto.VesselServiceClient
	pb.UnimplementedShippingServiceServer
}

// Создаём методы которые отражают методы из proto
func (s *service) CreateConsignment(
	ctx context.Context,
	req *pb.Consignment,
) (*pb.Response, error) {
	consignment, err := s.repo.Create(req)
	if err != nil {
		return nil, err
	}

	return &pb.Response{Created: true, Consignment: consignment}, nil
}

func (s *service) GetConsignments(
	context.Context,
	*pb.GetRequest,
) (*pb.Response, error) {
	consignments := s.repo.GetAll()
	return &pb.Response{Consignments: consignments}, nil
}

func main() {
	repo := &ConsignmentRepository{}

	listener, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterShippingServiceServer(grpcServer, &service{repo: repo})

	reflection.Register(grpcServer)

	log.Println("Running on port:", port)
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
