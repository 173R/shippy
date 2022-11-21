build:
	go mod tidy
	docker build -t service-consignment .
run:
	docker run -p 50051:50051 service-consignment