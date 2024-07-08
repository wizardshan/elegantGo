package repository

import (
	"context"
	"database/sql"
	"elegantGo/chapter-orm-4/repository/entity"
	"fmt"
	"strconv"
	"strings"
)

type Post struct {
	db *sql.DB
}

func NewPost(db *sql.DB) *Post {
	repo := new(Post)
	repo.db = db
	return repo
}

func (repo *Post) FetchByID(ctx context.Context, id int) *entity.Post {

	query := fmt.Sprintf("SELECT `id`, `hash_id`, `user_id`, `title`, `content`, `times_of_read`, `create_time`, `update_time` FROM posts WHERE `id`= %d", id)
	row := repo.db.QueryRowContext(ctx, query)

	var entityPost entity.Post
	row.Scan(
		&entityPost.ID,
		&entityPost.HashID,
		&entityPost.UserID,
		&entityPost.Title,
		&entityPost.Content,
		&entityPost.TimesOfRead,
		&entityPost.CreateTime,
		&entityPost.UpdateTime,
	)

	query = fmt.Sprintf("SELECT `id`, `mobile`, `password`, `nickname`, `avatar`, `bio`, `create_time`, `update_time` FROM users WHERE `id` = %d", entityPost.UserID)
	row = repo.db.QueryRowContext(ctx, query)

	var entityUser entity.User
	row.Scan(
		&entityUser.ID,
		&entityUser.Mobile,
		&entityUser.Password,
		&entityUser.Nickname,
		&entityUser.Avatar,
		&entityUser.Bio,
		&entityUser.CreateTime,
		&entityUser.UpdateTime,
	)
	entityPost.User = &entityUser

	return &entityPost
}

func (repo *Post) FetchMany(ctx context.Context) entity.Posts {

	query := "SELECT `id`, `hash_id`, `user_id`, `title`, `content`, `times_of_read`, `create_time`, `update_time` FROM posts ORDER BY create_time DESC"
	rows, _ := repo.db.QueryContext(ctx, query)
	defer rows.Close()

	var entityPosts entity.Posts
	for rows.Next() {
		var entityPost entity.Post
		rows.Scan(
			&entityPost.ID,
			&entityPost.HashID,
			&entityPost.UserID,
			&entityPost.Title,
			&entityPost.Content,
			&entityPost.TimesOfRead,
			&entityPost.CreateTime,
			&entityPost.UpdateTime,
		)
		entityPosts = append(entityPosts, &entityPost)
	}

	var userIDs []string
	for _, entityPost := range entityPosts {
		userIDs = append(userIDs, strconv.Itoa(entityPost.UserID))
	}

	query = fmt.Sprintf("SELECT `id`, `mobile`, `password`, `nickname`, `avatar`, `bio`, `create_time`, `update_time` FROM users WHERE `id` IN (%s)", strings.Join(userIDs, ","))
	rows, _ = repo.db.QueryContext(ctx, query)
	defer rows.Close()

	var entityUsers entity.Users
	for rows.Next() {
		var entityUser entity.User
		rows.Scan(
			&entityUser.ID,
			&entityUser.Mobile,
			&entityUser.Password,
			&entityUser.Nickname,
			&entityUser.Avatar,
			&entityUser.Bio,
			&entityUser.CreateTime,
			&entityUser.UpdateTime,
		)
		entityUsers = append(entityUsers, &entityUser)
	}

	entityUserIDMapping := make(map[int]*entity.User, len(entityUsers))
	for _, entityUser := range entityUsers {
		entityUserIDMapping[entityUser.ID] = entityUser
	}

	for _, entityPost := range entityPosts {
		entityPost.User = entityUserIDMapping[entityPost.UserID]
	}

	return entityPosts
}

