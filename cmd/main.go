package main

import (
	"fmt"
	"log"

	"github.com/j4ndrw/dungeons-and-dragons-template-creation/pkg/formrunners"
	"github.com/j4ndrw/dungeons-and-dragons-template-creation/pkg/utils"
)

func main() {
	utils.ClearScreen()

	numberOfPlayers, err := formrunners.NumberOfPlayers()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("You have", *numberOfPlayers, "players at the table.")

	creatures, err := formrunners.CharacterCreation(*numberOfPlayers)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("The creatures are:\n")
	for _, c := range creatures {
		fmt.Println(c.ToString())
	}
}
