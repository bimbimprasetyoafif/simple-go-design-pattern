package main

import (
	"fmt"
	minatoAndKushina "github.com/bimbimprasetyoafif/simple-go-design-pattern/naruto"
)

func main() {
	naruto := minatoAndKushina.CreateNaruto()
	clone := naruto.Clone()
	clone2 := clone.Clone()

	fmt.Println("naruto", naruto.GetCloneID())
	fmt.Println("clone", clone.GetCloneID())
	fmt.Println("clone2", clone2.GetCloneID())
	fmt.Println()
	fmt.Println("naruto addr", &naruto)
	fmt.Println("clone addr", &clone)
	fmt.Println("clone2 addr", &clone2)
}
