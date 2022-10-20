package actions

import (
	"snake/model"
	"time"

	"gopkg.in/mgo.v2/bson"
)

type msgStart struct {
	ps []*model.PlayerInfo
}

// syncStart 异步：开始游戏
/*
此处可以是一个通用的interface,改为通用接受的消息，异步消息统一处理了
*/
func syncStart(msg *msgStart) {
	session, err := model.MdbGetSession()
	if err != nil {
		time.Sleep(time.Second)
		syncStart(msg)
	}
	game := &model.Snake{
		ID:         bson.NewObjectId(),
		Players:    msg.ps,
		TurnuserId: 1,
		StartAt:    time.Now(),
		Actions: []*model.Action{{
			Type:      "start",
			Timestamp: time.Now().Unix(),
		}},
	}
	session.DB("game").C("game").Insert(game)
}
