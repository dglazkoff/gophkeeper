package main

import (
	"database/sql"
	"errors"
	"fmt"
	"gophkeeper/internal/api"
	"gophkeeper/internal/auth"
	"gophkeeper/internal/db"
	"gophkeeper/internal/logger"
	storageservice "gophkeeper/internal/service/storage"
	userservice "gophkeeper/internal/service/user"
	"net"
	"net/http"

	pbStorage "gophkeeper/internal/proto/storage"
	pbUser "gophkeeper/internal/proto/user"

	_ "github.com/jackc/pgx/v5/stdlib"
	"google.golang.org/grpc"
)

func main() {
	err := logger.Initialize()

	if err != nil {
		panic(err)
	}

	pgDB, err := sql.Open("pgx", "postgres://postgres:12345678@localhost:5432/gophkeeper")

	if err != nil {
		fmt.Println("Error on open db", err)
		panic(err)
	}
	defer pgDB.Close()

	store := db.New(pgDB)
	err = store.Bootstrap()

	if err != nil {
		logger.Log.Debug("Error on bootstrap db ", err)
		return
	}

	us := userservice.NewUserService(store)
	ss := storageservice.NewStorageService(store)

	listen, err := net.Listen("tcp", ":3000")
	if err != nil {
		panic(err)
	}

	interceptor := auth.NewInterceptorServer()
	grpcServer := grpc.NewServer(grpc.UnaryInterceptor(interceptor.Unary()))
	pbUser.RegisterUsersServer(grpcServer, api.NewUserServer(us))
	pbStorage.RegisterStorageServer(grpcServer, api.NewStorageServer(ss))

	if err := grpcServer.Serve(listen); err != nil && !errors.Is(err, http.ErrServerClosed) {
		panic(err)
	}
}
