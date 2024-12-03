package main

import (
	"context"
	"fmt"
	pbStorage "gophkeeper/internal/proto/storage"
)

type TextControl struct {
	*Client
}

func (c *TextControl) save() error {
	fmt.Println("Введите ключ для сохранения текста:")
	c.scanner.Scan()
	key := c.scanner.Text()

	fmt.Println("Введите текст:")
	c.scanner.Scan()
	text := c.scanner.Text()

	fmt.Println("Введите дополнительные данные (or press Enter):")
	c.scanner.Scan()
	md := c.scanner.Text()

	_, err := c.storageClient.SaveText(context.Background(), &pbStorage.SaveTextRequest{
		Key:      key,
		Text:     text,
		Metadata: &md,
	})

	if err != nil {
		fmt.Printf("Ошибка сохранения текста: %v\n", err)
		return err
	}

	fmt.Println("Текст успешно сохранен")
	return nil
}

func (c *TextControl) get() error {
	fmt.Println("Введите ключ для получения текста:")
	c.scanner.Scan()
	key := c.scanner.Text()

	res, err := c.storageClient.GetText(context.Background(), &pbStorage.GetTextRequest{
		Key: key,
	})

	if err != nil {
		fmt.Printf("Ошибка получения текста: %v\n", err)
		return err
	}

	fmt.Println("Текст:", res.GetText())
	fmt.Println("Дополнительные данные:", res.GetMetadata())

	return nil
}

func (c *TextControl) delete() error {
	fmt.Println("Введите ключ для удаления текста:")
	c.scanner.Scan()
	key := c.scanner.Text()

	_, err := c.storageClient.DeleteText(context.Background(), &pbStorage.DeleteTextRequest{
		Key: key,
	})

	if err != nil {
		fmt.Printf("Ошибка удаления текста: %v\n", err)
		return err
	}

	fmt.Println("Текст успешно удален")

	return nil
}
