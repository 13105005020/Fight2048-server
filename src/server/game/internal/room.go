package internal

import "github.com/name5566/leaf/gate"

var chatRoom room

type room struct {
	MaxUser       int
	Start         map[gate.Agent]int
	LocationSlice []user // 位置map
	ChessMap      map[int]map[int]int
}
