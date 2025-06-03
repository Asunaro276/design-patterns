package ware

type adidas struct {}

func (a *adidas) MakeShoe() IShoe {
	return &adidasShoe{
		shoe: shoe{
			logo: "adidas",
			size: 14,
		},
	}
}

func (a *adidas) MakeShirt() IShirt {
	return &adidasShirt{
		shirt: shirt{
			logo: "adidas",
			size: 14,
		},
	}
}

type adidasShoe struct {
	shoe
}

type adidasShirt struct {
	shirt
}
