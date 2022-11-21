module github.com/173R/shippy/service-consignment

go 1.19

//replace github.com/173R/shippy/service-vessel => ../service-vessel

require (
	github.com/173R/shippy/service-vessel v0.0.0-20221121104014-e9df736b3f0a
	google.golang.org/grpc v1.51.0
	google.golang.org/protobuf v1.28.1
)

require (
	github.com/golang/protobuf v1.5.2 // indirect
	golang.org/x/net v0.2.0 // indirect
	golang.org/x/sys v0.2.0 // indirect
	golang.org/x/text v0.4.0 // indirect
	google.golang.org/genproto v0.0.0-20221118155620-16455021b5e6 // indirect
)
