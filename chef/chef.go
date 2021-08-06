package chef

import "github.com/bimbimprasetyoafif/simple-go-design-pattern/food"

type DoChef interface {
	Preparation()
	AddFlavour(flavour string)
	AddOptional(opts ...string)
	Bake() food.Food
}
