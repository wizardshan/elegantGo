### engGo框架CRUD代码自动生成
对数据库的常规操作如下：
```go
Fetch      // 通过自增id获取单个记录
FetchOne   // 通过条件过滤获取单个记录
FetchMany  // 通过条件过滤获取多个记录
Count      // 通过条件过滤判断记录条数
Exist      // 通过条件过滤判断记录是否存在
Create     // 增加记录
Update     // 更新记录
Delete     // 删除记录
```
https://baomidou.com/guides/data-interface/

[源码链接](https://github.com/wizardshan/elegantGo/tree/main/app/chapter-orm-crud-3)
