package internal

import (
	"CW_DB/api/pb"
	"context"
	"fmt"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"log"
)

type Server struct {
	pb.UnimplementedGreetingServer
	dbPool *pgxpool.Pool
}

func NewServer(dbPool *pgxpool.Pool) *Server {
	return &Server{dbPool: dbPool}
}

func (s *Server) Login(ctx context.Context, req *pb.HelloRequest) (*pb.HelloResponse, error) {
	log.Printf("Получен запрос: логин=%s, пароль=%s\n", req.Login, req.Password)

	query := "SELECT register_client(@user_name, @pass_hash, @email_addr, @real_name, @passport_number_serial, @phone_number, @driving_experience_)"
	args := pgx.NamedArgs{
		"user_name":              req.Login,
		"pass_hash":              req.Password,
		"email_addr":             "example@domain.com",
		"real_name":              "Test User",
		"passport_number_serial": "123456789",
		"phone_number":           "+1234567890",
		"driving_experience_":    5,
	}

	_, err := s.dbPool.Exec(ctx, query, args)
	if err != nil {
		log.Printf("Ошибка выполнения хранимой функции: %v", err)
		return nil, fmt.Errorf("не удалось зарегистрировать клиента: %v", err)
	}

	message := fmt.Sprintf("Привет, %s! Вы успешно зарегистрированы.", req.Login)
	return &pb.HelloResponse{Message: message}, nil
}
