package main

import (
	"bufio"
	"fmt"
	"gophkeeper/internal/auth"
	"gophkeeper/internal/logger"
	"os"

	pbStorage "gophkeeper/internal/proto/storage"
	pbUser "gophkeeper/internal/proto/user"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type Client struct {
	conn            *grpc.ClientConn
	userClient      pbUser.UsersClient
	storageClient   pbStorage.StorageClient
	scanner         *bufio.Scanner
	authInterceptor *auth.InterceptorClient
}

func NewClient(address string) (*Client, error) {
	interceptor := auth.NewInterceptorClient()
	conn, err := grpc.NewClient(
		address,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithUnaryInterceptor(interceptor.Unary()),
	)
	if err != nil {
		return nil, err
	}

	userClient := pbUser.NewUsersClient(conn)
	storageClient := pbStorage.NewStorageClient(conn)

	scanner := bufio.NewScanner(os.Stdin)

	return &Client{
		conn:            conn,
		userClient:      userClient,
		storageClient:   storageClient,
		scanner:         scanner,
		authInterceptor: interceptor,
	}, nil
}

func shawAllCommands() {
	fmt.Println("Возможные команды:")
	fmt.Println("1 - Логин/пароль")

	fmt.Println("Ожидание ввода команды (для выхода нажмите Ctrl+D):")
}

func main() {
	err := logger.Initialize()
	if err != nil {
		panic(err)
	}

	client, err := NewClient(":3000")
	if err != nil {
		logger.Log.Error("Error while create client: ", err)
		return
	}

	defer client.conn.Close()

	client.authUser()
	shawAllCommands()

	for client.scanner.Scan() {
		line := client.scanner.Text()
		switch line {
		case "1":
			client.keepPassword()
		}

		shawAllCommands()
	}

	if err := client.scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "Ошибка:", err)
	}
}
