参数校验是日常开发中不可或缺的一环，也是繁琐、效率低下的一环。

<img src="../images/login.jpg" width="50%">

#### 案例：
如上图为常见的网站登陆场景，用户第一步输入手机号，点击获取短信验证码；第二步输入手机收到的短信验证码，点击登录按钮完成登录。<br>
```
发送验证码接口：
判断手机号非空，手机号数字型，手机号格式是否正确

登录接口：
判断手机号非空，手机号数字型，手机号格式是否正确
判断验证码非空，验证码数字型，验证码位数是否正确
```

V1版本：
```go
// controller.Sms
func (ctr *Sms) Captcha(c *gin.Context) (response.Data, error) {
    mobile := c.DefaultQuery("mobile", "")
    
    if mobile == "" {
        return nil, errors.New("手机号不能为空")
    }
    
    if matched, _ := regexp.MatchString(`^[0-9]+$`, mobile); !matched {
        return nil, errors.New("手机号必须数字")
    }
    
    if matched, _ := regexp.MatchString(`^(1[3-9][0-9]\d{8})$`, mobile); !matched {
        return nil, errors.New("手机号格式不正确")
    }
    
    return gin.H{"Mobile": mobile}, nil
}

// controller.User
func (ctr *User) Login(c *gin.Context) (response.Data, error) {
    mobile := c.DefaultQuery("mobile", "")
    captcha := c.DefaultQuery("captcha", "")
    
    if mobile == "" {
        return nil, errors.New("手机号不能为空")
    }
    
    if matched, _ := regexp.MatchString(`^[0-9]+$`, mobile); !matched {
        return nil, errors.New("手机号必须数字")
    }
    
    if matched, _ := regexp.MatchString(`^(1[3-9][0-9]\d{8})$`, mobile); !matched {
        return nil, errors.New("手机号格式不正确")
    }
    
    if captcha == "" {
        return nil, errors.New("验证码不能为空")
    }
    
    if matched, _ := regexp.MatchString(`^[0-9]+$`, captcha); !matched {
        return nil, errors.New("验证码必须数字")
    }
    
    if len(captcha) != 4 {
        return nil, errors.New("验证码必须4位")
    }
    
    return gin.H{"Mobile": mobile, "Captcha": captcha}, nil
}
```
[V1源码链接](../chapter-param-validate-1)

### 编程第一原则：DRY
`DRY`，不要重复自己(`Don’t repeat yourself`)，旨在减少代码的冗余和重复，提高代码的可维护性和复用性。

遵守DRY原则的优势：<br>
可维护：<br>
DRY原则的最大好处是可维护，DRY原则要求在编写代码时避免重复的逻辑和功能。如果同样的代码片段在多个地方出现，一旦需要对其中一个进行修改，就需要在其他地方都做相同的修改，增加了维护的困难度。

可复用：<br>
DRY原则鼓励将代码模块化，以便在不同的场景中重复使用，本质上促进了代码的复用。

可测试：<br>
模块化的代码很容易进行单元测试。

与之对应的就是`WET(Write Everything Twice)`，V1版本的代码就很湿，湿度100%，重复代码无处不在。

聪明的开发者肯定第一时间想到把校验逻辑提取成一个个独立的校验函数，按照此优化思路来到了V2版本：
```go
// controller.sms.go
func (ctr *Sms) Captcha(c *gin.Context) (response.Data, error) {
    mobile := c.DefaultQuery("mobile", "")
    
    if !notEmpty(mobile) {
        return nil, errors.New("手机号不能为空")
    }
    
    if !isNumber(mobile) {
        return nil, errors.New("手机号必须数字")
    }
    
    if !isMobile(mobile) {
        return nil, errors.New("手机号格式不正确")
    }
    
    return gin.H{"Mobile": mobile}, nil
}

// controller.user.go
func (ctr *User) Login(c *gin.Context) (response.Data, error) {
    mobile := c.DefaultQuery("mobile", "")
    captcha := c.DefaultQuery("captcha", "")
    
    if !notEmpty(mobile) {
        return nil, errors.New("手机号不能为空")
    }
    
    if !isNumber(mobile) {
        return nil, errors.New("手机号必须数字")
    }
    
    if !isMobile(mobile) {
        return nil, errors.New("手机号格式不正确")
    }
    
    if !notEmpty(captcha) {
        return nil, errors.New("验证码不能为空")
    }
    
    if !isNumber(captcha) {
        return nil, errors.New("验证码必须数字")
    }
    
    if !length(captcha, 4) {
        return nil, errors.New("验证码必须4位")
    }
    
    return gin.H{"Mobile": mobile, "Captcha": captcha}, nil
}

// controller.validate.go
func notEmpty(s string) bool {
    return s != ""
}

func isMobile(s string) bool {
    matched, _ := regexp.MatchString(`^(1[3-9][0-9]\d{8})$`, s)
    return matched
}

func length(s string, value int) bool {
    return len(s) == value
}

func isNumber(s string) bool {
    matched, _ := regexp.MatchString(`^[0-9]+$`, s)
    return matched
}
```
[V2源码链接](../chapter-param-validate-2)

