package singleton

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSingleton(t *testing.T) {

	gs := NewGame()

	assert.Equal(t, 1, gs.CurrentTurnIndex)

	gs.MakeATurn()
	assert.Equal(t, 2, gs.CurrentTurnIndex)

	// try to create new game so that turns begin from 1
	gs = NewGame()

	// but they are still the same because it's the same instance
	assert.Equal(t, 2, gs.CurrentTurnIndex)

	gs.MakeATurn()
	assert.Equal(t, 3, gs.CurrentTurnIndex)
}
