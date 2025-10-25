// https://habrastorage.org/r/w1560/getpro/habr/post_images/cad/355/d48/cad355d48e3b9a2debcad55bc6504beb.jpg
package observer

import "fmt"

////
// Subject
////

// what we need to observe aka subject
// it will notify observers of events
type GameState struct {
	ID               string
	CurrentTurnIndex int
	PlayerHealth     int
	PlayerMana       int
	TotalPoints      int
	Observers        map[string]ObserverInterface
}

func (gs *GameState) Attach(obs ObserverInterface) {
	if gs.Observers == nil {
		gs.Observers = make(map[string]ObserverInterface)
	}
	gs.Observers[obs.GetID()] = obs
}

func (gs *GameState) Detach(obs ObserverInterface) {
	delete(gs.Observers, obs.GetID())

}

// not necessary for the patterns, just for demo
// a method to change state according to logic
func (gs *GameState) Notify() {
	for _, obsInst := range gs.Observers {
		obsInst.Update()
	}
}

// not necessary for the patterns, just for demo
// a method to change state according to logic
func (gs *GameState) MakeATurn() {
	gs.CurrentTurnIndex++
	gs.PlayerHealth--
	gs.PlayerMana--
	gs.TotalPoints++
	gs.Notify()
}

////
// Observer
////

type ObserverInterface interface {
	GetID() string
	Update()
}

type Observer struct {
	ID           string
	UpdateCalled bool
	DebugVal     int
}

func (gs *Observer) GetID() string {
	return gs.ID
}

func (gs *Observer) Update() {
	fmt.Println("gs update", gs.ID)
	gs.DebugVal++
	gs.UpdateCalled = true
}
