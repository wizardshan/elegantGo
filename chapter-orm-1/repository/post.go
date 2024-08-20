package repository

import (
	"context"
	"database/sql"
	"elegantGo/chapter-orm-1/repository/entity"
	"fmt"
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

	query := fmt.Sprintf("SELECT posts.`id`, posts.`hash_id`, posts.`user_id`, posts.`title`, posts.`content`, posts.`views`, posts.`create_time`, posts.`update_time`, users.`nickname`, users.`avatar` FROM posts, users WHERE posts.user_id=users.id AND posts.`id`= %d", id)
	row := repo.db.QueryRowContext(ctx, query)

	var entPost entity.Post
	row.Scan(
		&entPost.ID,
		&entPost.HashID,
		&entPost.UserID,
		&entPost.Title,
		&entPost.Content,
		&entPost.Views,
		&entPost.CreateTime,
		&entPost.UpdateTime,

		&entPost.Nickname,
		&entPost.Avatar,
	)

	query = fmt.Sprintf("SELECT comments.`id`, comments.`user_id`, comments.`post_id`, comments.`content`, comments.`create_time`, comments.`update_time`, users.`nickname`, users.`avatar` FROM comments, users WHERE comments.user_id=users.id AND comments.`post_id`= %d ORDER BY comments.create_time DESC", id)
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

			&entComment.Nickname,
			&entComment.Avatar,
		)
		entComments = append(entComments, &entComment)
	}
	entPost.Comments = entComments

	return &entPost
}

func (repo *Post) FetchMany(ctx context.Context) []*entity.Post {

	query := "SELECT posts.`id`, posts.`hash_id`, posts.`user_id`, posts.`title`, posts.`content`, posts.`views`, posts.`create_time`, posts.`update_time`, users.`nickname`, users.`avatar` FROM posts, users WHERE posts.user_id=users.id ORDER BY posts.create_time DESC"
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
			&entPost.Views,
			&entPost.CreateTime,
			&entPost.UpdateTime,

			&entPost.Nickname,
			&entPost.Avatar,
		)
		entPosts = append(entPosts, &entPost)
	}

	return entPosts
}
