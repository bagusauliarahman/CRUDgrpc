package main

import (
	"context"
	"io"
	"log"
	"time"

	pb "belajar.grpc/crud/kendaraan"
	"google.golang.org/grpc"
)

const (
	address = "localhost:9090"
)

func main() {
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	client := pb.NewMobilClient(conn)

	runGetAllMobil(client)

	//
	//
	//
	//

}

func runGetAllMobil(client pb.MobilClient) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)

	defer cancel()

	req := &pb.Empty{}
	stream, err := client.GetAllMobil(ctx, req)
	if err != nil {
		log.Fatalf("%v.GetAllMobil(_) = _, %v", client, err)

	}
	for {
		row, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("%v.GetAllMobil(_) = _, %v", client, err)
		}

		log.Printf("Info Armada: %v", row)
	}
}

func runGetMobil(client pb.MobilClient, mobilid string) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	req := &pb.Id{Value: mobilid}
	res, err := client.GetMobil(ctx, req)
	if err != nil {
		log.Fatalf("%v.GetMobil(_) = _, %v", client, err)
	}
	log.Printf("Info Armada: %v", res)

}

func runCreateMobil(client pb.MobilClient, warna string, tahun int32, merek string, jenis string) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	req := &pb.InfoMobil{Warna: warna, Tahun: tahun, Manufaktur: &pb.Manufaktur{Merek: merek, Jenis: jenis}}

	res, err := client.CreateMobil(ctx, req)
	if err != nil {
		log.Fatalf("%v.CreateMobil(_) =  _, %v", client, err)
	}
	if res.GetValue() != "" {
		log.Printf("CreateMobil Id: %v", res)
	} else {
		log.Printf("CreateMobil Failed")
	}
}

func runUpdateMobil(client pb.MobilClient, mobilid string, warna string, tahun int32, merek string, jenis string) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	req := &pb.InfoMobil{Id: mobilid, Warna: warna, Tahun: tahun, Manufaktur: &pb.Manufaktur{Merek: merek, Jenis: jenis}}
	res, err := client.UpdateMobil(ctx, req)
	if err != nil {
		log.Fatalf("%v.UpdateMobil(_) = _, %v", client, err)
	}
	if int(res.GetValue()) == 1 {
		log.Printf("Update Armada Sukses")
	} else {
		log.Printf("Update Armada Failed")
	}
}

func runDeleteMobil(client pb.MobilClient, mobilid string) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	req := &pb.Id{Value: mobilid}
	res, err := client.DeleteMobil(ctx, req)
	if err != nil {
		log.Fatalf("%v.DeleteMobil(_) =  _, %v", client, err)
	}
	if int(res.GetValue()) == 1 {
		log.Printf("Delete Armada Sukses")
	} else {
		log.Printf(" Delete Armada Failed")
	}
}
