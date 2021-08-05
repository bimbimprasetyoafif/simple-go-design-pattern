package factory

import "github.com/bimbimprasetyoafif/simple-go-design-pattern/food"

type mcd struct {
}

func (m mcd) ServeDrink() food.Drink {
	return &food.CocaCola{}
}
func (m mcd) ServeMeal() food.Meal {
	return &food.Burger{}
}
