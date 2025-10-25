package state

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestObserver(t *testing.T) {

	gs := &GameState{
		ID:               "unique-id",
		CurrentTurnIndex: 1,
		PlayerHealth:     100,
		PlayerMana:       100,
		TotalPoints:      0,
	}

	gs.MakeATurn()
	assert.Equal(t, 99, gs.PlayerHealth)

	gs.MakeATurn()

	assert.Equal(t, 3, gs.CurrentTurnIndex)
	assert.Equal(t, 100, gs.PlayerHealth)

}
