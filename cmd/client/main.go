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

type StorageControl interface {
	save() error
	get() error
	delete() error
}

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
	fmt.Println("2 - Текст")
	fmt.Println("3 - Бинарные данные")
	fmt.Println("4 - Банковская карта")

	fmt.Println("Ожидание ввода команды (для выхода нажмите Ctrl+D):")
}

func (c *Client) storageControl(sc StorageControl) {
	for {
		fmt.Println("1 - Сохранить")
		fmt.Println("2 - Получить")
		fmt.Println("3 - Удалить")
		fmt.Println("0 - Выйти")

		c.scanner.Scan()
		line := c.scanner.Text()
		switch line {
		case "1":
			err := sc.save()
			if err == nil {
				break
			}
		case "2":
			err := sc.get()
			if err == nil {
				break
			}
		case "3":
			err := sc.delete()
			if err == nil {
				break
			}
		case "0":
			return
		}
	}
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

	passwordControl := PasswordControl{client}
	textControl := TextControl{client}
	binaryControl := BinaryControl{client}
	bankCardControl := BankCardControl{client}

	for client.scanner.Scan() {
		line := client.scanner.Text()
		switch line {
		case "1":
			client.storageControl(&passwordControl)
		case "2":
			client.storageControl(&textControl)
		case "3":
			client.storageControl(&binaryControl)
		case "4":
			client.storageControl(&bankCardControl)
		}

		shawAllCommands()
	}

	if err := client.scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "Ошибка:", err)
	}
}
