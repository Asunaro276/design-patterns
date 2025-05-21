package ware

type nike struct {}

func (n *nike) MakeShoe() IShoe {
	return &nikeShoe{
		shoe: shoe{
			logo: "nike",
			size: 14,
		},
	}
}

func (n *nike) MakeShirt() IShirt {
	return &nikeShirt{
		shirt: shirt{
			logo: "nike",
			size: 14,
		},
	}
}

type nikeShoe struct {
	shoe
}

type nikeShirt struct {
	shirt
}
