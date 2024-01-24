package controller

import (
	"bytes"
	"elegantGo/lession01/model"
	"encoding/csv"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"time"
)

func (ctr *User) ExportV2(c *gin.Context) {
	// 从数据库获取数据
	users := ctr.repo.All()

	// 构建csv内容
	csvContent := ctr.BuildCsv(users)

	// 返回csv文件
	csvName := fmt.Sprintf("users-%s", time.Now().Format("20060102150405"))
	ctr.ToCsv(c, csvName, csvContent)
}

func (ctr *User) BuildCsv(users model.Users) string {
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

	return buff.String()
}

func (ctr *User) ToCsv(c *gin.Context, name string, content string) {
	c.Writer.Header().Set("Content-type", "application/octet-stream")
	c.Writer.Header().Set("Content-Disposition", fmt.Sprintf("attachment;filename=%s.csv", name))
	c.String(http.StatusOK, content)
}
