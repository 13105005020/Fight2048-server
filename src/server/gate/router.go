package gate

import (
	"game/server/src/server/game"
	"game/server/src/server/msg"
)

func init() {
	msg.Processor.SetRouter(&msg.Demo{}, game.ChanRPC)
	msg.Processor.SetRouter(&msg.Hello{}, game.ChanRPC)
	msg.Processor.SetRouter(&msg.JoinRoom{}, game.ChanRPC)
	msg.Processor.SetRouter(&msg.Fight{}, game.ChanRPC)
	msg.Processor.SetRouter(&msg.CheckOnline{}, game.ChanRPC)
}
