module github.com/173R/shippy/shippy-service-consignment

go 1.19

replace github.com/173R/shippy/shippy-service-vessel => ../shippy-service-vessel

require (
	github.com/173R/shippy/shippy-service-vessel/proto/vessel v0.0.1
	google.golang.org/grpc v1.51.0
	google.golang.org/protobuf v1.28.1
)