package pizza

import (
	"fmt"
	"github.com/bimbimprasetyoafif/simple-go-design-pattern/chef"
	"github.com/bimbimprasetyoafif/simple-go-design-pattern/food"
)

type pizza struct {
	name    string
	flavour string
}

func (p pizza) GetName() string {
	return p.name
}
func (p pizza) GetFlavour() string {
	return p.flavour
}

type chefPizza struct {
	flavour string
}

func (p chefPizza) Preparation() {
	fmt.Println("chef preparing pizza")
}

func (p *chefPizza) AddFlavour(flavour string) {
	p.flavour = flavour
}

func (p *chefPizza) AddOptional(opts ...string) {
	for i, opt := range opts {
		if i == len(opts)-1 {
			p.flavour = fmt.Sprintf("%s and %s", p.flavour, opt)
			continue
		}
		p.flavour = fmt.Sprintf("%s, %s", p.flavour, opt)
	}
}

func (p chefPizza) Bake() food.Food {
	return &pizza{
		name:    "pizza",
		flavour: p.flavour,
	}
}

func CallPizzaChef() chef.DoChef {
	return &chefPizza{}
}
