package model

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)

type Action struct {
	Type      string      `bson:"type"` // start, turn, up, down, over, throw
	UserId    int64       `bson:"user_id"`
	Timestamp int64       `bson:"timestamp"`
	Extra     interface{} `bson:"extra"`
}

type PlayerInfo struct {
	UserID int64  `bson:"user_id" json:"user_id"`
	Name   string `bson:"name" json:"name"`
	Thumb  string `bson:"thumb" json:"thumb"`
	Status string `bson:"status" json:"status"`
	Score  int    `bson:"score" json:"score"`
}

type Snake struct {
	ID         bson.ObjectId `bson:"_id" json:"id"`
	Players    []*PlayerInfo `bson:"players" json:"players"`
	TurnuserId int64         `bson:"turn_index" json:"turn_index"`
	EndReason  string        `bson:"end_reason" json:"end_reason"`
	StartAt    time.Time     `bson:"start_at" json:"start_at"`
	EndAt      time.Time     `bson:"end_at" json:"end_at"`
	Actions    []*Action     `bson:"actions" json:"actions"`
}
