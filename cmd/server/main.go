package main

import (
	"fmt"
	db "github.com/TouchGlass/users-service/internal/database"
	transportgrpc "github.com/TouchGlass/users-service/internal/transport/grpc"
	"github.com/TouchGlass/users-service/internal/user"
	"log"
)

func main() {
	database, err := db.InitDB()
	if err != nil {
		log.Fatalf("Error connecting to database: %v", err)
	}

	uRepo := user.NewUserRepository(database)
	uService := user.NewUserService(uRepo)

	fmt.Println("gRPC сервер запущен на :50051")

	if err := transportgrpc.RunGRPC(uService); err != nil {
		log.Fatalf("Error starting GRPC server: %v", err)
	}
}
