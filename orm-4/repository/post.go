package repository

import (
	"context"
	"database/sql"
	"elegantGo/orm-4/repository/entity"
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

func (repo *Post) Fetch(ctx context.Context, id int) *entity.Post {

	query := fmt.Sprintf("SELECT `id`, `hash_id`, `user_id`, `title`, `content`, `view`, `create_time`, `update_time` FROM posts WHERE `id`= %d", id)
	row := repo.db.QueryRowContext(ctx, query)

	var entPost entity.Post
	row.Scan(
		&entPost.ID,
		&entPost.HashID,
		&entPost.UserID,
		&entPost.Title,
		&entPost.Content,
		&entPost.View,
		&entPost.CreateTime,
		&entPost.UpdateTime,
	)

	query = fmt.Sprintf("SELECT `id`, `nickname`, `avatar` FROM users WHERE `id` = %d", entPost.UserID)
	row = repo.db.QueryRowContext(ctx, query)

	var entUser entity.User
	row.Scan(
		&entUser.ID,
		&entUser.Nickname,
		&entUser.Avatar,
	)
	entPost.User = &entUser

	query = fmt.Sprintf("SELECT `id`, `user_id`, `post_id`, `content`, `create_time`, `update_time` FROM comments WHERE `post_id`= %d ORDER BY create_time DESC", id)
	rows, _ := repo.db.QueryContext(ctx, query)
	defer rows.Close()

	var entComments []*entity.Comment
	for rows.Next() {
		var entComment entity.Comment
		rows.Scan(
			&entComment.ID,
			&entComment.UserID,
			&entComment.PostID,
			&entComment.Content,
			&entComment.CreateTime,
			&entComment.UpdateTime,
		)
		entComments = append(entComments, &entComment)
	}

	var userIDs []string
	for _, entComment := range entComments {
		userIDs = append(userIDs, strconv.Itoa(entComment.UserID))
	}

	query = fmt.Sprintf("SELECT `id`, `nickname`, `avatar` FROM users WHERE `id` IN (%s)", strings.Join(userIDs, ","))
	rows, _ = repo.db.QueryContext(ctx, query)
	defer rows.Close()

	var entUsers []*entity.User
	for rows.Next() {
		var u entity.User
		rows.Scan(
			&u.ID,
			&u.Nickname,
			&u.Avatar,
		)
		entUsers = append(entUsers, &u)
	}

	entUserIDMapping := make(map[int]*entity.User, len(entUsers))
	for _, u := range entUsers {
		entUserIDMapping[u.ID] = u
	}

	for _, entityComment := range entComments {
		entityComment.User = entUserIDMapping[entityComment.UserID]
	}

	entPost.Comments = entComments

	return &entPost
}

func (repo *Post) FetchMany(ctx context.Context) []*entity.Post {

	query := "SELECT `id`, `hash_id`, `user_id`, `title`, `content`, `view`, `create_time`, `update_time` FROM posts ORDER BY create_time DESC"
	rows, _ := repo.db.QueryContext(ctx, query)
	defer rows.Close()

	var entPosts []*entity.Post
	for rows.Next() {
		var entPost entity.Post
		rows.Scan(
			&entPost.ID,
			&entPost.HashID,
			&entPost.UserID,
			&entPost.Title,
			&entPost.Content,
			&entPost.View,
			&entPost.CreateTime,
			&entPost.UpdateTime,
		)
		entPosts = append(entPosts, &entPost)
	}

	var userIDs []string
	for _, entPost := range entPosts {
		userIDs = append(userIDs, strconv.Itoa(entPost.UserID))
	}

	query = fmt.Sprintf("SELECT `id`, `nickname`, `avatar` FROM users WHERE `id` IN (%s)", strings.Join(userIDs, ","))
	rows, _ = repo.db.QueryContext(ctx, query)
	defer rows.Close()

	var entUsers []*entity.User
	for rows.Next() {
		var entUser entity.User
		rows.Scan(
			&entUser.ID,
			&entUser.Nickname,
			&entUser.Avatar,
		)
		entUsers = append(entUsers, &entUser)
	}

	entUserIDMapping := make(map[int]*entity.User, len(entUsers))
	for _, entUser := range entUsers {
		entUserIDMapping[entUser.ID] = entUser
	}

	for _, entPost := range entPosts {
		entPost.User = entUserIDMapping[entPost.UserID]
	}

	return entPosts
}

func (repo *Post) Comments(ctx context.Context) []*entity.Comment {

	query := "SELECT `id`, `user_id`, `post_id`, `content`, `create_time`, `update_time` FROM comments ORDER BY create_time DESC LIMIT 10"
	rows, _ := repo.db.QueryContext(ctx, query)
	defer rows.Close()

	var entComments []*entity.Comment
	for rows.Next() {
		var entComment entity.Comment
		rows.Scan(
			&entComment.ID,
			&entComment.UserID,
			&entComment.PostID,
			&entComment.Content,
			&entComment.CreateTime,
			&entComment.UpdateTime,
		)
		entComments = append(entComments, &entComment)
	}

	var userIDs []string
	for _, entComment := range entComments {
		userIDs = append(userIDs, strconv.Itoa(entComment.UserID))
	}

	query = fmt.Sprintf("SELECT `id`, `nickname`, `avatar` FROM users WHERE `id` IN (%s)", strings.Join(userIDs, ","))
	rows, _ = repo.db.QueryContext(ctx, query)
	defer rows.Close()

	var entUsers []*entity.User
	for rows.Next() {
		var entUser entity.User
		rows.Scan(
			&entUser.ID,
			&entUser.Nickname,
			&entUser.Avatar,
		)
		entUsers = append(entUsers, &entUser)
	}

	entUserIDMapping := make(map[int]*entity.User, len(entUsers))
	for _, entityUser := range entUsers {
		entUserIDMapping[entityUser.ID] = entityUser
	}

	var postIDs []string
	for _, entComment := range entComments {
		postIDs = append(postIDs, strconv.Itoa(entComment.PostID))
	}

	query = fmt.Sprintf("SELECT `id`, `hash_id`, `user_id`, `title`, `content`, `view`, `create_time`, `update_time` FROM posts WHERE `id` IN (%s)", strings.Join(postIDs, ","))
	rows, _ = repo.db.QueryContext(ctx, query)
	defer rows.Close()

	var entPosts []*entity.Post
	for rows.Next() {
		var entPost entity.Post
		rows.Scan(
			&entPost.ID,
			&entPost.HashID,
			&entPost.UserID,
			&entPost.Title,
			&entPost.Content,
			&entPost.View,
			&entPost.CreateTime,
			&entPost.UpdateTime,
		)
		entPosts = append(entPosts, &entPost)
	}

	entPostIDMapping := make(map[int]*entity.Post, len(entPosts))
	for _, entPost := range entPosts {
		entPostIDMapping[entPost.ID] = entPost
	}

	for _, entComment := range entComments {
		entComment.Post = entPostIDMapping[entComment.PostID]
		entComment.User = entUserIDMapping[entComment.UserID]
	}

	return entComments
}
