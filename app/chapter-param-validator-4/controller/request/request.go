package request

import (
	"strconv"
	"strings"
)

type MobileField struct {
	Mobile string `binding:"required,number,mobile"`
}

type CaptchaField struct {
	Captcha string `binding:"required,number,len=4"`
}

type IDField struct {
	ID string `binding:"required,number"`
}

type IDSField struct {
	IDS string `binding:"required"`
	ids []int
}

func (req *IDSField) Validate() error {

	idArray := strings.Split(req.IDS, ",")
	for _, str := range idArray {
		i, err := strconv.Atoi(str)
		if err != nil {
			return err
		}

		req.ids = append(req.ids, i)
	}
	return nil
}

func (req *IDSField) GetIDS() []int {
	return req.ids
}

type Query struct {
	Sort     string
	Order    string
	Page     int
	PageSize int
	Filter   string
	Operator string
}

func (req *Query) Validate() error {
	return nil
}
