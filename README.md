<p align="center">
<img src="./assets/logo.png" width=50% height=50%>
</p>
<p align="center">
<a href="https://pkg.go.dev/github.com/erfanmomeniii/factory?tab=doc"target="_blank">
    <img src="https://img.shields.io/badge/Go-1.20+-00ADD8?style=for-the-badge&logo=go" alt="go version" />
</a>

<img src="https://img.shields.io/badge/license-MIT-magenta?style=for-the-badge&logo=none" alt="license" />
<img src="https://img.shields.io/badge/Version-1.0.1-red?style=for-the-badge&logo=none" alt="version" />
</p>

# factory

`factory` is a lightweight package for generating fake data that helps generate realistic but fictional data for various purposes, such as testing, prototyping, or populating databases with sample data.

# Documentation

## Install

```bash
go get github.com/erfanmomeniii/factory
```   

Next, include it in your application:

```bash
import "github.com/erfanmomeniii/factory"
``` 

## Quick Start
The following examples illustrates how to use this package:

```go
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
```

## Contributing
Pull requests are welcome. For changes, please open an issue first to discuss what you would like to change.
Please make sure to update tests as appropriate.
