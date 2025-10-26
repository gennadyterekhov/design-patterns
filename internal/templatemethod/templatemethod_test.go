package templatemethod

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTemplateMethod(t *testing.T) {

	gs := NewGame()

	assert.Equal(t, 100, gs.PlayerHealth)

	gs.MakeATurn()

	assert.Equal(t, 99, gs.PlayerHealth)

	gs = NewModdedGame(CheatTurnHandler)
	gs.MakeATurn()

	assert.Equal(t, 1000, gs.PlayerHealth)
}
