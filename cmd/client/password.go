package main

import (
	"context"
	"fmt"
	pbStorage "gophkeeper/internal/proto/storage"
)

type PasswordControl struct {
	*Client
}

func (c *PasswordControl) save() error {
	fmt.Println("Введите ключ для сохранения пары логин/пароль:")
	c.scanner.Scan()
	key := c.scanner.Text()

	fmt.Println("Введите логин:")
	c.scanner.Scan()
	login := c.scanner.Text()

	fmt.Println("Введите пароль:")
	c.scanner.Scan()
	password := c.scanner.Text()

	fmt.Println("Введите дополнительные данные (or press Enter):")
	c.scanner.Scan()
	md := c.scanner.Text()

	_, err := c.storageClient.SavePassword(context.Background(), &pbStorage.SavePasswordRequest{
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

func (c *PasswordControl) get() error {
	fmt.Println("Введите ключ для получения пары логин/пароль:")
	c.scanner.Scan()
	key := c.scanner.Text()

	res, err := c.storageClient.GetPassword(context.Background(), &pbStorage.GetPasswordRequest{
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

func (c *PasswordControl) delete() error {
	fmt.Println("Введите ключ для удаления пары логин/пароль:")
	c.scanner.Scan()
	key := c.scanner.Text()

	_, err := c.storageClient.DeletePassword(context.Background(), &pbStorage.DeletePasswordRequest{
		Key: key,
	})

	if err != nil {
		fmt.Printf("Ошибка удаления пароля: %v\n", err)
		return err
	}

	fmt.Println("Пара логин/пароль успешно удалена")

	return nil
}
