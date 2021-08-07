package chef

import "github.com/bimbimprasetyoafif/simple-go-design-pattern/food"

type DoChef interface {
	Preparation() DoChef
	AddFlavour(flavour string) DoChef
	AddOptional(opts ...string) DoChef
	CookAndServe() food.AllFood
}
