package main

import (
	"film_grpc/serverGo/initializers"
	"google.golang.org/grpc"
	"log"
	"net"
	"os"

	pb "film_grpc/proto"
)

func init() {
	initializers.LoadEnvVar()
	initializers.ConnectToDB()
}

type Server struct {
	pb.UnimplementedFilmGenreServiceServer
}

func main() {
	port := os.Getenv("PORT_SERVER_GO")
	listener, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("Faild to start sever %v", err)
	}
	grpcServer := grpc.NewServer()
	pb.RegisterFilmGenreServiceServer(grpcServer, &Server{})
	log.Printf("Server started at %v", listener.Addr())
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("Failed to start %v", err)
	}
}
