package cashier

import (
	"github.com/bimbimprasetyoafif/simple-go-design-pattern/burger"
	"github.com/bimbimprasetyoafif/simple-go-design-pattern/pizza"
	"testing"
)

func Test_cashier_ServeFood_ChangeChef(t *testing.T) {
	cP := pizza.CallPizzaChef()
	cashier := CallCashier(cP)
	result := cashier.ServeFood("f1")

	if result.GetName() != "pizza" {
		t.Error("it should return pizza")
	}

	cB := burger.CallBurgerChef()
	cashier.ChangeChef(cB)
	result = cashier.ServeFood("f1")

	if result.GetName() != "burger" {
		t.Error("it should return burger")
	}

}

func Test_cashier_ServeFood_Pizza(t *testing.T) {
	cP := pizza.CallPizzaChef()
	cashier := CallCashier(cP)
	result := cashier.ServeFood("f1")

	if result.GetName() != "pizza" {
		t.Error("it should return pizza")
	}

	if result.GetFlavour() != "f1" {
		t.Error("it should return f1")
	}
}

func Test_cashier_ServeFood_Burger(t *testing.T) {
	cB := burger.CallBurgerChef()
	cashier := CallCashier(cB)
	result := cashier.ServeFood("f1")

	if result.GetName() != "burger" {
		t.Error("it should return burger")
	}

	if result.GetFlavour() != "f1" {
		t.Error("it should return f1")
	}
}

func Test_cashier_ServeFood_Pizza_WithOptional(t *testing.T) {
	cP := pizza.CallPizzaChef()
	cashier := CallCashier(cP)
	t.Run("2 optional", func(t *testing.T) {
		result := cashier.ServeFood("f1", "f2")

		if result.GetName() != "pizza" {
			t.Error("it should return pizza")
		}

		if result.GetFlavour() != "f1 and f2" {
			t.Error("it should return f1 and f2")
		}
	})

	t.Run("3 optional", func(t *testing.T) {
		result := cashier.ServeFood("f1", "f2", "f3")

		if result.GetName() != "pizza" {
			t.Error("it should return pizza")
		}

		if result.GetFlavour() != "f1, f2 and f3" {
			t.Error("it should return f1, f2 and f3")
		}
	})
}

func Test_cashier_ServeFood_Burger_WithOptional(t *testing.T) {
	cB := burger.CallBurgerChef()
	cashier := CallCashier(cB)
	t.Run("2 optional", func(t *testing.T) {
		result := cashier.ServeFood("f1", "f2")

		if result.GetName() != "burger" {
			t.Error("it should return burger")
		}

		if result.GetFlavour() != "f1 and f2" {
			t.Error("it should return f1 and f2")
		}
	})

	t.Run("3 optional", func(t *testing.T) {
		result := cashier.ServeFood("f1", "f2", "f3")

		if result.GetName() != "burger" {
			t.Error("it should return burger")
		}

		if result.GetFlavour() != "f1, f2 and f3" {
			t.Error("it should return f1, f2 and f3")
		}
	})
}
