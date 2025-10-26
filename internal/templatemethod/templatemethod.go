// https://habrastorage.org/r/w1560/getpro/habr/post_images/2d6/a9d/443/2d6a9d443800486ed3ba432da9954df8.jpg
package templatemethod

import "fmt"

////
// Subject - a game with modding capabilities
////

type GameState struct {
	ID               string
	CurrentTurnIndex int
	PlayerHealth     int
	PlayerMana       int
	TotalPoints      int
	TurnHandler      func(*GameState) error
}

func NewGame() *GameState {
	return &GameState{
		ID:               "unique-id",
		CurrentTurnIndex: 1,
		PlayerHealth:     100,
		PlayerMana:       100,
		TotalPoints:      0,
		TurnHandler:      defaultTurnHandler,
	}
}

func NewModdedGame(th func(*GameState) error) *GameState {
	return &GameState{
		ID:               "unique-id",
		CurrentTurnIndex: 1,
		PlayerHealth:     100,
		PlayerMana:       100,
		TotalPoints:      0,
		TurnHandler:      th,
	}
}

func (gs *GameState) SetTurnHandler(th func(*GameState) error) {
	gs.TurnHandler = th
}

func (gs *GameState) MakeATurn() error {
	fmt.Println("MakeATurn")
	fmt.Println("gs.TurnHandler == nil", gs.TurnHandler == nil)

	if gs.TurnHandler == nil {
		return fmt.Errorf("no turn handler")
	}
	return gs.TurnHandler(gs)
}

func defaultTurnHandler(gs *GameState) error {
	gs.CurrentTurnIndex++
	gs.PlayerHealth--
	gs.PlayerMana--
	gs.TotalPoints++
	return nil
}

// infinite health and mana
func CheatTurnHandler(gs *GameState) error {
	gs.CurrentTurnIndex++
	gs.PlayerHealth = 1000
	gs.PlayerMana = 1000
	gs.TotalPoints++
	return nil
}
