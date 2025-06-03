package ware

type shirt struct {
	logo string
	size int
}

type IShirt interface {
	SetLogo(logo string)
	SetSize(size int)
	GetLogo() string
	GetSize() int
}

func (s *shirt) SetLogo(logo string) {
	s.logo = logo
}

func (s *shirt) SetSize(size int) {
	s.size = size
}

func (s *shirt) GetLogo() string {
	return s.logo
}

func (s *shirt) GetSize() int {
	return s.size
}
