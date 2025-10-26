package state

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestState(t *testing.T) {

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

	// An event happens that changes your internal state
	// we dont EXPLICITLY change the mood, we just use some logic
	// that happens to change state
	// we dont have a method to change state to whatever we want (important)
	gs.HearBadNews()

	// and now the methods return something else
	assert.Equal(t, "I don't deserve it...", gs.ReactToCompliment())
	assert.Equal(t, "You're right, I'm the worst.", gs.ReactToCriticism())
}
