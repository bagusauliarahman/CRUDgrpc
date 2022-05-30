package main

import (
	"context"
	"log"

	"net"

	pb "belajar.grpc/crud/kendaraan"
	"google.golang.org/grpc"
)

const (
	port = ":9090"
)

var mobils []*pb.InfoMobil

type mobilServer struct {
	pb.UnimplementedMobilServer
}

func main() {
	initMobil()
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)

	}
	s := grpc.NewServer()

	pb.RegisterMobilServer(s, &mobilServer{})

	log.Printf("server is listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

func initMobil() {
	mobil1 := &pb.InfoMobil{Id: "B1000AR", Warna: "hijau", Tahun: 2005, Manufaktur: &pb.Manufaktur{Merek: "Toyota", Jenis: "limo"}}
	mobil2 := &pb.InfoMobil{Id: "B2222WA", Warna: "hitam", Tahun: 2009, Manufaktur: &pb.Manufaktur{Merek: "Mercedes Benz", Jenis: "C class"}}

	mobils = append(mobils, mobil1)
	mobils = append(mobils, mobil2)
}

func (s *mobilServer) GetAllMobil(in *pb.Empty, stream pb.Mobil_GetAllMobilServer) error {
	log.Printf("REceived: %v", in)

	for _, mobil := range mobils {
		if err := stream.Send(mobil); err != nil {
			return err
		}

	}
	return nil
}

func (s *mobilServer) GetMobil(ctx context.Context, in *pb.Id) (*pb.InfoMobil, error) {
	log.Printf("received: %v", in)

	res := &pb.InfoMobil{}

	for _, mobil := range mobils {
		if mobil.GetId() == in.GetValue() {
			res = mobil
			break
		}

	}
	return res, nil
}

func (s *mobilServer) CreateMobil(ctx context.Context, in *pb.InfoMobil) (*pb.Id, error) {
	log.Printf("Received: %v", in)
	res := pb.Id{}

	mobils = append(mobils, in)
	return &res, nil
}

func (s *mobilServer) UpdateMobil(ctx context.Context, in *pb.InfoMobil) (*pb.Status, error) {
	log.Printf("Received: %v", in)

	res := pb.Status{}
	for index, mobil := range mobils {
		if mobil.GetId() == in.GetId() {
			mobils = append(mobils[:index], mobils[index+1:]...)
			in.Id = mobil.GetId()
			mobils = append(mobils, in)
			res.Value = 1
			break
		}
	}
	return &res, nil
}

func (s *mobilServer) DeleteMobil(ctx context.Context, in *pb.Id) (*pb.Status, error) {
	log.Printf("received: %v", in)

	res := pb.Status{}
	for index, mobil := range mobils {
		if mobil.GetId() == in.GetValue() {
			mobils = append(mobils[:index], mobils[index+1:]...)
			res.Value = 1
			break

		}
	}
	return &res, nil
}
