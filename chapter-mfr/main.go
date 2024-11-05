package main

import (
	"fmt"
	"github.com/elliotchance/pie/v2"
	"github.com/samber/lo"
)

type User struct {
	ID   int
	Name string
}

type Users []*User

func (us Users) IDs() []int {
	return pie.Map(us, func(user *User) int {
		return user.ID
	})
}

func (us Users) UniqueIDs() []int {
	return pie.Unique(us.IDs())
}

func (us Users) GlueIDs() string {
	return pie.Join(us.UniqueIDs(), ",")
}

//func init() {
//	for i := 1; i <= 10; i++ {
//		u := new(User)
//		u.ID = i
//		u.Name = fmt.Sprintf("name%d", i)
//		users = append(users, u)
//	}
//}

func main() {

	fmt.Println(lo.SnakeCase("CreateTime"))
	return

}

//func Join[T comparable, U any](ss1 []T, fn func(T) U) string {
//	return pie.Join(pie.Unique(pie.Map(ss1, fn)), ",")
//}
//
//func (us Users) join() string {
//	return pie.Join(pie.Unique(pie.Map(us, func(user *User) int {
//		return user.ID
//	})), ",")
//}
