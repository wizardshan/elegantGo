package controller

import (
	"elegantGo/orm-crud-1/repository"
	"elegantGo/orm-crud-1/repository/ent"
	"elegantGo/orm-crud-1/repository/ent/user"
	"github.com/gin-gonic/gin"
	"net/http"
)

type User struct {
	repo *repository.User
}

func NewUser(repo *repository.User) *User {
	ctr := new(User)
	ctr.repo = repo
	return ctr
}

func (ctr *User) One(c *gin.Context) {
	id := 1
	entUser := ctr.repo.FetchByID(c.Request.Context(), id)
	c.JSON(http.StatusOK, entUser)

	//entUser := ctr.repo.FetchOne(c.Request.Context(), func(opt *ent.UserQuery) {
	//	opt.Where(user.ID(id))
	//})
	//c.JSON(http.StatusOK, entUser)

	// Fetch实现
	//entUser := ctr.repo.Fetch(c.Request.Context(), id)
	//c.JSON(http.StatusOK, entUser)
	//
	// FetchByMobile实现
	//mobile := "13000000001"
	//entUser := ctr.repo.FetchOne(c.Request.Context(), func(opt *ent.UserQuery) {
	//	opt.Where(user.Mobile(mobile))
	//})
	//c.JSON(http.StatusOK, entUser)

	// FetchByMobileAndPassword实现
	//mobile := "13000000001"
	//password := "a906449d5769fa7361d7ecc6aa3f6d28"
	//entUser := ctr.repo.FetchOne(c.Request.Context(), func(opt *ent.UserQuery) {
	//	opt.Where(user.Mobile(mobile), user.Password(password))
	//})
	//c.JSON(http.StatusOK, entUser)

}

func (ctr *User) Many(c *gin.Context) {

	// FetchByNickname实现
	//nickname := "昵称1"
	//entUsers := ctr.repo.FetchMany(c.Request.Context(), func(opt *ent.UserQuery) {
	//	opt.Where(user.Nickname(nickname)).Order(ent.Desc(user.FieldCreateTime))
	//})
	//c.JSON(http.StatusOK, entUsers)

	// 分页实现
	offset := 0
	limit := 10
	optionFunc := func(opt *ent.UserQuery) {
		opt.Offset(offset).Limit(limit).Order(ent.Desc(user.FieldCreateTime))
	}
	entUsers := ctr.repo.FetchMany(c.Request.Context(), optionFunc)
	totalCount := ctr.repo.Count(c.Request.Context(), optionFunc)

	c.JSON(http.StatusOK, gin.H{
		"users":      entUsers,
		"totalCount": totalCount,
	})
}

func (ctr *User) CUD(c *gin.Context) {
	mobile := "13000000001"
	nickname := "昵称1"
	entUser := ctr.repo.Create(c.Request.Context(), func(opt *ent.UserCreate) {
		opt.SetMobile(mobile).SetNickname(nickname)
	})
	c.JSON(http.StatusOK, entUser)

	//updateAffectedRows := ctr.repo.Update(c.Request.Context(), func(opt *ent.UserUpdate) {
	//	opt.SetMobile(mobile).Where(user.Nickname(nickname))
	//})
	//c.JSON(http.StatusOK, updateAffectedRows)
	//
	//deleteAffectedRows := ctr.repo.Delete(c.Request.Context(), func(opt *ent.UserDelete) {
	//	opt.Where(user.Nickname(nickname))
	//})
	//c.JSON(http.StatusOK, deleteAffectedRows)
}

func (ctr *User) Register(c *gin.Context) {
	entUser, err := ctr.repo.Register(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, entUser)
}
