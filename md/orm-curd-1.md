### 控制反转思想来提高CRUD效率
<img src="../images/crud.jpeg" width="100%">

开发中小伙伴都会遇到针对同一张表查询有N种的查询条件，例如users用户表经常会有以下的查询方法：
```go
// id查询
func (repo *User) FetchByID(ctx context.Context, id int) *ent.User {
	return repo.db.User.Query().Where(user.ID(id)).FirstX(ctx)
}
// 手机号查询
func (repo *User) FetchByMobile(ctx context.Context, mobile string) *ent.User {
	return repo.db.User.Query().Where(user.Mobile(mobile)).FirstX(ctx)
}
// 手机号和密码查询
func (repo *User) FetchByMobileAndPassword(ctx context.Context, mobile string, password string) *ent.User {
	return repo.db.User.Query().Where(user.Mobile(mobile), user.Password(password)).FirstX(ctx)
}
// 昵称查询
func (repo *User) FetchByNickname(ctx context.Context, nickname string) ent.Users {
	return repo.db.User.Query().Where(user.Nickname(nickname)).AllX(ctx)
}
```
上面的代码是不是很眼熟，相信在你也会经常写过或遇到这样的代码，这又是程序开发中没有技术含量的搬砖活，如果项目存在接口层，还需要在接口层定义接口函数，更加费时费力，有什么方式可以优化呢？

对上面代码的观察可以发现这几个函数都具有结构相似性，不同处在于查询参数和Where语句的不同，有没有一种方式可以在外部控制查询参数和Where语句，达到灵活使用的目的呢？

### InversionOfControl 控制反转
控制反转一种设计思想，有反就有正，先看一下正控制是怎么控制的。

上文中Fetch开头的函数都属于正控制，controller.One函数传递参数id调用FetchByID，One函数控制不了FetchByID函数内部怎么使用id参数，等于id、不等于id、大于、小于id都得看FetchByID内部如何实现，这种方式就是FetchByID控制One，属于正控制，符合常规逻辑。

反过来想，controller.One可以控制FetchByID内部对id参数的使用，那么就是控制反转的逻辑，是控制反转代码是怎么实现的呢？
```go
// controller层
func (ctr *User) One(c *gin.Context) {
    id := 1
    //entUser := ctr.repo.FetchByID(c.Request.Context(), id)

    entUser := ctr.repo.FetchOne(c.Request.Context(), func (opt *ent.UserQuery) {
        opt.Where(user.ID(id))
    })
    c.JSON(http.StatusOK, entUser)
}

// repository层
func (repo *User) FetchOne(ctx context.Context, optionFunc func(opt *ent.UserQuery)) *ent.User {
    builder := repo.db.User.Query()
    optionFunc(builder)
    return builder.FirstX(ctx)
}
```
通过匿名函数传参的方式我们实现了控制反转的逻辑，这样在controller层就可以灵活组合不同的查询条件实现多变的查询需求，同理增删改也可以通过这种思想来优化。
```go
// 增
func (repo *User) Create(ctx context.Context, optionFunc func(opt *ent.UserCreate)) *ent.User {
	builder := repo.db.User.Create()
	optionFunc(builder)
	return builder.SaveX(ctx)
}

// 删
func (repo *User) Delete(ctx context.Context, optionFunc func(opt *ent.UserDelete)) int {
    builder := repo.db.User.Delete()
    optionFunc(builder)
    return builder.ExecX(ctx)
}

// 改
func (repo *User) Update(ctx context.Context, optionFunc func(opt *ent.UserUpdate)) int {
    builder := repo.db.User.Update()
    optionFunc(builder)
    return builder.SaveX(ctx)
}
```

前文中的帖子需求优化如下：
```go
func (ctr *Post) Many(c *gin.Context) {
	entPosts := ctr.repo.FetchMany(c.Request.Context(), func(opt *ent.PostQuery) {
		opt.WithUser()
	})
	c.JSON(http.StatusOK, entPosts)
}

func (ctr *Post) One(c *gin.Context) {
	id := 1
	posts := ctr.repo.FetchOne(c.Request.Context(), func(opt *ent.PostQuery) {
            opt.WithUser().WithComments(func(o *ent.CommentQuery) {
                o.WithUser()
            }).Where(post.ID(id))
	})
	c.JSON(http.StatusOK, posts)
}

func (ctr *Post) Comments(c *gin.Context) {
	comments := ctr.repoComment.FetchMany(c.Request.Context(), func(opt *ent.CommentQuery) {
            opt.WithUser().WithPost().Order(ent.Desc(comment.FieldCreateTime))
	})
	c.JSON(http.StatusOK, comments)
}
```
[源码链接](https://github.com/wizardshan/elegantGo/tree/main/app/chapter-orm-crud-1)

众所周知增删改查在日常开发中占了很大的工作量，搬砖活不可避免，通过控制反转的思想提高增删改查的开发效率，我们可以做一个高效的CRUD程序员。

### entGo框架事务
```go
func (repo *User) Register(ctx context.Context) (*ent.User, error) {
	mobile := "13000000003"
	password := "a906449d5769fa7361d7ecc6aa3f6d28"
	level := 30
	nickname := "昵称3"
	avatar := "头像3.png"
	bio := "个人介绍3"

	entUser := repo.Create(ctx, func(opt *ent.UserCreate) {
		opt.SetMobile(mobile).SetPassword(password).SetLevel(level).SetNickname(nickname).SetAvatar(avatar).SetBio(bio)
	})

	hashID, err := hashid.EncodeUserID(entUser.ID)
	if err != nil {
		return nil, err
	}
	repo.Update(ctx, func(opt *ent.UserUpdate) {
		opt.SetHashID(hashID).Where(user.ID(entUser.ID))
	})

	return entUser, nil
}
```
这是一段用户注册的代码，第一步用户信息保存在数据库中，获取到自增id，然后对自增id进行加密操作，然后再把id加密结果保存在用户表里，我们使用了刚刚定义的CRUD函数，但是又需要使用事务来保持数据一致性，这又如何实现呢，请看下文分解。

