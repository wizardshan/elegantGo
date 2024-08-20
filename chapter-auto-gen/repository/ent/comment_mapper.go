package ent

import (
	"elegantGo/chapter-auto-gen/domain"
	"github.com/elliotchance/pie/v2"
)

func (entComment *Comment) Mapper() *domain.Comment {
	if entComment == nil {
		return nil
	}

	domComment := new(domain.Comment)
	domComment.UserID = entComment.UserID
	domComment.PostID = entComment.PostID
	domComment.Content = entComment.Content
	domComment.ID = entComment.ID
	domComment.CreateTime = entComment.CreateTime
	domComment.UpdateTime = entComment.UpdateTime
	domComment.Post = entComment.Edges.Post.Mapper()
	domComment.User = entComment.Edges.User.Mapper()

	return domComment
}

func (entComments Comments) Mapper() domain.Comments {
	return pie.Map(entComments, func(ent *Comment) *domain.Comment {
		return ent.Mapper()
	})
}