package client

import (
	"context"
	"fmt"
	pbUser "gophkeeper/internal/proto/user"
)

func (c *Client) Login() error {
	fmt.Println("Введите логин:")
	c.Scanner.Scan()
	login := c.Scanner.Text()

	fmt.Println("Введите пароль:")
	c.Scanner.Scan()
	password := c.Scanner.Text()

	res, err := c.UserClient.LoginUser(context.Background(), &pbUser.LoginUserRequest{
		Login:    login,
		Password: password,
	})

	if err != nil {
		fmt.Printf("Ошибка авторизации: %v\n", err)
		return err
	}

	token := res.GetAccessToken()
	c.AuthInterceptor.AccessToken = token

	return nil
}

func (c *Client) Register() error {
	fmt.Println("Введите логин:")
	c.Scanner.Scan()
	login := c.Scanner.Text()

	fmt.Println("Введите пароль:")
	c.Scanner.Scan()
	password := c.Scanner.Text()

	res, err := c.UserClient.RegisterUser(context.Background(), &pbUser.RegisterUserRequest{
		Login:    login,
		Password: password,
	})

	if err != nil {
		fmt.Printf("Ошибка регистрации: %v\n", err)
		return err
	}

	token := res.GetAccessToken()
	c.AuthInterceptor.AccessToken = token

	return nil
}

func (c *Client) AuthUser() {
	for {
		fmt.Println("Вы уже зарегистрированы? (y/n)")
		c.Scanner.Scan()
		answer := c.Scanner.Text()

		if answer == "y" {
			err := c.Login()

			if err == nil {
				break
			}
		} else if answer == "n" {
			err := c.Register()

			if err == nil {
				break
			}
		}
	}
}
