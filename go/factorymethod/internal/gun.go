package gun

type Gun interface {
	SetName(name string)
	SetPower(power int)
	GetName() string
	GetPower() int
}

type gun struct {
	name string
	power int
}	

func (g *gun) SetName(name string) {
	g.name = name
}

func (g *gun) GetName() string {
	return g.name
}

func (g *gun) SetPower(power int) {
	g.power = power
}

func (g *gun) GetPower() int {
	return g.power
}
