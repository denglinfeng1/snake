package model

import (
	"time"

	"gopkg.in/mgo.v2"
	//"gopkg.in/mgo.v2/bson"
)

func MdbGetSession() (session *mgo.Session, err error) {
	dail_info := &mgo.DialInfo{
		Addrs:  []string{"127.0.0.1"},
		Direct: false,
		Timeout: time.Second * 1,
		Database: "game",
		Username: "test1",
		Password: "123456",
		PoolLimit: 1024,
	}

	session, err = mgo.DialWithInfo(dail_info)
	return
}
