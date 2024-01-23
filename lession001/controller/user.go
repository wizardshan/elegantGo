package controller

import (
	"bytes"
	"elegantGo/lession02/repository"
	"encoding/csv"
	"fmt"
	"github.com/gin-gonic/gin"
	"strconv"
)

type User struct {
	repo *repository.User
}

func NewUser(repo *repository.User) *User {
	ctr := new(User)
	ctr.repo = repo
	return ctr
}

func (ctr *User) Export(c *gin.Context) {
	users := ctr.repo.All()

	//f := excelize.NewFile()
	//
	//// 在默认的 Sheet1 上写入数据
	//err := f.SetCellValue("Sheet1", "A1", "姓名")
	//if err != nil {
	//	log.Fatal(err)
	//}
	//f.SetCellValue("Sheet1", "B1", "年龄")
	//f.SetCellValue("Sheet1", "A2", "张三")
	//f.SetCellValue("Sheet1", "B2", 30)
	//
	//var buffer bytes.Buffer
	//_ = f.Write(&buffer)
	//content := bytes.NewReader(buffer.Bytes())
	//
	//fileName := fmt.Sprintf("%s%s%s.xlsx", time.Now().Format("2016-01-02"), `-`, "user")
	//c.Writer.Header().Add("Content-Disposition", fmt.Sprintf(`attachment; filename="%s"`, fileName))
	//c.Writer.Header().Add("Content-Type", "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet")
	//http.ServeContent(c.Writer, c.Request, fileName, time.Now(), content)
	//
	//return

	var heads []string
	var buff = new(bytes.Buffer)
	//写入UTF-8 BOM，否则windows汉字乱码
	//buff.WriteString("\xEF\xBB\xBF")
	heads = []string{"id", "level", "mobile", "nickname", "createTime"}
	wr := csv.NewWriter(buff)
	wr.Write(heads)
	for _, u := range users {
		wr.Write([]string{
			strconv.Itoa(u.ID),
			strconv.Itoa(u.Level),
			u.Mobile,
			u.Nickname,
			u.CreateTime.GoString(),
		})
	}

	wr.Flush()
	c.Writer.Header().Set("Content-type", "application/octet-stream")
	c.Writer.Header().Set("Content-Disposition", fmt.Sprintf("attachment;filename=%s.csv", "user"))
	c.String(200, buff.String())
}
