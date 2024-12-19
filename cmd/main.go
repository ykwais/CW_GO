package main

import (
	"CW_DB/api/pb"
	"CW_DB/db"
	"CW_DB/internal"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
)

func main() {
	// Подключение к базе данных
	dbPool, err := db.ConnectDB()
	if err != nil {
		log.Fatalf("Не удалось подключиться к базе данных: %v", err)
	}
	defer dbPool.Close()

	if err := db.RunMigrations(dbPool); err != nil {
		log.Fatalf("Ошибка выполнения миграций: %v", err)
	}

	// Настройка gRPC сервера
	lis, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalf("Не удалось открыть порт 8080: %v", err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterGreetingServer(grpcServer, internal.NewServer(dbPool))
	reflection.Register(grpcServer)

	log.Println("gRPC-сервер запущен на порту 8080...")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Ошибка запуска gRPC-сервера: %v", err)
	}
}
