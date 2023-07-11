package main

import (
	"fmt"
	"github.com/erfanmomeniii/factory"
)

type Info struct {
	Name  string
	Phone int
}

func main() {
	f := factory.NewFactory()

	instances := f.Model(Info{}).
		Set("Name", "Erfan").
		Generate(2)

	for _, i := range instances {
		instance := i.(Info)
		fmt.Println("Name : ", instance.Name, " | Phone : ", instance.Phone)
	}
}
