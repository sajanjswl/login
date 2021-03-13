# user-service

user-service currently consists of 7 RPC. The Api/proto definations is present in api/proto/v1 if you update the proto defination don't forget to run the proto-gen.sh script present in ./script directory at the root of the project. This script will generate the required pb.go, grpc.pb.go and pb.gw.go files in the path /pkg/api/v1.

We have grpc-rest-gateway implemented i.e all the rpc has their corresponding HTTP/REST enpoint.

####Steps to bootstrap the projcet
*   `docker-compose -f ./docker/postgres.yaml up -d`  : It start two container one for postgres and the other for PgAdmin.
* Wait for the containers to come up and then run `docker exec -i postgres psql -U postgres -d postgres < backup.sql` 
* `go mod tidy` : to download the go dependencies
* `./server.sh` : starts the server
* `./client.sh`: runs the client call for  the rpc that is not commented in cmd/client-grpc/main.go file
* We also have a rest-client: cmd/client-rest
* For PgAdmin setup: https://medium.com/@sjnjaiswal/postgres-in-a-docker-container-simplified-b7e97ef30cfb
  




