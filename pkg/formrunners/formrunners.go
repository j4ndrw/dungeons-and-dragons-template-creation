package formrunners

import (
	"strconv"

	"github.com/j4ndrw/dungeons-and-dragons-template-creation/internal/creature"
	"github.com/j4ndrw/dungeons-and-dragons-template-creation/pkg/forms"
)

func NumberOfPlayers() (*int, error) {
	var numberOfPlayersShell string
	numberOfPlayersForm := forms.NumberOfPlayers(&numberOfPlayersShell)
	err := numberOfPlayersForm.Run()
	if err != nil {
	  return nil, err
	}
	numberOfPlayers, err := strconv.Atoi(numberOfPlayersShell)
	if err != nil {
	  return nil, err
	}
	return &numberOfPlayers, nil
}

func CharacterCreation(numberOfPlayers int) ([]*creature.Creature, error) {
	var creatures []*creature.Creature
	for i := 0; i < numberOfPlayers; i++ {
		creature := creature.Empty()
		characterCreationForm := forms.CharacterCreation(creature)
		err := characterCreationForm.Run()
		if err != nil {
		  return nil, err
		}
		creature = creature.Build()
		creatures = append(creatures, creature)
	}
	return creatures, nil
}
