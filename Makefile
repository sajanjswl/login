GOCMD=go

BINARY_NAME=server
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
PATH_TO_SERVER_MAIN=./cmd/server
PARAMETES_TO_SERVER=grpc-port 8000

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


cert:
	cd cert; ./gen.sh; cd ..

# generates .pb files
generate:
	@protoc --proto_path=./api/proto/v1 --proto_path=./third_party --go_out=./pkg/api/v1 --go-grpc_out=./pkg/api/v1 user-service.proto
	@protoc --proto_path=./api/proto/v1 --proto_path=./third_party --grpc-gateway_out=logtostderr=true:./pkg/api/v1 user-service.proto
	@protoc --proto_path=./api/proto/v1 --proto_path=./third_party --swagger_out=logtostderr=true:./api/swagger/v1 user-service.proto

# to bootstrap the project
bootstrap:
	@echo "createing external network user-network for docker"
	@docker network create user-network
	@echo "pulling dependecies.."
	@go mod tidy
	@echo "deploying postgress and PgAdmin"
	@docker-compose -f ./postgres/postgres.yaml up -d
	@ sleep 10s
	@echo "restoring sql backup..."
	@docker exec -i postgres psql -U postgres -d postgres < backup.sql
	
# destroys  the postgres containers
destroy:
	@echo "removing postgress..."
	@docker-compose -f ./postgres/postgres.yaml down

# creates sqldump
backup:
	@echo "creating sql dump"
	@docker exec postgres pg_dump -U postgres postgres > backup.sql

# deploys server in a docker container
deploy1:
	@echo "building binary"
	@env GOOS=linux GOARCH=arm go build -o bin/server_linux ./cmd/server
	@echo "building docker image.."
	@docker build . -t user:latest1001
	@echo "deploying..."
	@docker-compose -f docker-compose.yaml up

# destorys the user-service container
destroyDeploy:
	@echo "removing container"
	@docker-compose -f docker-compose.yaml down


########################################################
################### TERRAFORM ##########################
########################################################

tf-init:
	docker-compose -f deploy/docker-compose.yml run --rm terraform init

tf-fmt:
	docker-compose -f deploy/docker-compose.yml run --rm terraform fmt

tf-validate:
	docker-compose -f deploy/docker-compose.yml run --rm terraform validate

tf-plan:
	docker-compose -f deploy/docker-compose.yml run --rm terraform plan

tf-apply:
	docker-compose -f deploy/docker-compose.yml run --rm terraform apply

tf-destroy:
	docker-compose -f deploy/docker-compose.yml run --rm terraform destroy 

tf-workspace-dev:
	docker-compose -f deploy/docker-compose.yml run --rm terraform workspace select dev

tf-workspace-staging:
	docker-compose -f deploy/docker-compose.yml run --rm terraform workspace select staging


tf-workspace-prod:
	docker-compose -f deploy/docker-compose.yml run --rm terraform workspace select prod

.PHONY: test
test:
	docker-compose run --rm app sh -c "python manage.py wait_for_db && python manage.py test && flake8"

tf-workspace-list:
	docker-compose -f deploy/docker-compose.yml run --rm terraform workspace list

tf-workspace-new-dev:
	docker-compose -f deploy/docker-compose.yml run --rm terraform workspace new dev


