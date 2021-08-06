package burger

import (
	"fmt"
	"github.com/bimbimprasetyoafif/simple-go-design-pattern/chef"
	"github.com/bimbimprasetyoafif/simple-go-design-pattern/food"
)

type burger struct {
	name    string
	flavour string
}

func (b burger) GetName() string {
	return b.name
}

func (b burger) GetFlavour() string {
	return b.flavour
}

type chefBurger struct {
	flavour string
}

func (c chefBurger) Preparation() {
	fmt.Println("chef preparing burger")
}

func (c *chefBurger) AddFlavour(flavour string) {
	c.flavour = flavour
}

func (c *chefBurger) AddOptional(opts ...string) {
	for i, opt := range opts {
		if i == len(opts)-1 {
			c.flavour = fmt.Sprintf("%s and %s", c.flavour, opt)
			continue
		}
		c.flavour = fmt.Sprintf("%s, %s", c.flavour, opt)
	}

}

func (c chefBurger) Bake() food.Food {
	return &burger{
		name:    "burger",
		flavour: c.flavour,
	}
}

func CallBurgerChef() chef.DoChef {
	return &chefBurger{}
}
