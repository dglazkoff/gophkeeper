package client

import (
	"context"
	"fmt"
	pbStorage "gophkeeper/internal/proto/storage"
)

type TextControl struct {
	*Client
}

func (c *TextControl) Save() error {
	fmt.Println("Введите ключ для сохранения текста:")
	c.Scanner.Scan()
	key := c.Scanner.Text()

	fmt.Println("Введите текст:")
	c.Scanner.Scan()
	text := c.Scanner.Text()

	fmt.Println("Введите дополнительные данные (or press Enter):")
	c.Scanner.Scan()
	md := c.Scanner.Text()

	_, err := c.StorageClient.SaveText(context.Background(), &pbStorage.SaveTextRequest{
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

func (c *TextControl) Get() error {
	fmt.Println("Введите ключ для получения текста:")
	c.Scanner.Scan()
	key := c.Scanner.Text()

	res, err := c.StorageClient.GetText(context.Background(), &pbStorage.GetTextRequest{
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

func (c *TextControl) Delete() error {
	fmt.Println("Введите ключ для удаления текста:")
	c.Scanner.Scan()
	key := c.Scanner.Text()

	_, err := c.StorageClient.DeleteText(context.Background(), &pbStorage.DeleteTextRequest{
		Key: key,
	})

	if err != nil {
		fmt.Printf("Ошибка удаления текста: %v\n", err)
		return err
	}

	fmt.Println("Текст успешно удален")

	return nil
}
