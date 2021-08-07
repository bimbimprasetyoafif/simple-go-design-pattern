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

func (c *cashier) ServeFood(flavour string, opts ...string) food.AllFood {
	return c.chef.Preparation().
		AddFlavour(flavour).
		AddOptional(opts...).
		CookAndServe()
}
