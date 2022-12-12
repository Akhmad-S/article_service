package main

import (
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	"github.com/uacademy/blogpost/article_service/config"
	"github.com/uacademy/blogpost/article_service/proto-gen/blogpost"
	"github.com/uacademy/blogpost/article_service/services/article"
	"github.com/uacademy/blogpost/article_service/services/author"
	"github.com/uacademy/blogpost/article_service/storage"
	"github.com/uacademy/blogpost/article_service/storage/postgres"
)

func main() {
	cfg := config.Load()

	psqlConString := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		cfg.PostgresHost,
		cfg.PostgresPort,
		cfg.PostgresUser,
		cfg.PostgresPassword,
		cfg.PostgresDatabase,
	)

	var stg storage.StorageI
	stg, err := postgres.InitDb(psqlConString)
	if err != nil {
		panic(err)
	}

	println("gRPC server tutorial in Go")

	listener, err := net.Listen("tcp", ":9001")
	if err != nil {
		panic(err)
	}

	s := grpc.NewServer()
	blogpost.RegisterArticleServiceServer(s, article.NewArticleService(stg))
	blogpost.RegisterAuthorServiceServer(s, author.NewAuthorService(stg))
	reflection.Register(s)
	if err := s.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
