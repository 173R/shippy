module github.com/173R/shippy/service-consignment

go 1.19

replace github.com/173R/shippy/service-vessel => ../service-vessel

require (
	google.golang.org/grpc v1.51.0
	google.golang.org/protobuf v1.28.1
)

require (
	github.com/173R/shippy/service-vessel v0.0.0-20221121083744-bb385b0553f1 // indirect
	github.com/golang/protobuf v1.5.2 // indirect
	golang.org/x/net v0.0.0-20220722155237-a158d28d115b // indirect
	golang.org/x/sys v0.0.0-20220722155257-8c9f86f7a55f // indirect
	golang.org/x/text v0.4.0 // indirect
	google.golang.org/genproto v0.0.0-20200526211855-cb27e3aa2013 // indirect
)
