package food

type Drink interface {
	Flavour() string
}

type Meal interface {
	Eat() string
}
