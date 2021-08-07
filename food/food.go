package food

type AllFood interface {
	GetName() string
	GetFlavour() string
}

type Food struct {
	Name    string
	Flavour string
}

func (p Food) GetName() string {
	return p.Name
}
func (p Food) GetFlavour() string {
	return p.Flavour
}
