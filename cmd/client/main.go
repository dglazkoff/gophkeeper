package main

import (
	"fmt"
	"gophkeeper/internal/client"
	"gophkeeper/internal/logger"
	"os"
)

var (
	BuildVersion = "N/A"
	BuildDate    = "N/A"
	BuildCommit  = "N/A"
)

func shawAllCommands() {
	fmt.Println("Возможные команды:")
	fmt.Println("1 - Логин/пароль")
	fmt.Println("2 - Текст")
	fmt.Println("3 - Бинарные данные")
	fmt.Println("4 - Банковская карта")

	fmt.Println("Ожидание ввода команды (для выхода нажмите Ctrl+D):")
}

// go run -ldflags "-X main.BuildVersion=v1.0.1 -X 'main.BuildDate=$(date +'%Y/%m/%d %H:%M:%S')'" ./cmd/client
func main() {
	err := logger.Initialize()
	if err != nil {
		panic(err)
	}

	fmt.Printf("Build version: %s\n", BuildVersion)
	fmt.Printf("Build date: %s\n", BuildDate)
	fmt.Printf("Build commit: %s\n", BuildCommit)

	c, err := client.NewClient(":3000")
	if err != nil {
		logger.Log.Error("Error while create client: ", err)
		return
	}

	defer c.Conn.Close()

	c.AuthUser()
	shawAllCommands()

	passwordControl := client.PasswordControl{Client: c}
	textControl := client.TextControl{Client: c}
	binaryControl := client.BinaryControl{Client: c}
	bankCardControl := client.BankCardControl{Client: c}

	for c.Scanner.Scan() {
		line := c.Scanner.Text()
		switch line {
		case "1":
			c.StorageControl(&passwordControl)
		case "2":
			c.StorageControl(&textControl)
		case "3":
			c.StorageControl(&binaryControl)
		case "4":
			c.StorageControl(&bankCardControl)
		}

		shawAllCommands()
	}

	if err := c.Scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "Ошибка:", err)
	}
}
