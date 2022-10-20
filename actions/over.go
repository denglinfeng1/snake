package actions

import (
	"snake/model"
	"time"

	"gopkg.in/mgo.v2/bson"
)

type msgOver struct {
	ps []*model.PlayerInfo
}

// syncStart 异步：开始游戏
/*
此处可以是一个通用的interface,改为通用接受的消息，异步消息统一处理了
*/
func syncOver(msg *msgStart) {
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
			Type:      "over",
			Timestamp: time.Now().Unix(),
		}},
	}
	// over
	session.DB("game").C("game").Update(game, bson.M{"over": 1})
}
