
all:	 run

generate:
	@buf generate

run:
	@echo "starting the server..."
	@go run main.go

mysql:
	@echo "msql and adminer..."
	@docker-compose -f mysql.yaml up

mysql-destroy:
	@echo "destroying mysql"
	@docker-compose -f mysql.yaml down -v

grpc: 
	@echo "starting grpc client..."
	@go run auth-client/grpc/main.go

rest: 
	@echo "starting rest client..."
	@go run auth-client/rest/main.go


