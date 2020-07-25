package factory

type Adidas struct {
}

func NewAdidas() *Adidas {
	return &Adidas{}
}

func (a *Adidas) MakeShoe() Shoe {
	return &adidasShoe{
		logo: "Adidas",
	}
}

func (a *Adidas) MakeShort() Short {
	return &adidasShort{
		size: 20,
	}
}

type adidasShoe struct {
	logo string
}

func (a *adidasShoe) SetLogo(logo string) {
	a.logo = logo
}

func (a *adidasShoe) GetLogo() string {
	return a.logo
}

type adidasShort struct {
	size int
}

func (a *adidasShort) SetSize(size int) {
	a.size = size
}

func (a *adidasShort) GetSize() int {
	return a.size
}
