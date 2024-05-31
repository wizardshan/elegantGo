package controller

import (
	"app/chapter-orm-crud-2/repository"
	"app/chapter-orm-crud-2/repository/ent"
	"database/sql"
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
	//id := 1
	//query := "select hash_id, title, content from articles where `id`=%d"
	//ctr.repo.QueryOne(c.Request.Context(), query, []any{id})
	//var entUser ent.User
	//if row.Err() != nil {
	//	article.Err = row.Err().Error()
	//}
	//row.Scan(&article.HashID, &article.Title, &article.Content)
	//return &article

	//entUser := ctr.repo.FetchByID(c.Request.Context(), id)
	//c.JSON(http.StatusOK, entUser)

	//c.JSON(http.StatusOK, entUser)

}

func (ctr *User) Many(c *gin.Context) {

	nickname := "昵称1"
	query := "select `id`, `hash_id`, `mobile`, `nickname`, `create_time`, `update_time` from users where `nickname`=?"
	var entUsers ent.Users
	err := ctr.repo.QueryMany(c.Request.Context(), query, []any{nickname}, func(rows *sql.Rows) {
		var entUser ent.User
		err := rows.Scan(
			&entUser.ID,
			&entUser.HashID,
			&entUser.Mobile,
			&entUser.Nickname,
			&entUser.CreateTime,
			&entUser.UpdateTime,
		)
		if err != nil {
			panic(err.Error())
		}

		entUsers = append(entUsers, &entUser)
	})

	c.JSON(http.StatusOK, gin.H{
		"users": entUsers,
		"error": err,
	})
}

func (ctr *User) Register(c *gin.Context) {
	entUser, err := ctr.repo.Register(c.Request.Context())
	c.JSON(http.StatusOK, gin.H{
		"user":  entUser,
		"error": err,
	})
}
