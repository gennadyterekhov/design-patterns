package memento

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMemento(t *testing.T) {

	gs := &GameState{
		ID:               "unique-id",
		CurrentTurnIndex: 1,
		PlayerHealth:     100,
		PlayerMana:       100,
		TotalPoints:      0,
	}

	saveOnTurn1 := gs.createMemento()

	gs.makeATurn()

	assert.Equal(t, 2, gs.CurrentTurnIndex)
	gs.makeATurn()
	assert.Equal(t, 3, gs.CurrentTurnIndex)

	gs.setMemento(saveOnTurn1)
	assert.Equal(t, 1, gs.CurrentTurnIndex)
}
