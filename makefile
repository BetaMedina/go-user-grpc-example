generate:
	@protoc --go_out=plugins=grpc:. ./proto/user.proto

run:
	@echo "---- Running Server ----"
	@go run cmd/main.go

run_client:
	@echo "---- Running Client ----"
	@go run ./client/client.go