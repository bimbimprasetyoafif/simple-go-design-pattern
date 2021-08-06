package cashier

import (
	"github.com/bimbimprasetyoafif/simple-go-design-pattern/chef"
	"github.com/bimbimprasetyoafif/simple-go-design-pattern/food"
)

type cashier struct {
	chef chef.DoChef
}

func CallCashier(chef chef.DoChef) *cashier {
	return &cashier{
		chef: chef,
	}
}

func (c *cashier) ChangeChef(chef chef.DoChef) {
	c.chef = chef
}

func (c *cashier) ServeFood(flavour string, opts ...string) food.Food {
	c.chef.Preparation()
	c.chef.AddFlavour(flavour)
	if opts != nil {
		c.chef.AddOptional(opts...)
	}
	return c.chef.Bake()
}
