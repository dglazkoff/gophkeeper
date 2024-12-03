package main

import (
	"context"
	"fmt"
	pbStorage "gophkeeper/internal/proto/storage"
)

type BankCardControl struct {
	*Client
}

func (c *BankCardControl) save() error {
	fmt.Println("Введите ключ для сохранения данных карты:")
	c.scanner.Scan()
	key := c.scanner.Text()

	fmt.Println("Введите номер карты:")
	c.scanner.Scan()
	num := c.scanner.Text()

	fmt.Println("Введите имя и фамилию владельца карты:")
	c.scanner.Scan()
	holder := c.scanner.Text()

	fmt.Println("Введите cvv код карты:")
	c.scanner.Scan()
	cvv := c.scanner.Text()

	fmt.Println("Введите дату окончания действия карты:")
	c.scanner.Scan()
	expireDate := c.scanner.Text()

	fmt.Println("Введите дополнительные данные (or press Enter):")
	c.scanner.Scan()
	md := c.scanner.Text()

	_, err := c.storageClient.SaveBankCard(context.Background(), &pbStorage.SaveBankCardRequest{
		Key:            key,
		Number:         num,
		Holder:         holder,
		Cvv:            cvv,
		ExpirationDate: expireDate,
		Metadata:       &md,
	})

	if err != nil {
		fmt.Printf("Ошибка сохранения банковской карты: %v\n", err)
		return err
	}

	fmt.Println("Банковская карта успешно сохранена")
	return nil
}

func (c *BankCardControl) get() error {
	fmt.Println("Введите ключ для получения данных банковской карты:")
	c.scanner.Scan()
	key := c.scanner.Text()

	res, err := c.storageClient.GetBankCard(context.Background(), &pbStorage.GetBankCardRequest{
		Key: key,
	})

	if err != nil {
		fmt.Printf("Ошибка получения банковской карты: %v\n", err)
		return err
	}

	fmt.Println("Номер карты:", res.GetNumber())
	fmt.Println("Имя и фамилия владельца карты:", res.GetHolder())
	fmt.Println("CVV код карты:", res.GetCvv())
	fmt.Println("Дата окончания действия карты:", res.GetExpirationDate())
	fmt.Println("Дополнительные данные:", res.GetMetadata())

	return nil
}

func (c *BankCardControl) delete() error {
	fmt.Println("Введите ключ для удаления банковской карты:")
	c.scanner.Scan()
	key := c.scanner.Text()

	_, err := c.storageClient.DeleteBankCard(context.Background(), &pbStorage.DeleteBankCardRequest{
		Key: key,
	})

	if err != nil {
		fmt.Printf("Ошибка удаления банкоыской карты: %v\n", err)
		return err
	}

	fmt.Println("Банковская карта успешно удалена")

	return nil
}
