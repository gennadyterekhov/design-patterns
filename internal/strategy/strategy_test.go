package strategy

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStrategy(t *testing.T) {

	gs := &GameState{
		ID:               "unique-id",
		CurrentTurnIndex: 1,
		PlayerHealth:     100,
		PlayerMana:       100,
		TotalPoints:      0,
		Player:           Player{Name: "sonic"},
	}

	assert.Equal(t, 0, gs.TurnsAsSonic)
	assert.Equal(t, 0, gs.TurnsAsTails)

	gs.MakeATurn()

	assert.Equal(t, 1, gs.TurnsAsSonic)
	assert.Equal(t, 0, gs.TurnsAsTails)

	tails := Player{Name: "tails"}
	// we explicitly change our strategy (difference with 'state' pattern)
	gs.SetPlayer(tails)

	assert.Equal(t, 1, gs.TurnsAsSonic)
	assert.Equal(t, 0, gs.TurnsAsTails)

	gs.MakeATurn()

	assert.Equal(t, 1, gs.TurnsAsSonic)
	assert.Equal(t, 1, gs.TurnsAsTails)

}
