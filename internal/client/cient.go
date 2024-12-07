package client

import (
	"bufio"
	"fmt"
	pbStorage "gophkeeper/internal/proto/storage"
	pbUser "gophkeeper/internal/proto/user"
	"log"
	"os"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

type StorageControl interface {
	Save() error
	Get() error
	Delete() error
}

type Client struct {
	Conn            *grpc.ClientConn
	UserClient      pbUser.UsersClient
	StorageClient   pbStorage.StorageClient
	Scanner         *bufio.Scanner
	AuthInterceptor *InterceptorClient
}

func NewClient(address string) (*Client, error) {
	creds, err := credentials.NewClientTLSFromFile("certs/server.crt", "localhost")
	if err != nil {
		log.Fatalf("failed to load client TLS credentials: %v", err)
	}

	interceptor := NewInterceptorClient()
	conn, err := grpc.NewClient(
		address,
		grpc.WithTransportCredentials(creds),
		grpc.WithUnaryInterceptor(interceptor.Unary()),
	)
	if err != nil {
		return nil, err
	}

	userClient := pbUser.NewUsersClient(conn)
	storageClient := pbStorage.NewStorageClient(conn)

	scanner := bufio.NewScanner(os.Stdin)

	return &Client{
		Conn:            conn,
		UserClient:      userClient,
		StorageClient:   storageClient,
		Scanner:         scanner,
		AuthInterceptor: interceptor,
	}, nil
}

func (c *Client) StorageControl(sc StorageControl) {
	for {
		fmt.Println("1 - Сохранить")
		fmt.Println("2 - Получить")
		fmt.Println("3 - Удалить")
		fmt.Println("0 - Выйти")

		c.Scanner.Scan()
		line := c.Scanner.Text()
		switch line {
		case "1":
			err := sc.Save()
			if err == nil {
				break
			}
		case "2":
			err := sc.Get()
			if err == nil {
				break
			}
		case "3":
			err := sc.Delete()
			if err == nil {
				break
			}
		case "0":
			return
		}
	}
}
