package pizza

import (
	"fmt"
	"github.com/bimbimprasetyoafif/simple-go-design-pattern/chef"
	"github.com/bimbimprasetyoafif/simple-go-design-pattern/food"
)

type chefPizza struct {
	food food.Food
}

func (p *chefPizza) Preparation() chef.DoChef {
	fmt.Println("chef preparing pizza")
	p.food.Name = "pizza"
	return p
}

func (p *chefPizza) AddFlavour(flavour string) chef.DoChef {
	p.food.Flavour = flavour
	return p
}

func (p *chefPizza) AddOptional(opts ...string) chef.DoChef {
	if opts == nil {
		return p
	}
	for i, opt := range opts {
		if i == len(opts)-1 {
			p.food.Flavour = fmt.Sprintf("%s and %s", p.food.GetFlavour(), opt)
			continue
		}
		p.food.Flavour = fmt.Sprintf("%s, %s", p.food.GetFlavour(), opt)
	}

	return p
}

func (p *chefPizza) CookAndServe() food.AllFood {
	return p.food
}

func CallPizzaChef() chef.DoChef {
	return &chefPizza{}
}
