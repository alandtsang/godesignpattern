# Abstract Factory

## Definition

Abstract Factory Design Pattern is a creational design pattern that lets you create a family of related objects.

It is an abstraction over the factory pattern.

## Demo

main.go

```go
package main

import (
    "log"

    "github.com/alandtsang/godesignpattern/abstractFactory/pkg/factory"
)

func main() {
    nikeSports, err := factory.GetSportsFactory("nike")
    if err != nil {
        log.Fatal("GetSportsFactory nike failed,", err)
    }

    adidasSports, err := factory.GetSportsFactory("adidas")
    if err != nil {
        log.Fatal("GetSportsFactory adidas failed,", err)
    }

    nikeShoe := nikeSports.MakeShoe()
    nikeShort := nikeSports.MakeShort()
    adidasShoe := adidasSports.MakeShoe()
    adidasShort := adidasSports.MakeShort()

    factory.PrintShoe(nikeShoe)
    factory.PrintShort(nikeShort)
    factory.PrintShoe(adidasShoe)
    factory.PrintShort(adidasShort)
}
```

```
$ go run cmd/main.go
shoe logo: Nike
short size: 14
shoe logo: Adidas
short size: 20

```