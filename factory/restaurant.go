package factory

import (
	"github.com/bimbimprasetyoafif/simple-go-design-pattern/food"
)

type Resto string

const (
	KFC Resto = "kfc"
	MCD Resto = "mcd"
)

type Restaurant interface {
	ServeDrink() food.Drink
	ServeMeal() food.Meal
}

func NewRestaurantFactory(resto Resto) Restaurant {
	switch resto {
	case MCD:
		return &mcd{}
	case KFC:
		return &kfc{}
	}

	return nil
}
