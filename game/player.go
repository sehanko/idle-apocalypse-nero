package game

type Player struct {
	Res Resources
}

func NewPlayer() Player {
	return Player{NewResources()}
}
