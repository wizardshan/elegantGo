package request

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"strconv"
	"strings"
)

var Validate = gin.Validate

type MobileField struct {
	Mobile string `binding:"required,number,mobile"`
}

type CaptchaField struct {
	Captcha string `binding:"required,number,len=4"`
}

type IDField struct {
	ID int `binding:"required,positivenumber"`
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

type QueryField struct {
	Sort     string
	Order    string `binding:"omitempty,oneof=desc asc"`
	Page     int    `binding:"required,min=1"`
	PageSize int    `binding:"required,min=10,max=100"`
}

type TimeRangeField string
type NumberRangeField string
type NumbersBySeparatorField string
type StringsBySeparatorField string

type EqualField string
type LikeField string
type BetweenField string
type InField string

func (req *NumbersBySeparatorField) UnmarshalJSON(b []byte) error {

	var data string
	if err := json.Unmarshal(b, &data); err != nil {
		return err
	}

	if data == "" {
		return nil
	}
	s := req.split(data)
	err := Validate.Var(s, "dive,number")
	if err != nil {
		return err
	}

	*req = NumbersBySeparatorField(data)
	return nil
}

func (req *NumbersBySeparatorField) split(s string) []string {
	return strings.Split(s, ",")
}

func (req *NumbersBySeparatorField) Able() bool {
	if req == nil {
		return false
	}

	if *req == "" {
		return false
	}

	return true
}

func (req *NumbersBySeparatorField) Numbers() []int {
	data := string(*req)
	result := strings.Split(data, ",")
	var numbers []int
	for _, item := range result {
		num, _ := strconv.Atoi(item)
		numbers = append(numbers, num)
	}
	return numbers
}
