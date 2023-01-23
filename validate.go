package main

import (
	"fmt"
	"gopkg.in/go-playground/validator.v9"
}

type User struct {
	UserName string `validate:"required"`
	UserAge  int    `validate:"gte=18,lte=65"`
	UserMail string `validate:"required,email"`
}


func main() {
	validate := validator.New()

	user := User{
		UserName: "Goodness",
		UserAge:  67,
		UserMail: "john@example.com",
	}

	err := validate.Struct(user)
	if err != nil {
		// action if  Struct is invalid
	}
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			fmt.Println(err.Field(), err.Tag())
		}
	}
}
