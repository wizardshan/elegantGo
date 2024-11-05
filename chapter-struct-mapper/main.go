package main

import (
	"elegantGo/chapter-struct-mapper/controller"
	"elegantGo/chapter-struct-mapper/repository"
	"fmt"
	"github.com/elliotchance/pie/v2"
	"github.com/gin-gonic/gin"
	"math"
)

type User struct {
	ID   int
	Name string
}

type UserT struct {
	ID   int
	Name string
}

func main() {

	v := []float64{1.1, 2.2, 3.3, 4.4, 5.5}

	avg := pie.Average(v)
	count := len(v)

	sum2 := pie.Reduce(
		v,
		func(acc, value float64) float64 {
			return acc + (value-avg)*(value-avg)
		},
	) - v[0] + (v[0]-avg)*(v[0]-avg)

	d := math.Sqrt(sum2 / float64(count))

	fmt.Printf("Standard deviation: %f\n", d)

	var users []*User
	for i := 0; i < 10; i++ {
		u := new(User)
		u.ID = i
		users = append(users, u)
	}

	var usersT []*UserT
	for i := 0; i < 10; i++ {
		u := new(UserT)
		u.ID = i
		u.Name = fmt.Sprintf("name%d", i)
		usersT = append(usersT, u)
	}

	pie.Join(pie.Map(users, func(user *User) int {
		return user.ID
	}), ",")

	//strconv.Itoa(int)

	return
	engine := gin.New()

	handler := new(controller.Handler)

	repoArticle := repository.NewArticle()
	ctrArticle := controller.NewArticle(repoArticle)
	engine.GET("/article", handler.Wrapper(ctrArticle.One))
	engine.GET("/articles", handler.Wrapper(ctrArticle.Many))

	engine.Run()
}
