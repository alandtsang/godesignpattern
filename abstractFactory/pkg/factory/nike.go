package factory

type Nike struct {
}

func NewNike() *Nike {
	return &Nike{}
}

func (n *Nike) MakeShoe() Shoe {
	return &nikeShoe{
		logo: "Nike",
	}
}

func (n *Nike) MakeShort() Short {
	return &nikeShort{
		size: 14,
	}
}

type nikeShoe struct {
	logo string
}

func (n *nikeShoe) SetLogo(logo string) {
	n.logo = logo
}

func (s *nikeShoe) GetLogo() string {
	return s.logo
}

type nikeShort struct {
	size int
}

func (n *nikeShort) SetSize(size int) {
	n.size = size
}

func (n *nikeShort) GetSize() int {
	return n.size
}
