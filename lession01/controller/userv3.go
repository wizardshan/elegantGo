package controller

import (
	"elegantGo/lession01/model"
	"elegantGo/lession01/pkg/util"
	"elegantGo/lession01/repository"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/xuri/excelize/v2"
	"time"
)

type UserV3 struct {
	repo *repository.User
}

func NewUserV3(repo *repository.User) *UserV3 {
	ctr := new(UserV3)
	ctr.repo = repo
	return ctr
}

func (ctr *UserV3) Export(c *gin.Context) {
	// 从数据库获取数据
	users := ctr.repo.All()

	excelFile := excelize.NewFile()
	defer excelFile.Close()

	// 构建excel文件内容
	ctr.BuildExcel(users, excelFile)

	// 返回excel文件
	excelName := fmt.Sprintf("users-%s", time.Now().Format("20060102150405"))
	ctr.ToExcel(c, excelName, excelFile)
}

func (ctr *UserV3) BuildExcel(users model.Users, excelFile *excelize.File) *excelize.File {
	sheetName := "Sheet1"
	sheetIndex, _ := excelFile.NewSheet(sheetName)

	cellIndex := 1
	excelFile.SetCellValue(sheetName, fmt.Sprintf("A%d", cellIndex), "用户id")
	excelFile.SetCellValue(sheetName, fmt.Sprintf("B%d", cellIndex), "等级")
	excelFile.SetCellValue(sheetName, fmt.Sprintf("C%d", cellIndex), "余额")
	excelFile.SetCellValue(sheetName, fmt.Sprintf("D%d", cellIndex), "手机号码")
	excelFile.SetCellValue(sheetName, fmt.Sprintf("E%d", cellIndex), "昵称")
	excelFile.SetCellValue(sheetName, fmt.Sprintf("F%d", cellIndex), "创建时间")

	balanceSum := 0
	for _, u := range users {
		if u.Status == model.UserStatusCanceled {
			continue
		}

		cellIndex++
		levelDesc := ""
		switch u.Level {
		case model.UserLevelDescNormal:
			levelDesc = "普通"
		case model.UserLevelDescSilver:
			levelDesc = "白银"
		case model.UserLevelDescGold:
			levelDesc = "黄金"
		case model.UserLevelDescPlatinum:
			levelDesc = "铂金"
		default:
			levelDesc = "未知等级"
		}

		mobile, _ := util.Decrypt(u.Mobile)

		excelFile.SetCellValue(sheetName, fmt.Sprintf("A%d", cellIndex), u.ID)
		excelFile.SetCellValue(sheetName, fmt.Sprintf("B%d", cellIndex), levelDesc)
		excelFile.SetCellValue(sheetName, fmt.Sprintf("C%d", cellIndex), u.Balance/100)
		excelFile.SetCellValue(sheetName, fmt.Sprintf("D%d", cellIndex), fmt.Sprintf("%s****%s", mobile[0:4], mobile[8:11]))
		excelFile.SetCellValue(sheetName, fmt.Sprintf("E%d", cellIndex), u.Nickname)
		excelFile.SetCellValue(sheetName, fmt.Sprintf("F%d", cellIndex), u.CreateTime.Format(time.DateTime))

		balanceSum += u.Balance
	}

	cellIndex++
	excelFile.SetCellValue(sheetName, fmt.Sprintf("A%d", cellIndex), fmt.Sprintf("用户总数：%d 用户总余额：%d", len(users), balanceSum/100))
	excelFile.SetActiveSheet(sheetIndex)

	return excelFile
}

func (ctr *UserV3) ToExcel(c *gin.Context, name string, excelFile *excelize.File) {
	c.Writer.Header().Add("Content-Type", "application/octet-stream")
	c.Writer.Header().Add("Content-disposition", fmt.Sprintf("attachment;filename=%s.xlsx", name))
	c.Writer.Header().Add("Content-Transfer-Encoding", "binary")
	excelFile.Write(c.Writer)
}
