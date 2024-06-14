package request

import (
	"encoding/json"
	"errors"
	"fmt"
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

type EqualField string
type LikeField string
type BetweenField string
type InField string

type NumberRangeField struct {
	Start     int
	End       int
	startAble bool
	endAble   bool
}

func (req *NumberRangeField) UnmarshalJSON(b []byte) error {

	var data string
	if err := json.Unmarshal(b, &data); err != nil {
		return err
	}

	if data == "" {
		req.startAble = false
		req.endAble = false
		return nil
	}

	numbers := strings.Split(data, ",")
	capacity := len(numbers)
	if capacity != 2 {
		return errors.New(fmt.Sprintf("the rangeField capacity expected value is 2, the result is %d", capacity))
	}

	for i, numStr := range numbers {
		if numStr == "" {
			continue
		}

		num, err := strconv.Atoi(numStr)
		if err != nil {
			return err
		}

		if i == 0 {
			req.startAble = true
			req.Start = num
		} else {
			req.endAble = true
			req.End = num
		}
	}
	return nil
}

func (req *NumberRangeField) StartAble() bool {
	if req == nil {
		return false
	}

	return req.startAble
}

func (req *NumberRangeField) EndAble() bool {
	if req == nil {
		return false
	}

	return req.endAble
}

type StringsBySeparatorField struct {
	Values []string
}

func (req *StringsBySeparatorField) UnmarshalJSON(b []byte) error {

	var data string
	if err := json.Unmarshal(b, &data); err != nil {
		return err
	}

	req.Values = strings.Split(data, ",")

	return nil
}

func (req *StringsBySeparatorField) Able() bool {
	if req == nil || req.Values == nil {
		return false
	}

	if len(req.Values) == 0 {
		return false
	}

	return true
}

type NumbersBySeparatorField struct {
	Values []int
}

func (req *NumbersBySeparatorField) UnmarshalJSON(b []byte) error {

	var data string
	if err := json.Unmarshal(b, &data); err != nil {
		return err
	}

	numbers := strings.Split(data, ",")
	for _, numStr := range numbers {
		num, err := strconv.Atoi(numStr)
		if err != nil {
			return nil
		}
		req.Values = append(req.Values, num)
	}
	return nil
}

func (req *NumbersBySeparatorField) Able() bool {
	if req == nil || req.Values == nil {
		return false
	}

	if len(req.Values) == 0 {
		return false
	}

	return true
}
