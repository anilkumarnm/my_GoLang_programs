package main

import "fmt"

//VSCode will automatically add an import statement since we used Print in main func

func main() {
	// cards := []string{"Ace of Diamonds", newCard()} //Slice
	// cards = append(cards, "Six of Spades")          // Add new element in slice

	cards := deck{"Ace of Diamonds", newCard()} //Slice
	cards = append(cards, "Six of Spades")      // Add new element in slice

	fmt.Println(cards)
	cards.print()

}

func newCard() string { //This will say that the return value of this function will be a "string"
	return "Five of Diamonds"
}
