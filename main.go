package main

import (
	"fmt"
	"github.com/bimbimprasetyoafif/simple-go-design-pattern/burger"
	"github.com/bimbimprasetyoafif/simple-go-design-pattern/cashier"
	"github.com/bimbimprasetyoafif/simple-go-design-pattern/food"
	"github.com/bimbimprasetyoafif/simple-go-design-pattern/pizza"
)

func main() {
	pChef := pizza.CallPizzaChef()
	rCashier := cashier.CallCashier(pChef)

	pSpicy := rCashier.ServeFood("spicy spicy")
	PrintFood(pSpicy)
	pSalty := rCashier.ServeFood("salltlttyyy", "paperoni")
	PrintFood(pSalty)

	bChef := burger.CallBurgerChef()
	rCashier.ChangeChef(bChef)

	bCheese := rCashier.ServeFood("chessyyy", "extra beef", "less sauce")
	PrintFood(bCheese)
	bSweet := rCashier.ServeFood("creammy sweet")
	PrintFood(bSweet)
}

func PrintFood(food food.Food) {
	fmt.Println("=========", food.GetName(), "=========")
	fmt.Println("it's", food.GetFlavour())
	fmt.Println("===========================")
	fmt.Println()
}
