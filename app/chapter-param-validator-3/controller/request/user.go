package request

import (
	"strconv"
	"strings"
)

type UserLogin struct {
	Mobile  string `binding:"required,number,mobile" label:"手机号码"`
	Captcha string `binding:"required,number,len=4" label:"手机验证码"`
}

type UserDelete struct {
	IDS string `binding:"required"`
	ids []int
}

func (req *UserDelete) Validate() error {

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

func (req *UserDelete) GetIDS() []int {
	return req.ids
}
