package factory

import "github.com/bimbimprasetyoafif/simple-go-design-pattern/food"

type kfc struct {
}

func (k kfc) ServeDrink() food.Drink {
	return &food.Sprite{}
}
func (k kfc) ServeMeal() food.Meal {
	return &food.Chicken{}
}
