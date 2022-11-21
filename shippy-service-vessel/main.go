package main

import (
	"context"
	pb "github.com/173R/shippy/shippy-service-vessel/proto/vessel"
	"log"

	//"sync"
	"errors"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"net"
)

type RepositoryI interface {
	FindAvailable(*pb.Specification) (*pb.Vessel, error)
}

type VesselRepository struct {
	vessels []*pb.Vessel
}

func (repo *VesselRepository) FindAvailable(spec *pb.Specification) (*pb.Vessel, error) {
	for _, vessel := range repo.vessels {
		if vessel.Capacity >= spec.Capacity && vessel.MaxWeight >= spec.MaxWeight {
			return vessel, nil
		}
	}

	return nil, errors.New("No vessel found by that spec")
}

type vesselService struct {
	repo RepositoryI
	pb.UnimplementedVesselServiceServer
}

func (s *vesselService) FindAvailable(
	ctx context.Context,
	in *pb.Specification,
) (*pb.Response, error) {
	vessel, err := s.repo.FindAvailable(in)
	if err != nil {
		return nil, err
	}

	return &pb.Response{Vessel: vessel}, nil
}

const port = ":50052"

func main() {
	vessels := []*pb.Vessel{
		&pb.Vessel{
			Id:        "vessel001",
			Name:      "Boaty McBoatface",
			MaxWeight: 200000,
			Capacity:  500,
		},
	}

	repo := &VesselRepository{vessels}

	listener, err := net.Listen("tcp", port)
	if err != nil {
		log.Panic(err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterVesselServiceServer(grpcServer, &vesselService{repo: repo})
	reflection.Register(grpcServer)
	log.Println("Running on port:", port)

	if err := grpcServer.Serve(listener); err != nil {
		log.Panic(err)
	}
}
