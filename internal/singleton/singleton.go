// https://habrastorage.org/r/w1560/getpro/habr/post_images/f80/871/aaf/f80871aaf46238adcc0cd2468f19a4c5.jpg
package singleton

var gameInst *gameState

// type must be private to ensure only 1 instance can be created
type gameState struct {
	ID               string
	CurrentTurnIndex int
	PlayerHealth     int
	PlayerMana       int
	TotalPoints      int
}

func NewGame() *gameState {
	if gameInst != nil {
		return gameInst
	}
	gameInst = &gameState{
		ID:               "unique-id",
		CurrentTurnIndex: 1,
		PlayerHealth:     100,
		PlayerMana:       100,
		TotalPoints:      0,
	}
	return gameInst
}

func (gs *gameState) MakeATurn() {
	gs.CurrentTurnIndex++
	gs.PlayerHealth--
	gs.PlayerMana--
	gs.TotalPoints++
}
