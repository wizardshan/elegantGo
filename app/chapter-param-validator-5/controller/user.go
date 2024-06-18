package controller

import (
	"app/chapter-param-validator-5/controller/request"
	"app/chapter-param-validator-5/controller/response"
	"app/chapter-param-validator-5/repository"
	"app/chapter-param-validator-5/repository/ent"
	"app/chapter-param-validator-5/repository/ent/user"
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
	if err := c.Validate(request); err != nil {
		return nil, err
	}

	entUsers := ctr.repo.FetchMany(c.Request.Context(), func(builder *ent.UserQuery) {
		if request.Filter.ID != nil {
			builder = builder.Where(user.ID(*request.Filter.ID))
		}

		if request.Filter.Nickname != nil {
			builder = builder.Where(user.NicknameContains(*request.Filter.Nickname))
		}

		if request.Filter.Amount.StartAble() {
			builder = builder.Where(user.AmountGTE(request.Filter.Amount.Start))
		}

		if request.Filter.Amount.EndAble() {
			builder = builder.Where(user.AmountLTE(request.Filter.Amount.End))
		}

		if request.Filter.Status.Able() {
			builder = builder.Where(user.StatusIn(request.Filter.Status.Values...))
		}

		if request.Filter.Level.Able() {
			builder = builder.Where(user.LevelIn(request.Filter.Level.Values...))
		}

		if request.Filter.CreateTime.StartAble() {
			builder = builder.Where(user.CreateTimeGTE(request.Filter.CreateTime.Start))
		}

		if request.Filter.CreateTime.EndAble() {
			builder = builder.Where(user.CreateTimeLTE(request.Filter.CreateTime.End))
		}

		page := request.Page.Value()
		pageSize := request.PageSize.Value()
		builder.Offset(pageSize * (page - 1)).Limit(pageSize)

		order := ent.Asc
		if request.Order.IsDesc() {
			order = ent.Desc
		}
		builder.Order(order(request.Sort.Value()))
	})

	//return request, nil
	return entUsers, nil
}
