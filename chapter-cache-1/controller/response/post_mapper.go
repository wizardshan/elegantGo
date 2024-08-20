package response

import (
	"elegantGo/chapter-cache-1/domain"
)

func (respPost *Post) Mapper(domPost *domain.Post) *Post {
	if domPost == nil {
		return nil
	}

	post := new(Post)
	post.HashID = domPost.HashID
	post.UserID = domPost.UserID
	post.Title = domPost.Title
	post.Content = domPost.Content
	post.TimesOfRead = domPost.TimesOfRead
	post.ID = domPost.ID
	post.CreateTime = domPost.CreateTime
	post.UpdateTime = domPost.UpdateTime
	post.Comments = post.Comments.Mapper(domPost.Comments)
	post.User = post.User.Mapper(domPost.User)

	return post
}

func (respPosts Posts) Mapper(domPosts domain.Posts) Posts {
	return mapper(domPosts, func(dom *domain.Post) (post *Post) {
		return post.Mapper(dom)
	})
}
