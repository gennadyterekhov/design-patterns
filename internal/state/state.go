// https://habrastorage.org/r/w1560/getpro/habr/post_images/4fb/72c/31e/4fb72c31e718d89f6cefb97af767b6ee.jpg
package state

type GameState struct {
	ID               string
	CurrentTurnIndex int
	PlayerHealth     int
	PlayerMana       int
	TotalPoints      int
}

func (gs *GameState) OddTurn() {
	gs.CurrentTurnIndex++
	gs.PlayerHealth--
	gs.PlayerMana--
	gs.TotalPoints--

}
func (gs *GameState) EvenTurn() {
	gs.CurrentTurnIndex++
	gs.PlayerHealth++
	gs.PlayerMana++
	gs.TotalPoints++

}

// not necessary for the patterns, just for demo
// a method to change state according to logic
func (gs *GameState) MakeATurn() {
	if gs.CurrentTurnIndex%2 == 0 {
		gs.EvenTurn()
	} else {
		gs.OddTurn()
	}
}