`变量 != ""`和`len(变量) != num`如此简单的逻辑有必要提取成`notEmpty`和`length`函数吗？

非常有必要，原因有二：<br>
1、开发中有几率错误的把`!=`写成`==`，`==`写成`!=`，每个程序员应该都犯过这种错误。<br>
2、函数名能更好的表达代码意图

代码分析：<br>
1、感官清爽了很多，校验逻辑抽取成独立函数，提高了代码的可维护性和可复用性；<br>
2、`notEmpty`、`isMobile`、`length`、`isNumber`函数可测试性非常高；<br>

存在问题：<br>
1、存在`errors.New`重复的错误描述代码；

V2版本代码：湿度60%、干度40%

把`errors.New`独立出来，来到了V3版本：
```go
// controller.sms.go
func (ctr *Sms) Captcha(c *gin.Context) (response.Data, error) {
    mobile := c.DefaultQuery("mobile", "")
    
    if !notEmpty(mobile) {
        return nil, ErrMobileNotEmpty
    }
    
    if !isNumber(mobile) {
        return nil, ErrMobileNotNumber
    }
    
    if !isMobile(mobile) {
        return nil, ErrMobileFormat
    }
    
    return gin.H{"Mobile": mobile}, nil
}


// controller.user.go
func (ctr *User) Login(c *gin.Context) (response.Data, error) {
    mobile := c.DefaultQuery("mobile", "")
    captcha := c.DefaultQuery("captcha", "")
    
    if !notEmpty(mobile) {
        return nil, ErrMobileNotEmpty
    }
    
    if !isNumber(mobile) {
        return nil, ErrMobileNotNumber
    }
    
    if !isMobile(mobile) {
        return nil, ErrMobileFormat
    }
    
    if !notEmpty(captcha) {
        return nil, ErrCaptchaNotEmpty
    }
    
    if !isNumber(captcha) {
        return nil, ErrCaptchaNotNumber
    }
    
    if !length(captcha, 4) {
        return nil, ErrCaptchaLength
    }
    
    return gin.H{"Mobile": mobile, "Captcha": captcha}, nil
}

// controller.validate.go
var (
    ErrMobileNotEmpty  = errors.New("手机号不能为空")
    ErrMobileNotNumber = errors.New("手机号必须数字")
    ErrMobileFormat    = errors.New("手机号格式不正确")
    
    ErrCaptchaNotEmpty  = errors.New("验证码不能为空")
    ErrCaptchaNotNumber = errors.New("验证码必须数字")
    ErrCaptchaLength    = errors.New("验证码必须4位")
)
```
[V3源码链接](../chapter-param-validate-3)

代码分析：<br>
1、消除了`errors.New`重复的错误描述代码<br>

存在问题：<br>
1、Err系列变量存在大量的重复片段，3个"手机号"、3个"验证码"、2个"不能为空"、2个"必须数字"<br>
2、存在大量重复`if`控制语句

V3版本代码：湿度50%、干度50%

疑问：<br>
`if`控制语句不都这样写吗，会有什么问题？

这里的if控制语句重复是指：<br>
```go
// controller.sms.go
if !notEmpty(mobile) {
    return nil, ErrMobileNotEmpty
}
// controller.user.go
if !notEmpty(mobile) {
    return nil, ErrMobileNotEmpty
}

if !notEmpty(captcha) {
    return nil, ErrCaptchaNotEmpty
}
// `if !notEmpty(***)` 这3个if控制语句重复


// controller.sms.go
if !isNumber(mobile) {
    return nil, ErrMobileNotNumber
}
// controller.user.go
if !isNumber(mobile) {
    return nil, ErrMobileNotNumber
}

if !isNumber(captcha) {
    return nil, ErrCaptchaNotNumber
}
// `if !isNumber(***)` 这3个if控制语句重复


// controller.sms.go
if !isMobile(mobile) {
    return nil, ErrMobileFormat
}
// controller.user.go
if !isMobile(mobile) {
    return nil, ErrMobileFormat
}
// `if !isMobile(***)` 这2个if控制语句重复
```
重复写`notEmpty isNumber isMobile` `if`代码块的时候很有可能会遗漏`!`符号，把`if`控制语句抽取出来就可以避免这种低级错误。

