package main

import (
	"app/chapter-param-validator-6/controller"
	"fmt"
	"github.com/elliotchance/pie/v2"
	"github.com/gin-gonic/gin"
	"strconv"
	"strings"
)

type User struct {
	ID int
}

func main() {

	u1 := &User{ID: 1}
	u2 := &User{ID: 2}
	u3 := &User{ID: 3}
	uu := []*User{
		u1,
		u2,
		u3,
	}

	elements := strings.Split(",1", ",")

	s := pie.Map(elements, func(value string) *int {
		num, err := strconv.Atoi(value)
		if err != nil {
			return nil
		}
		return &num
	})
	//ss := []int{1, 2, 4, 111, 4, 5, 5, 7}
	//uu2 := pie.Each(uu, func(value *User) {
	//	value.ID = value.ID * 2
	//})
	//pie.Each(uu, func(value *User) {
	//	fmt.Println(value.ID)
	//})
	//
	//pie.Each(uu2, func(value *User) {
	//	fmt.Println(value.ID)
	//})

	//uu2 := make([]User, 2)
	//copy(uu2, uu)
	//uu[0].ID = 111
	//fmt.Println(uu)
	//fmt.Println(uu2)

	//ss := []int{1, 2, 3, 111, 4, 5, 6, 7}
	//fmt.Println(pie.SortUsing(uu, func(a, b User) bool {
	//	return a.ID < b.ID
	//}))
	//
	//ss1 := []int{1, 2, 3}
	//ss2 := []float32{1.1, 2.2, 3.3, 4.4, 5.5}
	//fmt.Println(pie.Zip(ss2, ss1))
	//fmt.Println(pie.ZipLongest(ss2, ss1))

	return
	elements := strings.Split(",1", ",")

	s := pie.Map(elements, func(value string) *int {
		num, err := strconv.Atoi(value)
		if err != nil {
			return nil
		}
		return &num
	})
	fmt.Println(s)

	return

	//for _, e := range elementR {
	//	num, err := strconv.Atoi(e)
	//	if err != nil {
	//		continue
	//	}
	//
	//	numberR = append(numberR, num)
	//}

	//All(elementR, func(value string) (err error) {
	//	var num int
	//	if num, err = strconv.Atoi(value); err != nil {
	//		return
	//	}
	//	numberR = append(numberR, num)
	//	return
	//})
	//
	//fmt.Println(numberR)
	//
	//return
	//r1 := pie.All([]float64{12.3, 12.3}, func(value float64) bool {
	//	return value == 12.3
	//})
	//fmt.Println(r1)
	//r2 := pie.Any([]float64{12.1, 12.3}, func(value float64) bool {
	//	return value == 12.3
	//})

	pie.Ints([]string{"92.384", "823.324", "453"})
	i := pie.Join([]float64{92.384, 823.324, 453}, ",")
	fmt.Println(i)

	return
	engine := gin.New()
	handler := new(controller.Handler)

	ctrUser := controller.NewUser()
	engine.GET("/users", handler.Wrapper(ctrUser.Many))

	engine.Run()
}
