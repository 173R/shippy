package main

import (
	"fmt"
	proto "github.com/173R/shippy/user-service/proto"
)

func createUser(ctx context.Context, service micro.Service, user *proto.User) error {
	client := proto.NewUserService("shippy.service.user", service.Client())
	rsp, err := client.Create(ctx, user)
	if err != nil {
		return err
	}

	// print the response
	fmt.Println("Response: ", rsp.User)

	return nil
}

func main() {

}
