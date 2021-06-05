 protoc --go_out=plugins=grpc:. **/**.proto
 
 // windows
 protoc --go_out=. --go-grpc_opt=require_unimplemented_servers=false --go-grpc_out=.  hello/hello.proto

