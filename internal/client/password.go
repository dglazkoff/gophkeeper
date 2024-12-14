package client

import (
	"context"
	"fmt"
	pbStorage "gophkeeper/internal/proto/storage"
)

type PasswordControl struct {
	*Client
}

func (c *PasswordControl) Save() error {
	fmt.Println("Введите ключ для сохранения пары логин/пароль:")
	c.Scanner.Scan()
	key := c.Scanner.Text()

	fmt.Println("Введите логин:")
	c.Scanner.Scan()
	login := c.Scanner.Text()

	fmt.Println("Введите пароль:")
	c.Scanner.Scan()
	password := c.Scanner.Text()

	fmt.Println("Введите дополнительные данные (or press Enter):")
	c.Scanner.Scan()
	md := c.Scanner.Text()

	_, err := c.StorageClient.SavePassword(context.Background(), &pbStorage.SavePasswordRequest{
		Key:      key,
		Login:    login,
		Password: password,
		Metadata: &md,
	})

	if err != nil {
		fmt.Printf("Ошибка сохранения пароля: %v\n", err)
		return err
	}

	fmt.Println("Пара логин/пароль успешно сохранена")
	return nil
}

func (c *PasswordControl) Get() error {
	fmt.Println("Введите ключ для получения пары логин/пароль:")
	c.Scanner.Scan()
	key := c.Scanner.Text()

	res, err := c.StorageClient.GetPassword(context.Background(), &pbStorage.GetPasswordRequest{
		Key: key,
	})

	if err != nil {
		fmt.Printf("Ошибка получения пароля: %v\n", err)
		return err
	}

	fmt.Println("Логин:", res.GetLogin())
	fmt.Println("Пароль:", res.GetPassword())
	fmt.Println("Дополнительные данные:", res.GetMetadata())

	return nil
}

func (c *PasswordControl) Delete() error {
	fmt.Println("Введите ключ для удаления пары логин/пароль:")
	c.Scanner.Scan()
	key := c.Scanner.Text()

	_, err := c.StorageClient.DeletePassword(context.Background(), &pbStorage.DeletePasswordRequest{
		Key: key,
	})

	if err != nil {
		fmt.Printf("Ошибка удаления пароля: %v\n", err)
		return err
	}

	fmt.Println("Пара логин/пароль успешно удалена")

	return nil
}
