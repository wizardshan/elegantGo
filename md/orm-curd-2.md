### engGo框架事务
```go
func (repo *repoDB) withTx(ctx context.Context, client *ent.Client, fn func(tx *ent.Tx) error) error {
	tx, err := client.Tx(ctx)
	if err != nil {
		return err
	}
	defer func() {
		if v := recover(); v != nil {
			tx.Rollback()
			panic(v)
		}
	}()
	if err := fn(tx); err != nil {
		if errRollback := tx.Rollback(); errRollback != nil {
			err = errors.Wrapf(err, "rolling back transaction: %v", errRollback)
		}
		return err
	}
	if err := tx.Commit(); err != nil {
		return errors.Wrapf(err, "committing transaction: %v", err)
	}
	return nil
}
```
上述代码是entGo框架官方推荐的事务用法，我们看一下上文中的用户注册功能事务如何实现，代码如下：
```go
func (repo *User) Register(ctx context.Context) (*ent.User, error) {

	mobile := "13000000003"
	password := "a906449d5769fa7361d7ecc6aa3f6d28"
	level := 30
	nickname := "昵称3"
	avatar := "头像3.png"
	bio := "个人介绍3"

	var entUser *ent.User
	err := repo.withTx(ctx, repo.db, func(tx *ent.Tx) error {
		db := tx.Client()
		entUser = repo.create(ctx, db, func(builder *ent.UserCreate) {
			builder.SetMobile(mobile).SetPassword(password).SetLevel(level).SetNickname(nickname).SetAvatar(avatar).SetBio(bio)
		})
		hashID, err := hashid.EncodePostID(entUser.ID)
		if err != nil {
			return err
		}

		repo.update(ctx, db, func(builder *ent.UserUpdate) {
			builder.SetHashID(hashID).Where(user.ID(entUser.ID))
		})

		return nil
	})

	return entUser, err

}
```
增删改查函数为了适应事务的使用做了修改，我们用FetchOne举例，定义了fetchOne函数，通过传入事务db和非事务db，供事务和非事务操作使用，其他增删改查同理，这样就可以共用，减少重复代码。
```go
func (repo *User) FetchOne(ctx context.Context, optionFunc func(builder *ent.UserQuery)) *ent.User {
	return repo.fetchOne(ctx, repo.db, optionFunc)
}

func (repo *User) fetchOne(ctx context.Context, db *ent.Client, optionFunc func(builder *ent.UserQuery)) *ent.User {
	builder := db.User.Query()
	optionFunc(builder)
	return builder.FirstX(ctx)
}
```


### engGo框架使用自定义SQL查询
少数场景下需要我们手写SQL查询，请参照controller层User.One和Many方法。

### engGo框架的一些用法
**1、InsertOrUpdate**<br>
在generate.go文件中打开sql/upsert开关：<br> 
go:generate go run -mod=mod entgo.io/ent/cmd/ent generate --feature sql/upsert ./schema<br>
代码参照controller层User.Upsert方法。

**2、ORDER BY rand()**<br>
代码参照controller层User.Rand方法。

[源码链接](https://github.com/wizardshan/elegantGo/tree/main/app/chapter-orm-crud-2)

我们通过定义CRUD通用方法大大提高了增删改查开发效率，现在又出了新的问题，针对每张表定义CURD通用方法又成了效率低下的重复工作，我们略读repository层，user、post、comment的通用方法结构非常类似，这属于结构重复性代码，那我们又该怎么优化呢，请看下文分解。


