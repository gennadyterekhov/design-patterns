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
		Player:           Player{Mood: "happy"},
	}

	assert.Equal(t, "Thank you! That's so kind!", gs.ReactToCompliment())
	assert.Equal(t, "I'll try to do better, thanks!", gs.ReactToCriticism())

	gs.HearBadNews()

	assert.Equal(t, "I don't deserve it...", gs.ReactToCompliment())
	assert.Equal(t, "You're right, I'm the worst.", gs.ReactToCriticism())
}
