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
 
 
Testing menggunakan Postman
1. run: go run server/main.go
2. buka postman (versi 9.7 keatas)
3. File > New > gRPC Request
4. masukan URL: localhost:9090
5. import file .proto (protobuf definition) dari direktrori kendaraan di project ini
6. pilih service/method yg tersedia: GetAllMobil, GetMObil, CreateMobil, UpdateMObil, DeleteMobil
7. klik: Generate Example Message, untuk mengirim message sesuai service yg kita pilih
8. klik: Envoke
