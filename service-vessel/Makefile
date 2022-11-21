build:
	go mod tidy
	docker build -t service-vessel .
run:
	docker run -p 50052:50052 service-vessel