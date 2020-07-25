package factory

import (
	"fmt"
)

type SportsFactory interface {
	MakeShoe() Shoe
	MakeShort() Short
}

func GetSportsFactory(brand string) (SportsFactory, error) {
	if brand == "nike" {
		return NewNike(), nil
	} else if brand == "adidas" {
		return NewAdidas(), nil
	}
	return nil, fmt.Errorf("Wrong brand type passed")
}

type Shoe interface {
	SetLogo(logo string)
	GetLogo() string
}

type Short interface {
	SetSize(size int)
	GetSize() int
}

func PrintShoe(s Shoe) {
	fmt.Println("shoe logo:", s.GetLogo())
}

func PrintShort(s Short) {
	fmt.Println("short size:", s.GetSize())
}
