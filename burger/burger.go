package burger

import (
	"fmt"
	"github.com/bimbimprasetyoafif/simple-go-design-pattern/chef"
	"github.com/bimbimprasetyoafif/simple-go-design-pattern/food"
)

type chefBurger struct {
	food food.Food
}

func (c *chefBurger) Preparation() chef.DoChef {
	fmt.Println("chef preparing burger")
	c.food.Name = "burger"
	return c
}

func (c *chefBurger) AddFlavour(flavour string) chef.DoChef {
	c.food.Flavour = flavour
	return c
}

func (c *chefBurger) AddOptional(opts ...string) chef.DoChef {
	if opts == nil {
		return c
	}

	for i, opt := range opts {
		if i == len(opts)-1 {
			c.food.Flavour = fmt.Sprintf("%s and %s", c.food.GetFlavour(), opt)
			continue
		}
		c.food.Flavour = fmt.Sprintf("%s, %s", c.food.GetFlavour(), opt)
	}

	return c
}

func (c *chefBurger) CookAndServe() food.AllFood {
	return c.food
}

func CallBurgerChef() chef.DoChef {
	return &chefBurger{}
}
