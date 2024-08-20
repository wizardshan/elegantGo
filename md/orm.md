### 什么是ORM?
对象关系映射（Object Relational Mapping，简称ORM），是一种程序技术，用于实现面向对象编程里不同类型系统的数据之间的转换。

**ORM是一种程序技术用于实现面向对象编程里在不同类型系统的数据转换时保持对象之间的映射关系。**

很多同学认为ORM就是使用ORM框架实现数据的增删改查操作，但是根据ORM的定义，ORM是强调在不同类型的系统数据转换时要保持正确的对象映射关系，不同类型的系统并不仅仅是数据库系统的数据和应用程序的数据相互转换，也包括应用程序中的数据和客户端的数据相互转换；

**什么是对象关系？** <br>
微博用户发帖举例：<br>
**一对多关系**：用户可以拥有自己发布的多个帖子，用户对帖子是一对多关系；<br>
**一对一关系**：帖子只能拥有一个发表它的用户，帖子对用户是一对一关系；<br>
**一对多关系**：帖子可以拥有多个评论，帖子对评论是一对多关系；<br>
**一对一关系**：评论只能属于一个帖子，评论对帖子是一对一关系；<br>
**一对多关系**：用户可以针对一个或多个帖子发布评论，用户对评论是一对多关系；<br>

另外还有多对多关系，读者自行了解。

如何在Go语言中来体现这种对象关系呢？
```go
type User struct {
    Posts []*Post         // 用户可以拥有多个帖子
    Comments []*Comment   // 用户可以拥有多个评论
}

type Post struct {
    User *User            // 帖子只能拥有一个发表它的用户
    Comments []*Comment   // 帖子可以拥有多个针对它的评论
}

type Comment struct {
    User *User            // 评论只能拥有一个发表它的用户
    Post *Post            // 评论只能属于一个帖子
}
```

**思考：为什么要保持对象之间的映射关系？**

我们用代码来探索这个问题
```
帖子表posts：
id(自增) hash_id(加密id) user_id(用户id) title(标题) content(内容) views(浏览量) create_time(创建时间) update_time(更新时间)

用户表users：
id(自增)  mobile(手机号) password(密码) nickname(昵称) avatar(头像) bio(个人简介) create_time(创建时间) update_time(更新时间)

评论表comments：
id(自增)  user_id(用户id) post_id(帖子id) content(内容) create_time(创建时间) update_time(更新时间)
```

需求1：帖子列表，需要用户头像、用户昵称数据 <br/>
需求2：帖子详情，需要用户头像、用户昵称数据；需要评论列表数据，评论需要用户头像、用户昵称数据 <br/>

