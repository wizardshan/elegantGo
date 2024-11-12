# elegantGo
Go编程：优雅永不过时
- 1、[参数校验：代码由湿到干的过程](md/param-validate.md)
- 2、[参数校验：分离关注点-代码复杂度降维利器](md/param-validate-soc.md)
- 3、[参数校验：抽象到具体的思维方式](md/param-complex-validate.md)



- 1、[被误解的ORM，只是单纯的不手写SQL？](md/orm.md)
- 2、[entGo框架高效实现ORM](md/orm-entgo.md)
- 3、[控制反转提高CRUD开发效率](md/orm-curd-1.md)
- 4、[entGo框架CRUD事务优化](md/orm-curd-2.md)
- 5、[entGo框架CRUD通用代码自动化生成](md/orm-curd-3.md)

- 9、参数相关的安全漏洞：SQL注入
- 10、参数相关的安全漏洞：XSS 




- （去空格 支付宝账号只能绑定一个账号）
- 4、项目分层和模型转换那点事儿
- 5、使用pool池优化模型转换（实验性）
- 6、request和response分层模型灵活应用的实际案例 fastadmin 业务模型与数据库模型的区别 活动 是否开始 进行中 是否结束 倒计时



- 11、项目层级的缓存策略
- 12、缓存雪崩、击穿、穿透、缓存污染
- 13、项目结构分层的自动化生成
- 14、项目结构实践中的完善，user_id,重复提交**

the right status code is 4xx ("you messed up") or 5xx ("I messed up").

面向对象的难点：对象族

最后一章 应付差事的CodeReview

事务开始结束日志

MySQL事务隔离级别？ 扣库存的合适级别 为什么MySQL 默认隔离级别是RR，又被阿里设置为RC    RC级别需要程序员乐观锁解决
SELECT @@global.transaction_isolation


console.log(new Date("2024-07-03T08:00:10.325Z").toLocaleString())
console.log(new Date().toISOString())


1、理解ORM，减少多端沟通成本，避免效率低下手写SQL
2、控制反转优化增删改查
3、共用参数校验模型，减少重复代码 
4、自动生成模型转换函数，提高代码开发效率和执行效率
5、模型池化，减少GC次数，降低服务器负担
6、项目分层代码自动化
7、无缝感知的cache


https://help.aliyun.com/document_detail/472990.html

**整体方面：**<br>
这是日常开发中很常见的数据流水线，流程如下：<br>
1、获取原始数据 <br>
2、原始数据有效性校验<br>
3、原始数据转换成业务数据<br>
4、业务数据的应用<br>

>思考几分钟，1234步骤是否覆盖了日常开发中90%的业务流程，例如从数据库获取数据、从第三方接口获取数据。

UnmarshalJSON方法揉杂了数据的格式化、有效性校验、数据转化、业务逻辑判断等多种意图的代码，当业务有所变动时，定位需要改动的代码第一步就是通读整个方法，同时当出现问题定位bug时，第一步也是通读整个方法，前者就是代码的可读性差，后者就是代码的可维护行差。

对于复杂问题的解决办法


<img src="../images/postman-complex-parameters.jpg">

代码片段重复
方法内重复
1、封装成方法内全局变量或常量
2、封装成独立方法
结构体内跨方法重复
1、封装成独立方法
2、封装成包全局变量或常量
3、封装成独立结构体
跨结构体重复
1、封装成包全局变量或常量
2、封装成独立结构体


TZ="Asia/Shanghai" go run main.go

https://wangbjun.site/2021/coding/golang/event-bus.html


实际上，对于应用系统而言，只有三种类型的异常：
BizException：业务异常，有明确的业务语义，不需要记录Error日志，不需要Retry
SysException：已知的系统异常，需要记录Error日志，可以Retry
Exception：未知的其它异常，需要完整的Error Stack日志，可以Retry


clear架构demo：
https://github.com/manuelkiessling/go-cleanarchitecture
https://github.com/bxcodec/go-clean-arch

https://github.com/xxjwxc/uber_go_guide_cn

https://darjun.github.io/2020/04/01/godailylib/govaluate/

工具函数库
https://github.com/gookit/goutil
https://github.com/duke-git/lancet

https://github.com/hashicorp/golang-lru


Sort:CreateTime
Order:ASC
Offset:20
Limit:10
Amount:{"Start":10, "End":100}
Level:10
Level:20
Level:30
Status:normal
Status:cancel
Status:invalid
CreateTime:{"Start":"2024-01-01T00:00:00Z", "End":"2024-10-01T00:00:00Z"}
Filter:{"ID":1,"Nickname":"昵称","Amount":{"Start":10, "End":100},"Level":[10,20,30],"Status":["normal","cancel","invalid"],"CreateTime":{"Start":"2024-01-01T00:00:00Z", "End":"2024-10-01T00:00:00Z"}}

