// https://habrastorage.org/r/w1560/getpro/habr/post_images/4fb/72c/31e/4fb72c31e718d89f6cefb97af767b6ee.jpg
package strategy

import "fmt"

type GameState struct {
	ID               string
	CurrentTurnIndex int
	PlayerHealth     int
	PlayerMana       int
	TotalPoints      int
	TurnsAsSonic     int
	TurnsAsTails     int
	Player Player
}
type Player struct {
	Name string
}

func (gs *GameState) SetPlayer(p Player) {
	gs.Player = p
}

func (gs *GameState) TurnAsSonic() {
	// use Sonic controls etc
	fmt.Println("sonic")
	gs.TurnsAsSonic++

}

func (gs *GameState) TurnAsTails() {
	// use Tails controls etc
	fmt.Println("tails")
	gs.TurnsAsTails++

}

func (gs *GameState) MakeATurn() {
	if gs.Player.Name == "sonic" {
		gs.TurnAsSonic()
	} else {
		gs.TurnAsTails()
	}
}
