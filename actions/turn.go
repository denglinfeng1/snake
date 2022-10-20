package actions

import (
	"snake/model"
	"time"

	"gopkg.in/mgo.v2/bson"
)

type msgTurn struct {
	GameID bson.ObjectId
	ps     []*model.PlayerInfo
}

// syncStart 异步：开始游戏
/*
此处可以是一个通用的interface,改为通用接受的消息，异步消息统一处理了
*/
func syncTurn(msg *msgTurn) {
	game := model.GetGameById(msg.GameID.Hex())
	session, err := model.MdbGetSession()
	if err != nil {
		time.Sleep(time.Second)
		syncTurn(msg)
	}
	actions := append(game.Actions, &model.Action{
		Type:      "turn",
		Timestamp: time.Now().Unix(),
		UserId:    1,
	})
	session.DB("game").C("game").Update(bson.M{"_id": msg.GameID}, bson.M{"$set": bson.M{"actions": actions}})
	// ws send
	// done
	return
}
