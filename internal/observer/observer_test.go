package observer

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

	obs := &Observer{
		ID:           "unique-obs-id",
		UpdateCalled: false,
	}
	assert.Equal(t, false, obs.UpdateCalled)

	gs.Attach(obs)

	gs.MakeATurn()
	assert.Equal(t, true, obs.UpdateCalled)
}
