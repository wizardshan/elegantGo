package controller

import (
	"database/sql"
	"elegantGo/chapter-orm-crud-3/repository"
	"elegantGo/chapter-orm-crud-3/repository/ent"
	entsql "entgo.io/ent/dialect/sql"
	"errors"
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
	query := "select `id`, `hash_id`, `mobile`, `nickname`, `create_time`, `update_time` from users where `id`=?"

	var entUser ent.User
	err := ctr.repo.QueryOne(c.Request.Context(), query, []any{id}, func(row *sql.Row) error {
		return row.Scan(
			&entUser.ID,
			&entUser.HashID,
			&entUser.Mobile,
			&entUser.Nickname,
			&entUser.CreateTime,
			&entUser.UpdateTime,
		)
	})

	if errors.Is(err, sql.ErrNoRows) {
		c.JSON(http.StatusOK, gin.H{
			"user": nil,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"user": entUser,
	})
}

func (ctr *User) Many(c *gin.Context) {

	nickname := "昵称1"
	query := "select `id`, `hash_id`, `mobile`, `nickname`, `create_time`, `update_time` from users where `nickname`=?"
	var entUsers ent.Users
	err := ctr.repo.QueryMany(c.Request.Context(), query, []any{nickname}, func(rows *sql.Rows) error {
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
			return err
		}

		entUsers = append(entUsers, &entUser)

		return nil
	})

	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"users": entUsers,
	})
}

func (ctr *User) Register(c *gin.Context) {
	entUser, err := ctr.repo.Register(c.Request.Context())

	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"user": entUser,
	})
}

func (ctr *User) Upsert(c *gin.Context) {

	mobile := "13000000003"
	password := "a906449d5769fa7361d7ecc6aa3f6d28"
	level := 30
	nickname := "昵称3"
	avatar := "头像3.png"
	bio := "个人介绍3"

	entUser := ctr.repo.Create(c.Request.Context(), func(builder *ent.UserCreate) {
		builder.
			SetMobile(mobile).
			SetPassword(password).
			SetLevel(level).
			SetNickname(nickname).
			SetAvatar(avatar).
			SetBio(bio).
			OnConflict().
			UpdateNewValues().
			Update(func(u *ent.UserUpsert) {
				u.SetBio("个人介绍33")
			})
	})

	c.JSON(http.StatusOK, gin.H{
		"user": entUser,
	})
}

func (ctr *User) Rand(c *gin.Context) {
	entUsers := ctr.repo.FetchMany(c.Request.Context(), func(builder *ent.UserQuery) {
		builder.Order(func(s *entsql.Selector) {
			s.OrderBy("rand()")
		})
	})

	c.JSON(http.StatusOK, gin.H{
		"users": entUsers,
	})
}
