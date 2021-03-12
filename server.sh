
# cd ./cmd/server

# rm -rf server


go build -o server ./cmd/server
./server --grpc-port 8080

# go run main.go --grpc-port 8080