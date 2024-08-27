�����ᵽ�Ĳ���"IDs=1,2,3,4,5"����Ҫ����������ŷָ���ַ���ת��Ϊint��Ƭ[]int{1,2,3,4,5}�����ں��������ʹ�á�<br>
���˼·��
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
���������<br>
ͨ��request.UserManyģ�ͽ���IDs�����������Զ��׼���������һ���ӿ�Ҳ��ҪIDs����ʱ�����Ӧ��requestģ��ͬ����Ҫʵ��IDsValues�����������͵��µ��ظ����롣

���˼·��
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
���������<br>
ͨ������������IDsField���ѽ������̷�װ���ڲ���������һ���ӿ���ҪIDs����ʱ�����ü��ɣ�<br>
���ַ�ʽ�ִ�������һ�����⣬�����û��ȼ�Level������Ҫ��ѯ����û��ȼ�������Ҫ���½�һ��LevelsField�ṹ�壬���£�
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
���ǰ�˼ά�ٳ���һ�㣬�����ַ���"1,2,3,4,5"�������½�һ��ͨ�õ�NumbersFieldV1�ṹ�壬���£�
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
�������û�ID���û��ȼ��ȶ����ֲ����Ϳ��Ը���NumbersField��

������뱾�ĵ��ص㣬����������һ��NumbersField.Values������Values����������Ŀ���ǰѱ����ŷָ��string��ת��Ϊ[]int����Ϊ�����裺<br>
����1���ַ���SplitΪ[]string<br>
����2��forѭ��[]string��ȡ��Ƭ����ַ���<br>
����3���ַ���ת��Ϊint��ת������ȽϿ��ɣ�ʧ�ܺ��ԣ�ֻ�ռ�ת���ɹ�������<br>

ĿǰΪֹ��������ûʲô���⣬����Ҫ�����ϸ����У�����ת��ʧ�ܾͱ����������£�
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
