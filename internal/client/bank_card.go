package client

import (
	"context"
	"fmt"
	pbStorage "gophkeeper/internal/proto/storage"
)

type BankCardControl struct {
	*Client
}

func (c *BankCardControl) Save() error {
	fmt.Println("Введите ключ для сохранения данных карты:")
	c.Scanner.Scan()
	key := c.Scanner.Text()

	fmt.Println("Введите номер карты:")
	c.Scanner.Scan()
	num := c.Scanner.Text()

	fmt.Println("Введите имя и фамилию владельца карты:")
	c.Scanner.Scan()
	holder := c.Scanner.Text()

	fmt.Println("Введите cvv код карты:")
	c.Scanner.Scan()
	cvv := c.Scanner.Text()

	fmt.Println("Введите дату окончания действия карты:")
	c.Scanner.Scan()
	expireDate := c.Scanner.Text()

	fmt.Println("Введите дополнительные данные (or press Enter):")
	c.Scanner.Scan()
	md := c.Scanner.Text()

	_, err := c.StorageClient.SaveBankCard(context.Background(), &pbStorage.SaveBankCardRequest{
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

func (c *BankCardControl) Get() error {
	fmt.Println("Введите ключ для получения данных банковской карты:")
	c.Scanner.Scan()
	key := c.Scanner.Text()

	res, err := c.StorageClient.GetBankCard(context.Background(), &pbStorage.GetBankCardRequest{
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

func (c *BankCardControl) Delete() error {
	fmt.Println("Введите ключ для удаления банковской карты:")
	c.Scanner.Scan()
	key := c.Scanner.Text()

	_, err := c.StorageClient.DeleteBankCard(context.Background(), &pbStorage.DeleteBankCardRequest{
		Key: key,
	})

	if err != nil {
		fmt.Printf("Ошибка удаления банкоыской карты: %v\n", err)
		return err
	}

	fmt.Println("Банковская карта успешно удалена")

	return nil
}
