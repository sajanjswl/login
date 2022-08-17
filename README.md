#auth-service

* The auth service is build using `golang` [gRPC](https://grpc.io/),[gRPC-gateway](https://github.com/grpc-ecosystem/grpc-gateway) 
* Proto files are managed by [bufbuild](https://docs.buf.build/introduction)

* This auth service starts two server one runs `gRPC Protocol` on Port `8000` and the other runs `Rest Protocol` on Port `9000` which proxies the trafic to `gRPC Service` 


###Step to run the service
1) `make mysql` will brotstap mysql and adminer container to provide mysql service to the application
2) `make run` will start the auth service `gRPC on Port 8000` and `Rest on Port 9000`
3) `make grpc` will make a `gRPC` client call to the server 
4) `make rest` will make a `Rest` client call to the server
5) For login with google : [http://localhost:9000/v1/login/ui/](http://localhost:9000/v1/login/ui/)
6) `make generate` will generate update changes made to proto files.
  




