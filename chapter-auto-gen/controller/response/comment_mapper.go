package response

import (
    "elegantGo/chapter-auto-gen/domain"
)

func (respComment *Comment) Mapper(domComment *domain.Comment) *Comment {
    if domComment == nil {
        return nil
    }

    comment := new(Comment)
    comment.UserID = domComment.UserID
    comment.PostID = domComment.PostID
    comment.Content = domComment.Content
    comment.ID = domComment.ID
    comment.CreateTime = domComment.CreateTime
    comment.UpdateTime = domComment.UpdateTime
    comment.Post = comment.Post.Mapper(domComment.Post)
    comment.User = comment.User.Mapper(domComment.User)

    return comment
}

func (respComments Comments) Mapper(domComments domain.Comments) Comments {
    return mapper(domComments, func(dom *domain.Comment) (comment *Comment) {
        return comment.Mapper(dom)
    })
}

