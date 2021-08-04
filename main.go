package main

import (
	"fmt"
	"github.com/bimbimprasetyoafif/simple-go-design-pattern/factory"
)

func main() {
	duckyKeyboard := factory.NewKeyboardFactory(factory.Ducky)
	razorKeyboard := factory.NewKeyboardFactory(factory.Razor)

	fmt.Println(duckyKeyboard.OnBacklit())
	fmt.Println(duckyKeyboard.Type())

	fmt.Println(razorKeyboard.OnBacklit())
	fmt.Println(razorKeyboard.Type())
}
