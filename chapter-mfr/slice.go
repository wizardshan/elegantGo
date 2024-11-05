package main

import (
	"github.com/elliotchance/pie/v2"
	"github.com/kr/pretty"
	"github.com/samber/lo"
	"strconv"
)

func FilterMap() {
	ids := []string{"1", "", "3", "4"}
	ids1 := lo.FilterMap(ids, func(s string, index int) (int, bool) {
		if !IsInt(s) {
			return 0, false
		}
		num, _ := strconv.Atoi(s)
		return num, true
	})

	ids2 := pie.Ints(pie.Filter(ids, func(s string) bool {
		return IsInt(s)
	}))

	pretty.Println(ids1, ids2)
}

func Map() {
	ids := []string{"1", "", "3", "4"}
	ids1 := lo.Map(ids, func(s string, index int) int {
		num, _ := strconv.Atoi(s)
		return num
	})

	ids2 := pie.Map(ids, func(s string) int {
		num, _ := strconv.Atoi(s)
		return num
	})

	pretty.Println(ids1, ids2)
}

func Filter(users []*User) {
	u1 := lo.Filter(users, func(u *User, index int) bool {
		return u.ID%2 == 0
	})

	u2 := pie.Filter(users, func(u *User) bool {
		return u.ID%2 == 0
	})

	pretty.Println(u1, u2)
}
