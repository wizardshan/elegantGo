package controller

import (
	"elegantGo/lession01/pkg/util"
	"elegantGo/lession01/repository"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/xuri/excelize/v2"
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

	f := excelize.NewFile()
	defer f.Close()

	name := "Sheet1"
	index, _ := f.NewSheet(name)

	i := 1
	f.SetCellValue(name, fmt.Sprintf("A%d", i), "用户id")
	f.SetCellValue(name, fmt.Sprintf("B%d", i), "等级")
	f.SetCellValue(name, fmt.Sprintf("C%d", i), "余额")
	f.SetCellValue(name, fmt.Sprintf("D%d", i), "手机号码")
	f.SetCellValue(name, fmt.Sprintf("E%d", i), "昵称")
	f.SetCellValue(name, fmt.Sprintf("F%d", i), "创建时间")

	sum := 0
	for _, u := range users {
		if u.Status != 10 {
			i++
			desc := ""
			switch u.Level {
			case 0:
				desc = "普通"
			case 10:
				desc = "白银"
			case 20:
				desc = "黄金"
			case 30:
				desc = "铂金"
			default:
				desc = "未知等级"
			}

			mobile, _ := util.Decrypt(u.Mobile)

			f.SetCellValue(name, fmt.Sprintf("A%d", i), u.ID)
			f.SetCellValue(name, fmt.Sprintf("B%d", i), desc)
			f.SetCellValue(name, fmt.Sprintf("C%d", i), u.Balance/100)
			f.SetCellValue(name, fmt.Sprintf("D%d", i), fmt.Sprintf("%s****%s", mobile[0:4], mobile[8:11]))
			f.SetCellValue(name, fmt.Sprintf("E%d", i), u.Nickname)
			f.SetCellValue(name, fmt.Sprintf("F%d", i), u.CreateTime.Format(time.DateTime))

			sum += u.Balance
		}
	}

	i++
	f.SetCellValue(name, fmt.Sprintf("A%d", i), fmt.Sprintf("用户总数：%d 用户总余额：%d", len(users), sum/100))
	f.SetActiveSheet(index)

	c.Writer.Header().Add("Content-Type", "application/octet-stream")
	c.Writer.Header().Add("Content-disposition", fmt.Sprintf("attachment;filename=users-%s.xlsx", time.Now().Format("20060102150405")))
	c.Writer.Header().Add("Content-Transfer-Encoding", "binary")
	f.Write(c.Writer)
}
