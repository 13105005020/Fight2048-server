package internal

import (
	"game/server/src/server/msg"
	"github.com/name5566/leaf/gate"
	"github.com/name5566/leaf/log"
	"reflect"
	"strconv"
	"time"
)

func init() {
	// 向当前模块（game 模块）注册 Hello 消息的消息处理函数 handleHello
	handler(&msg.Demo{}, handleRoom)
	handler(&msg.JoinRoom{}, handleRoom)
	handler(&msg.Hello{}, handleHello)
	handler(&msg.Fight{}, handleFight)
	handler(&msg.CheckOnline{}, handleCheck)
	chatRoom.Start = make(map[gate.Agent]int)
	go checkOnline()
}

func checkOnline() {
	for {
		for k, v := range chatRoom.LocationSlice {
			if v.CheckTime != 0 && int(time.Now().Unix())-v.CheckTime > 5 {
				chatRoom.LocationSlice = append(chatRoom.LocationSlice[:k], chatRoom.LocationSlice[k:]...)
			}
		}
		time.Sleep(4 * time.Second)
	}
}

func handleCheck(args []interface{}) {
	// 收到的 Hello 消息
	m := args[0].(*msg.CheckOnline)
	// 消息的发送者
	a := args[1].(gate.Agent)
	flag := 0
	for k, v := range chatRoom.LocationSlice {
		if v.Agent == a {
			chatRoom.LocationSlice[k].CheckTime = m.CheckTime
			flag = 1
		}
	}
	a.WriteMsg(&msg.CheckOnline{
		Type:      flag,
		CheckTime: int(time.Now().Unix()),
	})
}

func handler(m interface{}, h interface{}) {
	skeleton.RegisterChanRPC(reflect.TypeOf(m), h)
}

func handleRoom(args []interface{}) {
	// 收到的 Hello 消息
	//m := args[0].(*msg.JoinRoom)
	// 消息的发送者
	a := args[1].(gate.Agent)
	if !inRoom(a) {
		chatRoom.LocationSlice = append(chatRoom.LocationSlice, user{Agent: a, Health: 5, Number: len(chatRoom.LocationSlice)})
		a.WriteMsg(&msg.Hello{
			Name: "加入成功，编号" + strconv.Itoa(len(chatRoom.LocationSlice)) + ",当前" + strconv.Itoa(len(chatRoom.LocationSlice)) + "位用户",
		})
	} else {
		// 给发送者回应一个 Hello 消息
		a.WriteMsg(&msg.Hello{
			Name: "already in the room",
		})
	}
}

func inRoom(a gate.Agent) bool {
	var flag = false
	for _, v := range chatRoom.LocationSlice {
		if a == v.Agent {
			flag = true
		}
	}
	return flag
}

func handleHello(args []interface{}) {
	// 收到的 Hello 消息
	m := args[0].(*msg.Hello)
	// 消息的发送者
	a := args[1].(gate.Agent)
	// 输出收到的消息的内容
	log.Debug("hello %v", m.Name)
	// 给发送者回应一个 Hello 消息
	a.WriteMsg(&msg.Hello{
		Name: "启动成功",
	})
}

func handleFight(args []interface{}) {
	// 消息的发送者
	a := args[1].(gate.Agent)
	chatRoom.Start[a]++
	//给房间中所有人发送结果
	for _, v := range chatRoom.LocationSlice {
		v.Agent.WriteMsg(&msg.Hello{})
	}
}
