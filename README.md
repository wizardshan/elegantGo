# elegantGo
Go编程：优雅永不过时
- 1、[被误解的ORM，只是单纯的不手写SQL？](md/orm.md)
- 2、entGo框架高效实现ORM
- 3、控制反转思想来提高CRUD开发效率
- 4、entGo框架事务优化 的特殊用法 事务 insertupdate rand



- 1、参数校验三种方式的最优解
- 2、参数相关的三种安全漏洞之一：SQL注入
- 3、参数相关的三种安全漏洞之二：XSS （去空格 支付宝账号只能绑定一个账号）
- 4、项目分层和模型转换那点事儿
- 5、使用pool池优化模型转换（实验性）
- 6、request和response分层模型灵活应用的实际案例 fastadmin 业务模型与数据库模型的区别 活动 是否开始 进行中 是否结束 倒计时



- 11、项目层级的缓存策略
- 12、缓存雪崩、击穿、穿透、缓存污染
- 13、项目结构分层的自动化生成
- 14、项目结构实践中的完善，user_id,重复提交**

面向对象的难点：对象族

最后一章 应付差事的CodeReview

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