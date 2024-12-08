package main

import (
	"database/sql"
	"errors"
	"fmt"
	"gophkeeper/internal/api"
	"gophkeeper/internal/auth"
	"gophkeeper/internal/db"
	"gophkeeper/internal/logger"
	serverservice "gophkeeper/internal/service/server"
	storageservice "gophkeeper/internal/service/storage"
	userservice "gophkeeper/internal/service/user"
	"log"
	"net"
	"net/http"
	"os"

	pbServer "gophkeeper/internal/proto/server"
	pbStorage "gophkeeper/internal/proto/storage"
	pbUser "gophkeeper/internal/proto/user"

	_ "github.com/jackc/pgx/v5/stdlib"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

var (
	BuildVersion = "N/A"
	BuildDate    = "N/A"
	BuildCommit  = "N/A"
)

// go run -ldflags "-X main.BuildVersion=v1.0.1 -X 'main.BuildDate=$(date +'%Y/%m/%d %H:%M:%S')'" ./cmd/server
func main() {
	err := logger.Initialize()

	if err != nil {
		panic(err)
	}

	fmt.Printf("Build version: %s\n", BuildVersion)
	fmt.Printf("Build date: %s\n", BuildDate)
	fmt.Printf("Build commit: %s\n", BuildCommit)

	pgDB, err := sql.Open("pgx", os.Getenv("DATABASE_URL")) //postgres://postgres:12345678@localhost:5432/gophkeeper

	if err != nil {
		logger.Log.Debug("Error on open db", err)
		panic(err)
	}
	defer pgDB.Close()

	s3DB, err := db.NewS3()

	if err != nil {
		logger.Log.Debug("Error on open db", err)
		panic(err)
	}

	store := db.New(pgDB)
	err = store.Bootstrap()

	if err != nil {
		logger.Log.Debug("Error on bootstrap db ", err)
		return
	}

	us := userservice.NewUserService(store)
	storageService := storageservice.NewStorageService(store, s3DB)
	serverService := serverservice.NewServerService(store)

	creds, err := credentials.NewServerTLSFromFile("./certs/server.crt", "./certs/server.key")
	if err != nil {
		log.Fatalf("failed to load server TLS credentials: %v", err)
	}

	listen, err := net.Listen("tcp", ":3000")
	if err != nil {
		panic(err)
	}

	interceptor := auth.NewInterceptorServer()
	grpcServer := grpc.NewServer(grpc.Creds(creds), grpc.UnaryInterceptor(interceptor.Unary()))
	pbUser.RegisterUsersServer(grpcServer, api.NewUserServer(us))
	pbStorage.RegisterStorageServer(grpcServer, api.NewStorageServer(storageService))
	pbServer.RegisterServerServer(grpcServer, api.NewServer(serverService))

	if err := grpcServer.Serve(listen); err != nil && !errors.Is(err, http.ErrServerClosed) {
		panic(err)
	}
}
