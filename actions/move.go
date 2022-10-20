package actions

import (
	"snake/model"
	"time"

	"gopkg.in/mgo.v2/bson"
)

type msgMove struct {
	GameID bson.ObjectId
	ps     []*model.PlayerInfo
}

// syncStart 异步
/*
此处可以是一个通用的interface,改为通用接受的消息，异步消息统一处理了
*/
func syncMove(msg *msgMove) {
	game := model.GetGameById(msg.GameID.Hex())
	session, err := model.MdbGetSession()
	if err != nil {
		time.Sleep(time.Second)
		syncMove(msg)
	}
	actions := append(game.Actions, &model.Action{
		Type:      "move",
		Timestamp: time.Now().Unix(),
		UserId:    1,
		Extra: "up/down",
	})
	session.DB("game").C("game").Update(bson.M{"_id": msg.GameID}, bson.M{"$set": bson.M{"actions": actions}})
	// ws send
	// done
	return
}
