package repository

import (
	"context"
	"database/sql"
	"elegantGo/chapter-orm-3/repository/entity"
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

func (repo *Post) FetchByID(ctx context.Context, id int) *entity.Post {

	query := fmt.Sprintf("SELECT posts.`id`, posts.`hash_id`, posts.`user_id`, posts.`title`, posts.`content`, posts.`times_of_read`, posts.`create_time`, posts.`update_time`, users.`nickname`, users.`avatar`, users.`level` FROM posts, users WHERE posts.user_id=users.id AND posts.`id`= %d", id)
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

		&entityPost.UserNickname,
		&entityPost.UserAvatar,
		&entityPost.UserLevel,
	)

	return &entityPost
}

func (repo *Post) FetchMany(ctx context.Context) entity.Posts {

	query := "SELECT posts.`id`, posts.`hash_id`, posts.`user_id`, posts.`title`, posts.`content`, posts.`times_of_read`, posts.`create_time`, posts.`update_time`, users.`nickname`, users.`avatar`, users.`level` FROM posts, users WHERE posts.user_id=users.id ORDER BY posts.create_time DESC"
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

			&entityPost.UserNickname,
			&entityPost.UserAvatar,
			&entityPost.UserLevel,
		)
		entityPosts = append(entityPosts, &entityPost)
	}

	return entityPosts
}

func (repo *Post) Comments(ctx context.Context, postID int) entity.Comments {

	query := fmt.Sprintf("SELECT comments.`id`, comments.`user_id`, comments.`post_id`, comments.`content`, comments.`create_time`, comments.`update_time`, users.`nickname`, users.`avatar`, users.`level` FROM comments, users WHERE comments.user_id=users.id AND comments.`post_id`= %d ORDER BY comments.create_time DESC", postID)
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

			&entityComment.UserNickname,
			&entityComment.UserAvatar,
			&entityComment.UserLevel,
		)
		entityComments = append(entityComments, &entityComment)
	}

	return entityComments
}
