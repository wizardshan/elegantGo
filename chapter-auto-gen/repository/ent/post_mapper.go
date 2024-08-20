package ent

import (
	"elegantGo/chapter-auto-gen/domain"
	"github.com/elliotchance/pie/v2"
)

func (entPost *Post) Mapper() *domain.Post {
	if entPost == nil {
		return nil
	}

	domPost := new(domain.Post)
	domPost.HashID = entPost.HashID
	domPost.UserID = entPost.UserID
	domPost.Title = entPost.Title
	domPost.Content = entPost.Content
	domPost.TimesOfRead = entPost.TimesOfRead
	domPost.ID = entPost.ID
	domPost.CreateTime = entPost.CreateTime
	domPost.UpdateTime = entPost.UpdateTime
	
	domPost.Comments = Comments(entPost.Edges.Comments).Mapper()
	domPost.User = entPost.Edges.User.Mapper()

	return domPost
}

func (entPosts Posts) Mapper() domain.Posts {
	return pie.Map(entPosts, func(ent *Post) *domain.Post {
		return ent.Mapper()
	})
}