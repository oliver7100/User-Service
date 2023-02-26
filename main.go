package main

import (
	"log"
	"net"

	"google.golang.org/grpc"

	"github.com/oliver7100/user-service/database"
	"github.com/oliver7100/user-service/proto"
)

func main() {
	dbConnection, err := database.NewDatabaseConnection(
		"services:AVNS_zWWallm_soPdGTwcPQJ@tcp(db-mysql-fmf-do-user-7517862-0.b.db.ondigitalocean.com:25060)/db_user_service?charset=utf8mb4&parseTime=True&loc=Local",
	)

	if err != nil {
		log.Fatalf("Cant connect to db")
	}

	listener, err := net.Listen("tcp", ":9000")

	if err != nil {
		log.Fatalf("failed to listen")
	}

	s := grpc.NewServer()

	proto.RegisterUserServiceServer(
		s,
		proto.CreateNewService(dbConnection),
	)

	if err := s.Serve(listener); err != nil {
		panic(err)
	}
}
