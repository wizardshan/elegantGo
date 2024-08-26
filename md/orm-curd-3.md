### entGo框架CRUD通用代码自动化生成
对数据表的常规操作如下：
```go
Fetch      // 通过自增ID获取单个记录
FetchOne   // 通过条件过滤获取单个记录
FetchMany  // 通过条件过滤获取多个记录
Count      // 通过条件过滤判断记录条数
Exist      // 通过条件过滤判断记录是否存在
Create     // 增加记录
Update     // 更新记录
Delete     // 删除记录
```
通过模板文件自动生成表对应的操作对象，速度又快又不容易出错，具体生成代码参照template目录的gen_repo.go。

执行生成命令：<br>
```go
//go:generate go run gen_repo.go -entName User

User为模型参数名
```
命令执行成功后会在repository目录中生成user_crud.go文件。

```go
// 模板repo.tmpl文件部分代码

type {{.entName}} struct {
	repo
}

func New{{.entName}}(db *ent.Client) *{{.entName}} {
	repo := new({{.entName}})
	repo.db = db
	return repo
}

func (repo *{{.entName}}) Fetch(ctx context.Context, id int) *ent.{{.entName}} {
	return repo.fetchOne(ctx, repo.db, func(opt *ent.{{.entName}}Query) {
		opt.Where({{.entName | ToLower}}.ID(id))
	})
}

func (repo *{{.entName}}) FetchOne(ctx context.Context, optionFunc func(opt *ent.{{.entName}}Query)) *ent.{{.entName}} {
	return repo.fetchOne(ctx, repo.db, optionFunc)
}

func (repo *{{.entName}}) fetchOne(ctx context.Context, db *ent.Client, optionFunc func(opt *ent.{{.entName}}Query)) *ent.{{.entName}} {
	builder := db.{{.entName}}.Query()
	optionFunc(builder)
	return builder.FirstX(ctx)
}
```
[源码链接](../chapter-orm-crud-3)

具体生成哪些通用的操作方法，可以参照Java的MyBatis-Plus框架，写一个大而全的repo模板。
```go
saveBatch        // 插入（批量）
saveOrUpdate     // 插入或批量
removeById       // 根据 ID 删除
removeByIds      // 删除（根据ID 批量删除）
updateById       // 根据 ID 选择修改
updateBatchById  // 根据ID 批量更新
等等...
```
[MyBatis-Plus官网](https://baomidou.com/guides/data-interface/)

**Go Generation**<br>
Go Generation是解决结构重复性代码的利器，我们使用代码生成工具来生成一些重复的、机械的代码，例如序列化和反序列化代码、API定义代码、数据库访问代码等。这些代码通常具有固定的模板和格式，并且不需要手动编写。使用代码生成工具可以自动创建这些代码，同时也可以使代码更加一致和易于维护，从而减少编写代码的时间和出错的机会。

后文中会多次使用代码自动化生成的模式，请大家仔细体会。









