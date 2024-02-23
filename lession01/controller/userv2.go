package controller

import (
	"bytes"
	"elegantGo/lession01/pkg/util"
	"elegantGo/lession01/repository"
	"encoding/csv"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/xuri/excelize/v2"
	"net/http"
	"strconv"
	"time"
)

type UserV2 struct {
	repo *repository.User
}

func NewUserV2(repo *repository.User) *UserV2 {
	ctr := new(UserV2)
	ctr.repo = repo
	return ctr
}

func (ctr *UserV2) Export(c *gin.Context) {
	users := ctr.repo.All()

	fileFormat := c.DefaultQuery("fileFormat", "excel")
	if fileFormat == "csv" {
		var buffer = new(bytes.Buffer)
		wr := csv.NewWriter(buffer)
		heads := []string{"用户id", "等级", "余额", "手机号码", "昵称", "创建时间"}
		wr.Write(heads)

		sum := 0
		for _, u := range users {
			if u.Status != 10 {
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

				wr.Write([]string{
					strconv.Itoa(u.ID),
					desc,
					strconv.Itoa(u.Balance / 100),
					fmt.Sprintf("%s****%s", mobile[0:4], mobile[8:11]),
					u.Nickname,
					u.CreateTime.Format(time.DateTime),
				})

				sum += u.Balance
			}
		}

		wr.Write([]string{
			fmt.Sprintf("用户总数：%d 用户总余额：%d", len(users), sum/100),
		})
		wr.Flush()

		c.Writer.Header().Set("Content-type", "application/octet-stream")
		c.Writer.Header().Set("Content-Disposition", fmt.Sprintf("attachment;filename=%s.csv", fmt.Sprintf("users-%s", time.Now().Format("20060102150405"))))
		c.String(http.StatusOK, buffer.String())

	} else {
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
}