先用手写SQL的方式来实现：
```json
帖子列表:
posts表联users表查询：SELECT posts.`id`, posts.`hash_id`, posts.`user_id`, posts.`title`, posts.`content`, posts.`views`, posts.`create_time`, posts.`update_time`, users.`nickname`, users.`avatar` FROM posts, users WHERE posts.user_id=users.id ORDER BY posts.create_time DESC
[
    {
        "ID": 1,
        "HashID": "oKqk6tMl7z",
        "UserID": 1,
        "Title": "标题1",
        "Content": "内容1",
        "Views": 100,
        "CreateTime": "2024-08-01T00:00:00Z",
        "UpdateTime": "2024-08-02T00:00:00Z",
        "Nickname": "昵称1",
        "Avatar": "头像1.png",
        "Comments": null
    },
    {
        "ID": 2,
        "HashID": "02qN7SQyOb",
        "UserID": 2,
        "Title": "标题2",
        "Content": "内容2",
        "Views": 200,
        "CreateTime": "2024-08-01T00:00:00Z",
        "UpdateTime": "2024-08-02T00:00:00Z",
        "Nickname": "昵称2",
        "Avatar": "头像2.png",
        "Comments": null
    }
]

帖子详情：
posts表联users表查询：SELECT posts.`id`, posts.`hash_id`, posts.`user_id`, posts.`title`, posts.`content`, posts.`views`, posts.`create_time`, posts.`update_time`, users.`nickname`, users.`avatar` FROM posts, users WHERE posts.user_id=users.id AND posts.`id`= %d
comments表联users表查询：SELECT comments.`id`, comments.`user_id`, comments.`post_id`, comments.`content`, comments.`create_time`, comments.`update_time`, users.`nickname`, users.`avatar` FROM comments, users WHERE comments.user_id=users.id AND comments.`post_id`= %d ORDER BY comments.create_time DESC
{
    "ID": 1,
    "HashID": "oKqk6tMl7z",
    "UserID": 1,
    "Title": "标题1",
    "Content": "内容1",
    "Views": 100,
    "CreateTime": "2024-08-01T00:00:00Z",
    "UpdateTime": "2024-08-02T00:00:00Z",
    "Nickname": "昵称1",
    "Avatar": "头像1.png",
    "Comments": [
        {
        "ID": 1,
        "UserID": 1,
        "PostID": 1,
        "Content": "评论1",
        "CreateTime": "2024-08-01T00:00:00Z",
        "UpdateTime": "2024-08-02T00:00:00Z",
        "Nickname": "昵称1",
        "Avatar": "头像1.png"
        }
    ]
}
```
[源码链接](https://github.com/wizardshan/elegantGo/tree/main/app/chapter-orm-1)

当对接接口的程序员看到这两个JSON结构，会有两个疑问：<br/>
1、post数据里包含Nickname、Avatar属性，这两个属性是post自带属性吗？<br/>
2、同上，comment数据里也包含Nickname、Avatar属性，这两个属性是comment自带属性吗？<br/>

通过沟通或者看字段注释才能确定这两个属性是user的属性，这样无形中就降低了开发效率。

我们可以通过增加属性前缀解决属性归属的不确定性，例如帖子详情JSON数据格式改成如下：
```json
{
    "ID": 1,
    "HashID": "oKqk6tMl7z",
    "UserID": 1,
    "Title": "标题1",
    "Content": "内容1",
    "TimesOfRead": 100,
    "CreateTime": "2024-08-01T00:00:00Z",
    "UpdateTime": "2024-08-02T00:00:00Z",
    "UserNickname": "昵称1",              // Nickname => UserNickname
    "UserAvatar": "头像1.png",            // Avatar   => UserAvatar
    "Comments": [
        {
            "ID": 1,
            "UserID": 1,
            "PostID": 1,
            "Content": "评论1",
            "CreateTime": "2024-08-01T00:00:00Z",
            "UpdateTime": "2024-08-02T00:00:00Z",
            "UserNickname": "昵称1",
            "UserAvatar": "头像1.png"
        }
    ]
}
```
[源码链接](https://github.com/wizardshan/elegantGo/tree/main/app/chapter-orm-2)

Nickname、Avatar属性名称加上模型名称前缀User，重新命名为UserNickname、UserAvatar，从而解决了属性归属的不确定性。

这样还会有什么问题呢？我们先看一下客户端如何对接帖子详情接口：
```java
//客户端对接帖子详情接口过程（Android举例）

// Comment模型
public class Comment {
    public Integer ID;   
    public Integer UserID;
    public Integer PostID;
    public String Content;   
    public String CreateTime;
    public String UpdateTime;
    
    public String UserNickname;
    public String UserAvatar;
}

// Post模型
public class Post {
    public Integer ID;   
    public String HashID;   
    public Integer UserID;   
    public String Title;    
    public String Content;    
    public Integer TimesOfRead;
    public String CreateTime;
    public String UpdateTime;
    
    public String UserNickname;
    public String UserAvatar;
    
    public ArrayList<Comment> Comments; 
}
```
对接完成之后，随着业务的变动，产品经理提出新需求：
> 增加用户等级功能，帖子和评论增加用户等级字段

服务端程序开发过程：<br/>
步骤一：用户表增加level字段<br/>
步骤二：查询users表SQL增加level<br/>
步骤三：Post、Comment模型增加UserLevel属性<br/>

客户端开发过程：<br/>
Post、Comment模型，增加UserLevel属性<br/>
[源码链接](https://github.com/wizardshan/elegantGo/tree/main/app/chapter-orm-3)

>思考：为什么users表增加1个level字段导致服务端要修改3条SQL、2个模型，并且连带客户端也要修改2个模型？

本质上是代码实现的接口数据格式没有遵循ORM标准，破坏了对象关系映射，原本实现业务需求底层数据用到了三张表(posts、comments、users)，到了代码层只有Post和Comment两个模型，并且这两个模型错误的拥有了用户的部分属性。

我们看一下遵循ORM规范的接口数据结构是什么样子的，以及如何应对这样的需求变化

```json
帖子列表：
[
    {
        "ID": 1,
        "HashID": "oKqk6tMl7z",
        "UserID": 1,
        "Title": "标题1",
        "Content": "内容1",
        "TimesOfRead": 100,
        "CreateTime": "2024-08-01T00:00:00Z",
        "UpdateTime": "2024-08-02T00:00:00Z",
        "User": {
            "ID": 1,
            "Nickname": "昵称1",
            "Avatar": "头像1.png"
        },
        "Comments": null
    },
    {
        "ID": 2,
        "HashID": "02qN7SQyOb",
        "UserID": 2,
        "Title": "标题2",
        "Content": "内容2",
        "TimesOfRead": 200,
        "CreateTime": "2024-08-01T00:00:00Z",
        "UpdateTime": "2024-08-02T00:00:00Z",
        "User": {
            "ID": 2,
            "Nickname": "昵称2",
            "Avatar": "头像2.png"
        },
        "Comments": null
    }
]

帖子详情：
{
    "ID": 1,
    "HashID": "oKqk6tMl7z",
    "UserID": 1,
    "Title": "标题1",
    "Content": "内容1",
    "TimesOfRead": 100,
    "CreateTime": "2024-08-01T00:00:00Z",
    "UpdateTime": "2024-08-02T00:00:00Z",
    "User": {
        "ID": 1,
        "Nickname": "昵称1",
        "Avatar": "头像1.png"
    },
    "Comments": [
        {
            "ID": 1,
            "UserID": 1,
            "PostID": 1,
            "Content": "评论1",
            "CreateTime": "2024-08-01T00:00:00Z",
            "UpdateTime": "2024-08-02T00:00:00Z",
            "User": {
                "ID": 1,
                "Nickname": "昵称1",
                "Avatar": "头像1.png"
            },
            "Post": null
        }
    ]
}
```
[源码链接](https://github.com/wizardshan/elegantGo/tree/main/app/chapter-orm-4)

```java
//客户端对接帖子详情接口过程（Android举例）

// User模型
public class User {
    public Integer ID;   
    public String Nickname;    
    public String Avatar;
}

// Comment模型
public class Comment {
    public Integer ID;   
    public Integer UserID;
    public Integer PostID;
    public String Content;   
    public String CreateTime;
    public String UpdateTime;
    
    public User User; 
    public User Post; 
}

// Post模型
public class Post {
    public Integer ID;   
    public String HashID;   
    public Integer UserID;   
    public String Title;    
    public String Content;    
    public Integer TimesOfRead;
    public String CreateTime;
    public String UpdateTime;
    
    public User User; 
    public ArrayList<Comment> Comments; 
}
```
实现增加用户等级功能<br/>
服务端程序开发过程：<br/>
步骤一：用户表增加level字段<br/>
步骤二：查询users表SQL增加level<br/>
步骤三：User模型增加Level属性<br/>

客户端开发过程：<br/>
User模型增加Level属性<br/>

相对于联表实现方式，遵循ORM规范的JSON数据格式减少了服务器端和客户端代码改动范围，同时在服务器端和客户端保持了对象关系映射，对接接口的程序员可以很容易理解数据之间的层级结构和属性范围，减少了沟通成本。

再看一个稍微复杂的例子加深理解<br/>
需求：最新评论接口，需要评论内容、帖子主题、用户昵称数据
```json
[
  {
    "ID": 1,
    "UserID": 1,
    "PostID": 1,
    "Content": "评论1",
    "CreateTime": "2024-08-01T00:00:00Z",
    "UpdateTime": "2024-08-02T00:00:00Z",
    "User": {
      "ID": 1,
      "Nickname": "昵称1",
      "Avatar": "头像1.png"
    },
    "Post": {
      "ID": 1,
      "HashID": "oKqk6tMl7z",
      "UserID": 1,
      "Title": "标题1",
      "Content": "内容1",
      "Views": 100,
      "CreateTime": "2024-08-01T00:00:00Z",
      "UpdateTime": "2024-08-02T00:00:00Z",
      "User": null,
      "Comments": null
    }
  },
  {
    "ID": 2,
    "UserID": 2,
    "PostID": 2,
    "Content": "评论2",
    "CreateTime": "2024-08-01T00:00:00Z",
    "UpdateTime": "2024-08-02T00:00:00Z",
    "User": {
      "ID": 2,
      "Nickname": "昵称2",
      "Avatar": "头像2.png"
    },
    "Post": {
      "ID": 2,
      "HashID": "02qN7SQyOb",
      "UserID": 2,
      "Title": "标题2",
      "Content": "内容2",
      "Views": 200,
      "CreateTime": "2024-08-01T00:00:00Z",
      "UpdateTime": "2024-08-02T00:00:00Z",
      "User": null,
      "Comments": null
    }
  }
]
```

大家可能会有疑问，这样的模型嵌套模型的json数据格式，在客户端、服务器是否容易解析，会增加对接的工作量吗？<br/>

不同语言、不同平台对于模型嵌套数据格式都有很成熟的解析库：<br/>
python：Dataclasses JSON：[链接](https://github.com/lidatong/dataclasses-json)  <br/>
java：Gson FastJson <br/>
golang：resty [链接](https://github.com/go-resty/resty)  <br/>
typescript：class-transformer [链接](https://github.com/typestack/class-transformer)  <br/>

ios swift：MJExtension： [链接](https://github.com/CoderMJLee/MJExtension)  <br/>
android kotlin java.：Moshi： [链接](https://github.com/square/moshi)  <br/>


回到正文，凡事有利有弊，遵循ORM规范的代码存在什么问题呢？

通过帖子列表代码实现过程可以发现，遵循ORM规范的代码量差不多联表实现代码量的两倍：<br/>
步骤一：获取帖子列表<br/>
步骤二：通过帖子列表得到用户ID数组<br/>
步骤三：通过用户ID数组获取用户列表<br/>
步骤四：用户列表mapping格式化，得到用户ID和用户对象的映射map，目的是避免双循环，提高效率<br/>
步骤五：循环帖子列表，通过映射map获取到用户对象，赋值给帖子对象中的用户属性<br/>

这还只两张表实现ORM的代码量，如果三张表或者更多表参与，代码量更多，参照post/comments接口。

那有什么好方法来解决这个问题呢？请看下文分解。



