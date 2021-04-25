## grpc-client


# go build -o bin/grpc-client ./cmd/grpc-client
# ./bin/grpc-client --server=localhost:8080


## rest-client

# go build -o bin/rest-client ./cmd/rest-client
# ./bin/rest-client --server=http://localhost:8085

# For running client with Nginx Proxy user Port 8080 for GRPC and PORT 8085 for REST



# for aws-deployment grpc

# go build -o bin/grpc-client ./cmd/grpc-client
# ./bin/grpc-client --server=http://user-service-default-main-1728128331.us-east-1.elb.amazonaws.com:80


## for aws-deployment rest

go build -o bin/rest-client ./cmd/rest-client
./bin/rest-client --server=http://user-service-default-main-1728128331.us-east-1.elb.amazonaws.com:80
