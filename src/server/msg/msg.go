package msg

import (
	"github.com/name5566/leaf/network/json"
)

var Processor = json.NewProcessor()

func init() {
	Processor.Register(&Demo{})
	Processor.Register(&Hello{})
	Processor.Register(&JoinRoom{})
	Processor.Register(&Fight{})
	Processor.Register(&CheckOnline{})
}

type Demo struct {
}

type Hello struct {
	Name string
}

type Fight struct {
}

type CheckOnline struct {
	Type      int
	CheckTime int
}

// 加入房间
type JoinRoom struct {
	Type int
}
