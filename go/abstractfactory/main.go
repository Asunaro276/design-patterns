package main

import (
	"fmt"

	ware "github.com/asunaro276/design-patterns/go/abstractfactory/internal"
)

func main() {
	adidasFactory, _ := ware.GetSportsFactory("adidas")
	nikeFactory, _ := ware.GetSportsFactory("nike")

	adidasShoe := adidasFactory.MakeShoe()
	nikeShoe := nikeFactory.MakeShoe()

	adidasShirt := adidasFactory.MakeShirt()
	nikeShirt := nikeFactory.MakeShirt()

	printShoeDetails(adidasShoe)
	printShoeDetails(nikeShoe)
	printShirtDetails(adidasShirt)
	printShirtDetails(nikeShirt)
}

func printShoeDetails(s ware.IShoe) {
	fmt.Printf("Logo: %s", s.GetLogo())
	fmt.Println()
	fmt.Printf("Size: %d", s.GetSize())
	fmt.Println()
}

func printShirtDetails(s ware.IShirt) {
	fmt.Printf("Logo: %s", s.GetLogo())
	fmt.Println()
	fmt.Printf("Size: %d", s.GetSize())
	fmt.Println()
}
