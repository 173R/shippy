package main

import (
	"context"
	"errors"
	pb "github.com/173R/shippy/user-service/proto/user"
	"golang.org/x/crypto/bcrypt"
)

type authableI interface {
	Decode(token string) (*CustomClaims, error)
	Encode(user *pb.User) (string, error)
}

type handler struct {
	repository   UserRepositoryI
	tokenService authableI
	pb.UnimplementedUserServiceServer
}

func (self *handler) Get(ctx context.Context, req *pb.User) (*pb.Response, error) {
	result, err := self.repository.Get(ctx, req.Id)
	if err != nil {
		return nil, err
	}

	return &pb.Response{User: UnmarshalUser(result)}, nil
}

func (self *handler) GetAll(ctx context.Context, req *pb.Request) (*pb.Response, error) {
	results, err := self.repository.GetAll(ctx)
	if err != nil {
		return nil, err
	}

	return &pb.Response{Users: UnmarshalUserCollection(results)}, nil
}

func (self *handler) Auth(ctx context.Context, req *pb.User) (*pb.Token, error) {
	user, err := self.repository.GetByEmail(ctx, req.Email)
	if err != nil {
		return nil, err
	}

	if err := bcrypt.CompareHashAndPassword(
		[]byte(user.Password),
		[]byte(req.Password),
	); err != nil {
		return nil, err
	}

	token, err := self.tokenService.Encode(req)
	if err != nil {
		return nil, err
	}

	return &pb.Token{Token: token}, nil
}

func (self *handler) Create(
	ctx context.Context,
	req *pb.User,
) (*pb.Response, error) {
	hashedPass, err := bcrypt.GenerateFromPassword(
		[]byte(req.Password),
		bcrypt.DefaultCost,
	)
	if err != nil {
		return nil, err
	}

	req.Password = string(hashedPass)
	if err := self.repository.Create(ctx, MarshalUser(req)); err != nil {
		return nil, err
	}

	// Strip the password back out, so's we're not returning it
	req.Password = ""

	return &pb.Response{User: req}, nil
}

func (self *handler) ValidateToken(
	ctx context.Context,
	req *pb.Token,
) (*pb.Token, error) {
	claims, err := self.tokenService.Decode(req.Token)
	if err != nil {
		return nil, err
	}

	if claims.User.Id == "" {
		return nil, errors.New("invalid user")
	}

	return &pb.Token{Valid: true}, nil
}
