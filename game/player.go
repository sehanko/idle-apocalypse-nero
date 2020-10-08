package game

type Player struct {
	Minerals Minerals
}

func NewPlayer() Player {
	return Player{ZeroMinerals}
}