package main

import (
	"context"
	"fmt"
	pbStorage "gophkeeper/internal/proto/storage"
	"os"
	"path/filepath"
)

type BinaryControl struct {
	*Client
}

func (c *BinaryControl) save() error {
	fmt.Println("Введите путь для получения бинарных данных:")
	c.scanner.Scan()
	path := c.scanner.Text()

	data, err := os.ReadFile(path)

	if err != nil {
		fmt.Printf("Ошибка чтения файла: %s\n", err)
		return err
	}

	fmt.Println("Введите дополнительные данные (or press Enter):")
	c.scanner.Scan()
	md := c.scanner.Text()

	_, err = c.storageClient.SaveBinary(context.Background(), &pbStorage.SaveBinaryRequest{
		Key:      filepath.Base(path),
		Value:    data,
		Metadata: &md,
	})

	if err != nil {
		fmt.Printf("Ошибка сохранения бинарных данных: %v\n", err)
		return err
	}

	fmt.Printf("Бинарные данные успешно сохранены по ключу %s\n", filepath.Base(path))
	return nil
}

func (c *BinaryControl) get() error {
	fmt.Println("Введите ключ для получения данных:")
	c.scanner.Scan()
	key := c.scanner.Text()

	fmt.Println("Введите путь для сохранения бинарных данных:")
	c.scanner.Scan()
	path := c.scanner.Text()

	file, err := os.Create(path)
	defer file.Close()

	if err != nil {
		fmt.Printf("Ошибка создания файла: %s\n", err)
		return err
	}

	res, err := c.storageClient.GetBinary(context.Background(), &pbStorage.GetBinaryRequest{
		Key: key,
	})

	if err != nil {
		fmt.Printf("Ошибка получения бинарных данныъ: %v\n", err)
		return err
	}

	_, err = file.Write(res.GetValue())

	if err != nil {
		fmt.Printf("Ошибка записи файла: %s\n", err)
		return err
	}

	fmt.Println("Данные успешно записаны")
	fmt.Println("Дополнительные данные:", res.GetMetadata())

	return nil
}

func (c *BinaryControl) delete() error {
	fmt.Println("Введите ключ для удаления бинарных данных:")
	c.scanner.Scan()
	key := c.scanner.Text()

	_, err := c.storageClient.DeleteBinary(context.Background(), &pbStorage.DeleteBinaryRequest{
		Key: key,
	})

	if err != nil {
		fmt.Printf("Ошибка удаления данных: %v\n", err)
		return err
	}

	fmt.Println("Бинарные данные успешно удалены")

	return nil
}
