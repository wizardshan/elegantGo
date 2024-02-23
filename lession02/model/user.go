package model

import (
	"fmt"
	"strconv"
	"time"
)

type User struct {
	ID         int
	Level      int
	Mobile     string
	Nickname   string
	Avatar     string
	Amount     int
	LoginTime  time.Time
	CreateTime time.Time
}

func (user *User) LevelDesc() string {

	switch user.Level {
	case 0:
		return "普通"
	case 10:
		return "白银"
	case 20:
		return "黄金"
	case 30:
		return "铂金"
	}

	return "未知等级"
}

func (user *User) ToString() []string {

	return []string{
		strconv.Itoa(user.ID),
		user.LevelDesc(),
		strconv.Itoa(user.Amount / 100),
		fmt.Sprintf("%s****%s", user.Mobile[0:4], user.Mobile[8:11]),
		user.Nickname,
		user.CreateTime.Format(time.DateTime),
	}
}
