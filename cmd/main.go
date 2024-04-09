package main

import (
	"fmt"
	"log"
	"strconv"

	"github.com/j4ndrw/dungeons-and-dragons-template-creation/internal/creature"
	"github.com/j4ndrw/dungeons-and-dragons-template-creation/pkg/forms"
	"github.com/j4ndrw/dungeons-and-dragons-template-creation/pkg/utils"
)

func main() {
	utils.ClearScreen()

	var numberOfPlayersShell string
	numberOfPlayersForm := forms.NumberOfPlayers(&numberOfPlayersShell)
	err := numberOfPlayersForm.Run()
	if err != nil {
		log.Fatal(err)
	}
	numberOfPlayers, err := strconv.Atoi(numberOfPlayersShell)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("You have", numberOfPlayers, "players at the table.")

	var creatures []*creature.Creature
	for i := 0; i < numberOfPlayers; i++ {
		creature := creature.Empty()
		characterCreationForm := forms.CharacterCreation(creature)
		err := characterCreationForm.Run()
		if err != nil {
			log.Fatal(err)
		}
		creature = creature.Build()
		creatures = append(creatures, creature)
		fmt.Printf("Created creature: %s\n", creature.Name)
	}

	fmt.Printf("The creatures are:\n")
	for _, c := range creatures {
		fmt.Println(c.ToString())
	}
}
