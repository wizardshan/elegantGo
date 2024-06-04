package controller

import (
	"app/chapter-orm-crud-2/repository"
	"app/chapter-orm-crud-2/repository/ent"
	"database/sql"
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
