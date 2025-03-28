package main

import (
	"fmt"
	"log"
	"net"
	"time"

	pb "github.com/agrotention/pb-user"
	"github.com/agrotention/service-user/cfg"
	"github.com/agrotention/service-user/repo"
	"github.com/agrotention/service-user/service"
	"google.golang.org/grpc"
)

func main() {
	db := repo.NewDB()
	repo := repo.NewRepoUser(db)
	service := service.NewService(repo) // FIXME: inject with repo
	server := grpc.NewServer(
		grpc.ConnectionTimeout(time.Second*30),
		grpc.InitialWindowSize(1024),
	)
	pb.RegisterServiceUserServer(server, service)
	addr := fmt.Sprintf("%s:%s", cfg.SERVICE_USER_HOST, cfg.SERVICE_USER_PORT)
	ln, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalf("listener error: %v", err)
	}

	log.Printf("Server listening on %s", addr)
	err = server.Serve(ln)
	if err != nil {
		log.Fatalf("server serve error: %v", err)
	}
}
