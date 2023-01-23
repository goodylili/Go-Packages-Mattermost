package main


import (
	"fmt"
	"syreclabs.com/go/faker"
)

func main() {
	name := faker.Name()
	email := faker.App()
	btc := faker.Bitcoin()

	fmt.Println("Name:", name)
	fmt.Println("Email:", email)
	fmt.Println("Bitcoin:", btc)
}
