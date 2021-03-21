GOCMD=go

BINARY_NAME=server
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
PATH_TO_SERVER_MAIN=./cmd/server
PARAMETES_TO_SERVER=grpc-port 8456

all:	clean build run

# clean the binary
clean: 
	@echo "cleaning the environment..."
	@$(GOCLEAN)
	@rm -f ./bin/$(BINARY_NAME)

# but the server binary
build: 
	@echo "building server..." 
	@$(GOBUILD) -v -o ./bin/$(BINARY_NAME) $(PATH_TO_SERVER_MAIN)
	
# runs the server	
run:
	@echo "starting the server..."
	@./bin/$(BINARY_NAME) --$(PARAMETES_TO_SERVER)

# generates .pb files
generate:
	@protoc --proto_path=./api/proto/v1 --proto_path=./third_party --go_out=./pkg/api/v1 --go-grpc_out=./pkg/api/v1 user-service.proto
	@protoc --proto_path=./api/proto/v1 --proto_path=./third_party --grpc-gateway_out=logtostderr=true:./pkg/api/v1 user-service.proto
	@protoc --proto_path=./api/proto/v1 --proto_path=./third_party --swagger_out=logtostderr=true:./api/swagger/v1 user-service.proto

# to bootstrap the project
bootstrap:
	@echo "pulling dependecies.."
	@go mod tidy
	@echo "deploying postgress and PgAdmin"
	@docker-compose -f ./docker/postgres.yaml up -d
	@ sleep 10s
	@echo "restoring sql backup..."
	@docker exec -i postgres psql -U postgres -d postgres < backup.sql
	
# destroys  the postgres containers
destroy:
	@echo "removing postgress..."
	@docker-compose -f ./docker/postgres.yaml down

# creates sqldump
backup:
	@echo "creating sql dump"
	@docker exec postgres pg_dump -U postgres postgres > backup.sql

# deploys server in a docker container
deploy:
	@echo "building binary"
	@env GOOS=linux GOARCH=arm go build -o bin/server_linux ./cmd/server
	@echo "building docker image.."
	@docker build . -t user:v1
	@echo "deploying..."
	@docker-compose -f docker-compose.yaml up

# destorys the user-service container
destroyDeploy:
	@echo "removing container"
	@docker-compose -f docker-compose.yaml down