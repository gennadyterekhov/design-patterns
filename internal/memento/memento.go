// essentially it's like a save file
// https://habrastorage.org/r/w1560/getpro/habr/post_images/c08/bf1/7ee/c08bf17ee80d42272441cafbcce1a2dd.jpg
package memento

// object with some internal state aka Originator
type GameState struct {
	ID               string
	CurrentTurnIndex int
	PlayerHealth     int
	PlayerMana       int
	TotalPoints      int
}

// memento - just the data we need to save/remember
type GameStateDto struct {
	CurrentTurnIndex int
	PlayerHealth     int
	PlayerMana       int
	TotalPoints      int
}

func (gs *GameState) createMemento() *GameStateDto {
	return &GameStateDto{
		CurrentTurnIndex: gs.CurrentTurnIndex,
		PlayerHealth:     gs.PlayerHealth,
		PlayerMana:       gs.PlayerMana,
		TotalPoints:      gs.TotalPoints,
	}
}

func (gs *GameState) setMemento(mem *GameStateDto) {
	gs.CurrentTurnIndex = mem.CurrentTurnIndex
	gs.PlayerHealth = mem.PlayerHealth
	gs.PlayerMana = mem.PlayerMana
	gs.TotalPoints = mem.TotalPoints
}

// not necessary for the patterns, just for demo
// a method to change state according to logic
func (gs *GameState) makeATurn() {
	gs.CurrentTurnIndex++
	gs.PlayerHealth--
	gs.PlayerMana--
	gs.TotalPoints++
}
