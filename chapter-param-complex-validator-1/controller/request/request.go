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
	ID int `binding:"required,min=1"`
}

type IDsField string

func (req *IDsField) Values() []int {
	ss := strings.Split(string(*req), ",")
	var values []int
	for _, s := range ss {
		value, err := strconv.Atoi(s)
		if err != nil {
			continue
		}
		values = append(values, value)
	}
	return values
}
