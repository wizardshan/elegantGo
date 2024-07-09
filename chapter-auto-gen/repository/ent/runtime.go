// Code generated by ent, DO NOT EDIT.

package ent

import (
	"elegantGo/chapter-auto-gen/repository/ent/comment"
	"elegantGo/chapter-auto-gen/repository/ent/post"
	"elegantGo/chapter-auto-gen/repository/ent/schema"
	"elegantGo/chapter-auto-gen/repository/ent/user"
	"time"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	commentMixin := schema.Comment{}.Mixin()
	commentMixinFields0 := commentMixin[0].Fields()
	_ = commentMixinFields0
	commentFields := schema.Comment{}.Fields()
	_ = commentFields
	// commentDescCreateTime is the schema descriptor for create_time field.
	commentDescCreateTime := commentMixinFields0[0].Descriptor()
	// comment.DefaultCreateTime holds the default value on creation for the create_time field.
	comment.DefaultCreateTime = commentDescCreateTime.Default.(func() time.Time)
	// commentDescUpdateTime is the schema descriptor for update_time field.
	commentDescUpdateTime := commentMixinFields0[1].Descriptor()
	// comment.DefaultUpdateTime holds the default value on creation for the update_time field.
	comment.DefaultUpdateTime = commentDescUpdateTime.Default.(func() time.Time)
	// comment.UpdateDefaultUpdateTime holds the default value on update for the update_time field.
	comment.UpdateDefaultUpdateTime = commentDescUpdateTime.UpdateDefault.(func() time.Time)
	// commentDescUserID is the schema descriptor for user_id field.
	commentDescUserID := commentFields[0].Descriptor()
	// comment.DefaultUserID holds the default value on creation for the user_id field.
	comment.DefaultUserID = commentDescUserID.Default.(int)
	// commentDescPostID is the schema descriptor for post_id field.
	commentDescPostID := commentFields[1].Descriptor()
	// comment.DefaultPostID holds the default value on creation for the post_id field.
	comment.DefaultPostID = commentDescPostID.Default.(int)
	// commentDescContent is the schema descriptor for content field.
	commentDescContent := commentFields[2].Descriptor()
	// comment.DefaultContent holds the default value on creation for the content field.
	comment.DefaultContent = commentDescContent.Default.(string)
	postMixin := schema.Post{}.Mixin()
	postMixinFields0 := postMixin[0].Fields()
	_ = postMixinFields0
	postFields := schema.Post{}.Fields()
	_ = postFields
	// postDescCreateTime is the schema descriptor for create_time field.
	postDescCreateTime := postMixinFields0[0].Descriptor()
	// post.DefaultCreateTime holds the default value on creation for the create_time field.
	post.DefaultCreateTime = postDescCreateTime.Default.(func() time.Time)
	// postDescUpdateTime is the schema descriptor for update_time field.
	postDescUpdateTime := postMixinFields0[1].Descriptor()
	// post.DefaultUpdateTime holds the default value on creation for the update_time field.
	post.DefaultUpdateTime = postDescUpdateTime.Default.(func() time.Time)
	// post.UpdateDefaultUpdateTime holds the default value on update for the update_time field.
	post.UpdateDefaultUpdateTime = postDescUpdateTime.UpdateDefault.(func() time.Time)
	// postDescHashID is the schema descriptor for hash_id field.
	postDescHashID := postFields[0].Descriptor()
	// post.DefaultHashID holds the default value on creation for the hash_id field.
	post.DefaultHashID = postDescHashID.Default.(string)
	// postDescUserID is the schema descriptor for user_id field.
	postDescUserID := postFields[1].Descriptor()
	// post.DefaultUserID holds the default value on creation for the user_id field.
	post.DefaultUserID = postDescUserID.Default.(int)
	// postDescTitle is the schema descriptor for title field.
	postDescTitle := postFields[2].Descriptor()
	// post.DefaultTitle holds the default value on creation for the title field.
	post.DefaultTitle = postDescTitle.Default.(string)
	// postDescContent is the schema descriptor for content field.
	postDescContent := postFields[3].Descriptor()
	// post.DefaultContent holds the default value on creation for the content field.
	post.DefaultContent = postDescContent.Default.(string)
	// postDescTimesOfRead is the schema descriptor for times_of_read field.
	postDescTimesOfRead := postFields[4].Descriptor()
	// post.DefaultTimesOfRead holds the default value on creation for the times_of_read field.
	post.DefaultTimesOfRead = postDescTimesOfRead.Default.(int)
	userMixin := schema.User{}.Mixin()
	userMixinFields0 := userMixin[0].Fields()
	_ = userMixinFields0
	userFields := schema.User{}.Fields()
	_ = userFields
	// userDescCreateTime is the schema descriptor for create_time field.
	userDescCreateTime := userMixinFields0[0].Descriptor()
	// user.DefaultCreateTime holds the default value on creation for the create_time field.
	user.DefaultCreateTime = userDescCreateTime.Default.(func() time.Time)
	// userDescUpdateTime is the schema descriptor for update_time field.
	userDescUpdateTime := userMixinFields0[1].Descriptor()
	// user.DefaultUpdateTime holds the default value on creation for the update_time field.
	user.DefaultUpdateTime = userDescUpdateTime.Default.(func() time.Time)
	// user.UpdateDefaultUpdateTime holds the default value on update for the update_time field.
	user.UpdateDefaultUpdateTime = userDescUpdateTime.UpdateDefault.(func() time.Time)
	// userDescHashID is the schema descriptor for hash_id field.
	userDescHashID := userFields[0].Descriptor()
	// user.DefaultHashID holds the default value on creation for the hash_id field.
	user.DefaultHashID = userDescHashID.Default.(string)
	// userDescMobile is the schema descriptor for mobile field.
	userDescMobile := userFields[1].Descriptor()
	// user.DefaultMobile holds the default value on creation for the mobile field.
	user.DefaultMobile = userDescMobile.Default.(string)
	// userDescPassword is the schema descriptor for password field.
	userDescPassword := userFields[2].Descriptor()
	// user.DefaultPassword holds the default value on creation for the password field.
	user.DefaultPassword = userDescPassword.Default.(string)
	// userDescLevel is the schema descriptor for level field.
	userDescLevel := userFields[3].Descriptor()
	// user.DefaultLevel holds the default value on creation for the level field.
	user.DefaultLevel = userDescLevel.Default.(int)
	// userDescNickname is the schema descriptor for nickname field.
	userDescNickname := userFields[4].Descriptor()
	// user.DefaultNickname holds the default value on creation for the nickname field.
	user.DefaultNickname = userDescNickname.Default.(string)
	// userDescAvatar is the schema descriptor for avatar field.
	userDescAvatar := userFields[5].Descriptor()
	// user.DefaultAvatar holds the default value on creation for the avatar field.
	user.DefaultAvatar = userDescAvatar.Default.(string)
	// userDescBio is the schema descriptor for bio field.
	userDescBio := userFields[6].Descriptor()
	// user.DefaultBio holds the default value on creation for the bio field.
	user.DefaultBio = userDescBio.Default.(string)
}
