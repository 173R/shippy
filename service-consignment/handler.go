package main

import (
	"context"
	"errors"
	pb "github.com/173R/shippy/service-consignment/proto/consignment"
	vesselProto "github.com/173R/shippy/service-vessel/proto/vessel"
)

type handler struct {
	consignmentRepo consignmentRepoI
	vesselClient    vesselProto.VesselServiceClient
	pb.UnimplementedShippingServiceServer
}

func (self *handler) CreateConsignment(
	ctx context.Context,
	req *pb.Consignment,
) (*pb.Response, error) {
	vesselResponse, err := self.vesselClient.FindAvailable(
		ctx,
		&vesselProto.Specification{
			MaxWeight: req.Weight,
			Capacity:  int32(len(req.Containers)),
		},
	)

	if err != nil {
		return nil, err
	}
	if vesselResponse == nil {
		return nil, errors.New("error fetching vessel, returned nil")
	}

	req.VesselId = vesselResponse.Vessel.Id
	if err = self.consignmentRepo.Create(ctx, MarshalConsignment(req)); err != nil {
		return nil, err
	}

	return &pb.Response{Created: true}, nil
}

func (self *handler) GetConsignments(
	ctx context.Context,
	req *pb.GetRequest,
) (*pb.Response, error) {
	consignments, err := self.consignmentRepo.GetAll(ctx)
	if err != nil {
		return nil, err
	}
	return &pb.Response{
		Consignments: UnmarshalConsignmentCollection(consignments),
	}, nil
}
