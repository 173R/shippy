module github.com/173R/shippy/service-consignment

go 1.19

//replace github.com/173R/shippy/service-vessel => ../service-vessel

require (
	google.golang.org/grpc v1.51.0
	google.golang.org/protobuf v1.28.1
)