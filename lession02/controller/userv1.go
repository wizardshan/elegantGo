package controller

import (
	"bytes"
	"elegantGo/lession02/repository"
	"encoding/csv"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"time"
)

type UserV1 struct {
	repo *repository.User
}

func NewUserV1(repo *repository.User) *UserV1 {
	ctr := new(UserV1)
	ctr.repo = repo
	return ctr
}

func (ctr *UserV1) Export(c *gin.Context) {
	users := ctr.repo.All()

	var buff = new(bytes.Buffer)
	wr := csv.NewWriter(buff)
	heads := []string{"用户id", "等级", "余额", "手机号码", "昵称", "创建时间"}
	wr.Write(heads)

	amountTotal := 0
	for _, u := range users {
		levelDesc := ""
		switch u.Level {
		case 0:
			levelDesc = "普通"
		case 10:
			levelDesc = "白银"
		case 20:
			levelDesc = "黄金"
		case 30:
			levelDesc = "铂金"
		}

		wr.Write([]string{
			strconv.Itoa(u.ID),
			levelDesc,
			strconv.Itoa(u.Amount / 100),
			fmt.Sprintf("%s****%s", u.Mobile[0:4], u.Mobile[8:11]),
			u.Nickname,
			u.CreateTime.Format(time.DateTime),
		})

		amountTotal += u.Amount
	}

	wr.Write([]string{
		fmt.Sprintf("用户总数：%d 用户总余额：%d", len(users), amountTotal/100),
	})

	wr.Flush()

	c.Writer.Header().Set("Content-type", "application/octet-stream")
	c.Writer.Header().Set("Content-Disposition", fmt.Sprintf("attachment;filename=%s.csv", fmt.Sprintf("users-%s", time.Now().Format("20060102150405"))))
	c.String(http.StatusOK, buff.String())
}
