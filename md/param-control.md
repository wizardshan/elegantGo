上文提到的参数"IDs=1,2,3,4,5"，需要把这个被逗号分割的字符串转换为int切片[]int{1,2,3,4,5}，用于后续代码的使用。<br>
解决思路：
```go
type UserMany struct {
    IDs string
}

func (req *UserMany) IDsValues() (values []int) {
    nums := strings.Split(req.IDs, ",")
    for _, num := range nums {
        value, _ := strconv.Atoi(num)
        values = append(values, value)
    }
    return
}
```
代码分析：<br>
通过request.UserMany模型解析IDs参数，问题显而易见，当另外一个接口也需要IDs参数时，想对应的request模型同样需要实现IDsValues函数，这样就导致的重复代码。

解决思路：
```go
type UserMany struct {
    IDs IDsField
}

type IDsField string

func (req *IDsField) Values() (values []int) {
    nums := strings.Split(string(*req), ",")
    for _, num := range nums {
        value, _ := strconv.Atoi(num)
        values = append(values, value) 
    }
    return
}
```
代码分析：<br>
通过独立出来的IDsField，把解析过程封装其内部，当另外一个接口需要IDs参数时，重用即可；<br>
这种方式又带来另外一个问题，例如用户等级Level，当需要查询多个用户等级，就需要在新建一个LevelsField结构体，如下：
```go
type LevelsField string

func (req *LevelsField) Values() (values []int) {
    nums := strings.Split(string(*req), ",")
    for _, num := range nums {
        value, _ := strconv.Atoi(num)
        values = append(values, value)
    }
    return
}
```
我们把思维再抽象一点，根据字符串"1,2,3,4,5"的特征新建一个通用的NumbersFieldV1结构体，如下：
```go
type NumbersFieldV1 string

func (req *NumbersFieldV1) Values() (values []int) {
    nums := strings.Split(string(*req), ",")
    for _, num := range nums {
        value, _ := strconv.Atoi(num)
        values = append(values, value)
    }
    return
}

type UserMany struct {
    IDs    NumbersFieldV1
    Levels NumbersFieldV1
}
```
这样，用户ID和用户等级等多数字参数就可以复用NumbersField。

下面进入本文的重点，我们先审视一下NumbersField.Values方法，Values方法方法的目的是把被逗号分割的string，转换为[]int，分为三步骤：<br>
步骤1：字符串Split为[]string<br>
步骤2：for循环[]string获取切片里的字符串<br>
步骤3：字符串转换为int，转换规则比较宽松，失败忽略，只收集转换成功的数据<br>

目前为止，看起来没什么问题，当需要增加严格参数校验规则，转化失败就报错，代码如下：
```go
func (req *NumbersField) MustValues() (values []int, err error) {
    nums := strings.Split(string(*req), ",")
    for _, num := range nums {
        var value int
        if value, err = strconv.Atoi(num); err != nil {
            return
        }
        values = append(values, value)
    }
    return
}
```






https://blog.csdn.net/weixin_44358480/article/details/139697331

https://www.jianshu.com/p/8b917163b557

https://blog.jaronnie.com/%e7%bc%96%e7%a8%8b%e7%9a%84%e6%9c%ac%e8%b4%a8%ef%bc%9alogic-%e4%b8%8e-control-%e5%88%86%e7%a6%bb%ef%bc%81/

https://time.geekbang.org/column/article/2751