来到了V4版本:
```go
// controller.sms.go
func (ctr *Sms) Captcha(c *gin.Context) (response.Data, error) {
    mobile := c.DefaultQuery("mobile", "")
    
    fields := []*Field{
        {name: "手机号", funcs: []string{NotEmpty, IsNumber, IsMobile}, value: mobile},
    }
    
    if err := validate(fields); err != nil {
        return nil, err
    }
    
    return gin.H{"Mobile": mobile}, nil
}

// controller.user.go
func (ctr *User) Login(c *gin.Context) (response.Data, error) {
    mobile := c.DefaultQuery("mobile", "")
    captcha := c.DefaultQuery("captcha", "")
    
    fields := []*Field{
        {name: "手机号", funcs: []string{NotEmpty, IsNumber, IsMobile}, value: mobile},
        {name: "验证码", funcs: []string{NotEmpty, IsNumber, Length}, value: captcha, args: []any{4}},
    }
    
    if err := validate(fields); err != nil {
        return nil, err
    }
    
    return gin.H{"Mobile": mobile, "Captcha": captcha}, nil
}

// controller.validate.go
type Validator func() bool

func (field *Field) validate() (errs Errs) {
    for _, funcName := range field.funcNames {
        vf, ok := ValidatorFuncMapping[funcName]
        if !ok {
            errs = append(errs, fmt.Errorf(ErrValidateFuncNotExistMsg, field.name))
            continue
        }
        
        validator, errMsg := vf(field)
        // 前文中所有的if控制语句都抽取成这一个if代码块了，并且重复使用也遗漏不了!符号
        if !validator() {
            errs = append(errs, errors.New(errMsg))
        }
    }
    return
}
```
[V4源码链接](../chapter-param-validate-4)

代码分析：<br>
1、通过可配置的方式实现了参数检验功能<br>
2、错误描述语句得到了重用<br>
3、`if`控制语句得到了重用

V4版本代码：湿度10%、干度90%

此时的参数校验代码就很像gin框架自带的模型绑定校验了，我们用gin模型绑定校验重新实现一遍，来到了V5版本：
```go
type SmsCaptcha struct {
    Mobile string `binding:"required,number,mobile"`
}

type UserLogin struct {
    Mobile  string `binding:"required,number,mobile"`
    Captcha string `binding:"required,number,len=4"`
}

func (ctr *Sms) Captcha(c *gin.Context) (response.Data, error) {
    request := new(request.SmsCaptcha)
    if err := c.ShouldBind(request); err != nil {
        return nil, err
    }
    
    return request, nil
}

func (ctr *User) Login(c *gin.Context) (response.Data, error) {
    request := new(request.UserLogin)
    if err := c.ShouldBind(request); err != nil {
        return nil, err
    }
    
    return request, nil
}
```
[V5源码链接](../chapter-param-validate-5)

此时的代码非常清爽，令人愉悦。

这么优雅的代码还会有什么可优化空间吗？<br>
`SmsCaptcha.Mobile`字段和`UserLogin.Mobile`字段`binding`标签内容重复

解决字段校验配置重复的解决办法是拆解字段为独立结构体，多个字段结构体的相互组合成业务所需的结构体，来到了V6版本：
```go
type MobileField struct {
    Mobile string `binding:"required,number,mobile"`
}

type CaptchaField struct {
    Captcha string `binding:"required,number,len=4"`
}

type IDField struct {
    ID int `binding:"required,min=1"`
}

type SmsCaptcha struct {
    MobileField
}

type UserLogin struct {
    MobileField
    CaptchaField
}

type UserRegister struct {
    MobileField
    CaptchaField
    user.NicknameField
    user.PasswordField
    user.RePasswordField
}
```
[V6源码链接](../chapter-param-validate-6)

代码分析：<br>
1、独立的字段结构体通常以表名为包名定义范围，比如商品名称和用户名称字段名都为Name，但是所需定义的校验逻辑（字符长度等）很有可能不同；<br>
2、复用Mobile、Captcha、ID、CreateTime等公共字段结构体；

