package controller

import (
	"elegantGo/chapter-param-complex-validator-4/controller/request"
	"elegantGo/chapter-param-complex-validator-4/controller/response"
	"elegantGo/chapter-param-complex-validator-4/repository"
	"elegantGo/chapter-param-complex-validator-4/repository/ent"
	"elegantGo/chapter-param-complex-validator-4/repository/ent/user"
	"github.com/gin-gonic/gin"
)

type User struct {
	repo *repository.User
}

func NewUser(repo *repository.User) *User {
	ctr := new(User)
	ctr.repo = repo
	return ctr
}

func (ctr *User) Many(c *gin.Context) (response.Data, error) {
	request := new(request.UserMany)
	if err := c.ShouldBind(request); err != nil {
		return nil, err
	}

	ctr.repo.FetchMany(c.Request.Context(), func(builder *ent.UserQuery) {
		if request.Filter.ID != nil {
			builder = builder.Where(user.ID(*request.Filter.ID))
		}

		if request.Filter.Nickname != nil {
			builder = builder.Where(user.NicknameContains(*request.Filter.Nickname))
		}

		if request.Filter.Amount.StartAble() {
			builder = builder.Where(user.AmountGTE(*request.Filter.Amount.Start))
		}

		if request.Filter.Amount.EndAble() {
			builder = builder.Where(user.AmountLTE(*request.Filter.Amount.End))
		}

		if request.Filter.Status.Able() {
			builder = builder.Where(user.StatusIn(request.Filter.Status...))
		}

		if request.Filter.Level.Able() {
			builder = builder.Where(user.LevelIn(request.Filter.Level...))
		}

		if request.Filter.CreateTime.StartAble() {
			builder = builder.Where(user.CreateTimeGTE(request.Filter.CreateTime.Start))
		}

		if request.Filter.CreateTime.EndAble() {
			builder = builder.Where(user.CreateTimeLTE(request.Filter.CreateTime.End))
		}

		builder.Offset(request.Offset.Value()).Limit(request.Limit.Value()).Order(request.Order.By(request.Sort.Value()))
	})

	return request, nil
}
