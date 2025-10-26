// https://habrastorage.org/r/w1560/getpro/habr/post_images/4fb/72c/31e/4fb72c31e718d89f6cefb97af767b6ee.jpg
package state

type GameState struct {
	ID               string
	CurrentTurnIndex int
	PlayerHealth     int
	PlayerMana       int
	TotalPoints      int
	Player
}
type Player struct {
	Mood string
}

func (gs *GameState) ReactToCompliment() string {
	if gs.Player.Mood == "sad" {
		return "I don't deserve it..."
	}
	return "Thank you! That's so kind!"
}

func (gs *GameState) ReactToCriticism() string {
	if gs.Player.Mood == "sad" {
		return "You're right, I'm the worst."
	}
	return "I'll try to do better, thanks!"
}

func (gs *GameState) HearBadNews() {
	gs.Player.Mood = "sad"
}
