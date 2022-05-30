# CRUDgrpc
simple gRPC CRUD API project

├── kendaraan
│   └── kendaraan.proto
│
├── server
│   └── main.go
│
├── client
│   └── main.go

langkah2

 go mod init belajar.grpc/crud
 go get google.golang.org/grpc
 protoc --go_out=paths=source_relative:. --go-grpc_out=paths=source_relative:. kendaraan/kendaraan.proto
 
 
