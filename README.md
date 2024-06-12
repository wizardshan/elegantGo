# elegantGo
Go编程：优雅永不过时
- 1、被误解的ORM，只是单纯的不手写SQL？
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


console.log(new Date("2024-05-14T17:23:52.418905+08:00").toLocaleString())

1、理解ORM，减少多端沟通成本，避免效率低下手写SQL
2、控制反转优化增删改查
3、共用参数校验模型，减少重复代码 
4、自动生成模型转换函数，提高代码开发效率和执行效率
5、模型池化，减少GC次数，降低服务器负担
6、项目分层代码自动化
7、无缝感知的cache


https://help.aliyun.com/document_detail/472990.html