func (repo *Post) Comments(ctx context.Context, postID int) entity.Comments {

	query := fmt.Sprintf("SELECT `id`, `user_id`, `post_id`, `content`, `create_time`, `update_time` FROM comments WHERE `post_id`= %d ORDER BY create_time DESC", postID)
	rows, _ := repo.db.QueryContext(ctx, query)
	defer rows.Close()

	var entityComments entity.Comments
	for rows.Next() {
		var entityComment entity.Comment
		rows.Scan(
			&entityComment.ID,
			&entityComment.UserID,
			&entityComment.PostID,
			&entityComment.Content,
			&entityComment.CreateTime,
			&entityComment.UpdateTime,
		)
		entityComments = append(entityComments, &entityComment)
	}

	var userIDs []string
	for _, entityComment := range entityComments {
		userIDs = append(userIDs, strconv.Itoa(entityComment.UserID))
	}

	query = fmt.Sprintf("SELECT `id`, `mobile`, `password`, `nickname`, `avatar`, `bio`, `create_time`, `update_time` FROM users WHERE `id` IN (%s)", strings.Join(userIDs, ","))
	rows, _ = repo.db.QueryContext(ctx, query)
	defer rows.Close()

	var entityUsers entity.Users
	for rows.Next() {
		var entityUser entity.User
		rows.Scan(
			&entityUser.ID,
			&entityUser.Mobile,
			&entityUser.Password,
			&entityUser.Nickname,
			&entityUser.Avatar,
			&entityUser.Bio,
			&entityUser.CreateTime,
			&entityUser.UpdateTime,
		)
		entityUsers = append(entityUsers, &entityUser)
	}

	entityUserIDMapping := make(map[int]*entity.User, len(entityUsers))
	for _, entityUser := range entityUsers {
		entityUserIDMapping[entityUser.ID] = entityUser
	}

	for _, entityComment := range entityComments {
		entityComment.User = entityUserIDMapping[entityComment.UserID]
	}

	return entityComments
}

func (repo *Post) LatestComments(ctx context.Context) entity.Comments {

	query := "SELECT `id`, `user_id`, `post_id`, `content`, `create_time`, `update_time` FROM comments ORDER BY create_time DESC LIMIT 10"
	rows, _ := repo.db.QueryContext(ctx, query)
	defer rows.Close()

	var entityComments entity.Comments
	for rows.Next() {
		var entityComment entity.Comment
		rows.Scan(
			&entityComment.ID,
			&entityComment.UserID,
			&entityComment.PostID,
			&entityComment.Content,
			&entityComment.CreateTime,
			&entityComment.UpdateTime,
		)
		entityComments = append(entityComments, &entityComment)
	}

	var userIDs []string
	for _, entityComment := range entityComments {
		userIDs = append(userIDs, strconv.Itoa(entityComment.UserID))
	}

	query = fmt.Sprintf("SELECT `id`, `mobile`, `password`, `nickname`, `avatar`, `bio`, `create_time`, `update_time` FROM users WHERE `id` IN (%s)", strings.Join(userIDs, ","))
	rows, _ = repo.db.QueryContext(ctx, query)
	defer rows.Close()

	var entityUsers entity.Users
	for rows.Next() {
		var entityUser entity.User
		rows.Scan(
			&entityUser.ID,
			&entityUser.Mobile,
			&entityUser.Password,
			&entityUser.Nickname,
			&entityUser.Avatar,
			&entityUser.Bio,
			&entityUser.CreateTime,
			&entityUser.UpdateTime,
		)
		entityUsers = append(entityUsers, &entityUser)
	}

	entityUserIDMapping := make(map[int]*entity.User, len(entityUsers))
	for _, entityUser := range entityUsers {
		entityUserIDMapping[entityUser.ID] = entityUser
	}

	var postIDs []string
	for _, entityComment := range entityComments {
		postIDs = append(postIDs, strconv.Itoa(entityComment.PostID))
	}

	query = fmt.Sprintf("SELECT `id`, `hash_id`, `user_id`, `title`, `content`, `times_of_read`, `create_time`, `update_time` FROM posts WHERE `id` IN (%s)", strings.Join(postIDs, ","))
	rows, _ = repo.db.QueryContext(ctx, query)
	defer rows.Close()

	var entityPosts entity.Posts
	for rows.Next() {
		var entityPost entity.Post
		rows.Scan(
			&entityPost.ID,
			&entityPost.HashID,
			&entityPost.UserID,
			&entityPost.Title,
			&entityPost.Content,
			&entityPost.TimesOfRead,
			&entityPost.CreateTime,
			&entityPost.UpdateTime,
		)
		entityPosts = append(entityPosts, &entityPost)
	}

	entityPostIDMapping := make(map[int]*entity.Post, len(entityPosts))
	for _, entityPost := range entityPosts {
		entityPostIDMapping[entityPost.ID] = entityPost
	}

	for _, entityComment := range entityComments {
		entityComment.Post = entityPostIDMapping[entityComment.PostID]
		entityComment.User = entityUserIDMapping[entityComment.UserID]
	}

	return entityComments
}
