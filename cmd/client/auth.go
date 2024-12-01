package main

import (
	"context"
	"fmt"
	pbUser "gophkeeper/internal/proto/user"
)

func (c *Client) login() error {
	fmt.Println("Введите логин:")
	c.scanner.Scan()
	login := c.scanner.Text()

	fmt.Println("Введите пароль:")
	c.scanner.Scan()
	password := c.scanner.Text()

	res, err := c.userClient.LoginUser(context.Background(), &pbUser.LoginUserRequest{
		Login:    login,
		Password: password,
	})

	if err != nil {
		fmt.Printf("Ошибка авторизации: %v\n", err)
		return err
	}

	token := res.GetAccessToken()
	c.authInterceptor.AccessToken = token

	return nil
}

func (c *Client) register() error {
	fmt.Println("Введите логин:")
	c.scanner.Scan()
	login := c.scanner.Text()

	fmt.Println("Введите пароль:")
	c.scanner.Scan()
	password := c.scanner.Text()

	res, err := c.userClient.RegisterUser(context.Background(), &pbUser.RegisterUserRequest{
		Login:    login,
		Password: password,
	})

	if err != nil {
		fmt.Printf("Ошибка регистрации: %v\n", err)
		return err
	}

	token := res.GetAccessToken()
	c.authInterceptor.AccessToken = token

	return nil
}

func (c *Client) authUser() {
	for {
		fmt.Println("Вы уже зарегистрированы? (y/n)")
		c.scanner.Scan()
		answer := c.scanner.Text()

		if answer == "y" {
			err := c.login()

			if err == nil {
				break
			}
		} else if answer == "n" {
			err := c.register()

			if err == nil {
				break
			}
		}
	}
}
