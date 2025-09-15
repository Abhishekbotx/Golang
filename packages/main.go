package main

import (
	"fmt"

	"github.com/abhishekbotx/golang/packages/auth"
	"github.com/abhishekbotx/golang/packages/user"
	"github.com/fatih/color"
)

func main() {
	auth.LoginWithCredentials("abhi","secret")
	session:=auth.GetSession()
	fmt.Println("session",session)

	user:=user.User{
		Email: "abhi@gmail.com",
		Name:"John Doe",
	}

	fmt.Println(user.Email,user.Name)

	color.Red(user.Email)
}