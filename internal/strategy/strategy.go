// https://habrastorage.org/r/w1560/getpro/habr/post_images/8d8/303/cdb/8d8303cdbc70de33f376454c2eb6934a.jpg
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
	Player           Player
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