V6版本代码：湿度0%、干度100%

总结：<br>
一、验证逻辑封装在各自的属性结构体中，由request层实体负责验证逻辑，验证逻辑不会散落在项目代码的各个地方，当验证逻辑改变时，修改对应的结构体就可以了，这就是代码的高内聚；

二、通过不同字段结构体的组合就可以实现多样的验证需求，提高了代码的复用率，这就是代码的可重用性

独立字段结构体组合成不同的校验结构体，这种方式在实际项目开发中有很大的灵活性，可以满足参数校验比较多变复杂的需求场景，小伙伴可以在项目开发中慢慢体会。

1、校验结构体为什么没有form标签？
```go
type UserLogin struct {
    Mobile  string `form:"mobile" binding:"required,number,mobile"`
    Captcha string `form:"captcha" binding:"required,number,len=4"`
}
```
form标签目的是为了绑定指定名称的from表单参数，通常是struct结构体属性名称首字母小写，这里也属于变相的重复代码，当我们通过复制这一行代码新建一个新的属性的时候，就需要改两个地方，实际开发过程中笔者真的发生过几次把form标签漏改的情况，导致bug耽误了时间，这种bug还特别让人恼火。

为了一劳永逸直接把from标签干掉，默认的struct结构体属性名称做为参数名称，省时又方便，当然习惯了首字母小写的同学感觉别扭，请克服心理障碍，可以参考阿里的接口文档：[链接](https://help.aliyun.com/document_detail/472990.html)，请求参数返回参数都保持首字母大写的格式，希望可以帮助你打败习惯的心魔。

2、参数校验失败的错误提示是否需要汉化？<br>
错误提示理论上用户是看不到的，因为前端也需要检验用户输入，用户看到的是前端代码实现的提示信息，后端的校验提示主要是给对接接口的程序员看的，转换为中文意义不大，建议是不要转，无端增加工作量，供大家参考。

3、gin自带的go-playground/validator结构体校验框架功能异常强大，能满足开发中99.99%的校验需求，请仔细学习。

### 控制逻辑业务逻辑傻傻分不清楚
```go
if !notEmpty(mobile) {
    return nil, errors.New("手机号不能为空")
}
```
在这个代码块中，`if ! {}`是控制逻辑，`notEmpty`是业务逻辑，V4版本分离了控制逻辑和业务逻辑，所有的校验业务逻辑抽象成Validator类型，`if ! {}`控制逻辑得到了复用。

```go
type Validator func() bool

// 前文中所有的if控制语句都抽取成这一个if代码块了，并且重复使用也遗漏不了!符号
if !validator() {
    errs = append(errs, errors.New(errMsg))
}
```
gin的validator框架就把分离控制逻辑业务逻辑做到了极致，它抽取了常见的校验业务逻辑，形成了一个个独立的函数达到代码复用的目的。

分离控制逻辑业务逻辑带来的好处：<br>
1、分离后的业务逻辑代码更少，没有了控制逻辑代码的干扰代码更好理解，降低了代码复杂度

2、控制逻辑大多数情况大同小异，复用控制逻辑，降低了手写出错几率，提升开发效率又减少bug

控制逻辑与业务逻辑的知识点，请参看陈皓先生的文章： [编程的本质](https://time.geekbang.org/column/article/2751)

### 分离关注点
关注点分离（Separation of Concerns，简称SOC）是软件工程中的一个设计原则，它鼓励将一个复杂的问题分解成多个更小的、更易于管理的部分。每个部分解决问题的一个特定方面，即一个“关注点”。通过这种方式，关注点分离旨在提高软件的可维护性、可扩展性和可复用性，同时减少代码的复杂度。

在参数校验的业务中，对参数使用哪些校验规则是开发者的关注点，而不是一直不停的写`if`判断语句；使用gin的validator框架，我们的关注点是配置校验`tag`标签，控制逻辑和业务逻辑都可以复用，从而提高了开发效率。

这篇文章的重点并不是讲如何使用gin框架校验参数，不要只是觉得gin参数检验功能很方便，而是学习其中包含的编程思想，从而提高自己的开发水平和开发效率，DRY、分离Control、Logic、SOC后续还会继续深入学习，请大家保持耐心。

开发中经常会出现格式如`1,2,3,4,5`的参数，那么我们又要如何优雅的解析呢，请看下文分解。
