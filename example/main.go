package main

import (
	"fmt"
	"github.com/erfanmomeniii/factory"
)

type Name struct {
	FirstName string
	LastName  string
}
type Info struct {
	Name  Name
	Phone int
}

func main() {
	f := factory.NewFactory()

	instances := f.Model(Info{}).
		Generate(2)

	for _, i := range instances {
		instance := i.(Info)
		fmt.Println("Name : ", instance.Name.FirstName+" "+instance.Name.LastName, " | Phone : ", instance.Phone)
	}
}
