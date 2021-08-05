package main

import (
	"fmt"
	"github.com/bimbimprasetyoafif/simple-go-design-pattern/factory"
)

func main() {
	cashierMCD := factory.NewRestaurantFactory(factory.MCD)
	cashierKFC := factory.NewRestaurantFactory(factory.KFC)

	fmt.Println("=============== DRINK ==============")
	drinkMCD := cashierMCD.ServeDrink()
	drinkKFC := cashierKFC.ServeDrink()
	fmt.Println(drinkMCD.Flavour())
	fmt.Println(drinkKFC.Flavour())

	fmt.Println("=============== DRINK ==============")
	mealMCD := cashierMCD.ServeMeal()
	mealKFC := cashierKFC.ServeMeal()
	fmt.Println(mealMCD.Eat())
	fmt.Println(mealKFC.Eat())

}
