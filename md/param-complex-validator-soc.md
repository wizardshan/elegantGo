�����ᵽ�Ĳ���"IDs=1,2,3,4,5"����Ҫ����������ŷָ���ַ���ת��Ϊint��Ƭ[]int{1,2,3,4,5}�����ں��������ʹ�á�<br>
���˼·��
```go
type UserMany struct {
    IDs string
}

func (req *UserMany) IDsValues() []int {
    ss := strings.Split(req.IDs, ",")
    var values []int
    for _, s := range ss {
        value, err := strconv.Atoi(s)
        if err != nil {
            continue
        }
        values = append(values, value)
    }
    return values
}
```
[Դ������](../chapter-param-complex-validator-1)<br>
���������<br>
ͨ��request.UserManyģ�ͽ���IDs�ַ��������������Զ��׼���������һ���ӿ�Ҳ��ҪIDs����ʱ�����Ӧ��requestģ��ͬ����Ҫʵ��IDsValues�����������͵��µ��ظ����롣

���˼·��
```go
type UserMany struct {
    IDs IDsField
}

type IDsField string

func (req *IDsField) Values() []int {
    ss := strings.Split(string(*req), ",")
    var values []int
    for _, s := range ss {
        value, err := strconv.Atoi(s)
        if err != nil {
            continue
        }
        values = append(values, value)
    }
    return values
}
```
���������<br>
ͨ������������IDsField���ѽ������̷�װ���ڲ���������һ���ӿ���ҪIDs����ʱ�����ü��ɣ�<br>
���ַ�ʽ�ִ�������һ�����⣬�����û��ȼ�Level������Ҫ��ѯ����û��ȼ�������Ҫ�ٸ���ճ��LevelsField�ṹ�壬���£�
```go
type LevelsField string

func (req *LevelsField) Values() []int {
    ss := strings.Split(string(*req), ",")
    var values []int
    for _, s := range ss {
        value, err := strconv.Atoi(s)
        if err != nil {
            continue
        }
        values = append(values, value)
    }
    return values
}
```
���ǰѽǶ��ٰθ�һ��˼ά�ٳ���һ�㣬�����ַ���"1,2,3,4,5"�������½�һ��ͨ�õ�IntsFieldV1�ṹ�壬���£�
```go
type UserMany struct {
    IDs    IntsFieldV1
    Levels IntsFieldV1
}

type IntsFieldV1 string

func (req *IntsFieldV1) Values() []int {
    ss := strings.Split(string(*req), ",")
    var values []int
    for _, s := range ss {
        value, err := strconv.Atoi(s)
        if err != nil {
            continue
        }
        values = append(values, value)
    }
    return values
}
```
�������û�ID���û��ȼ��ȶ��������ֲ����Ϳ��Ը���IntsFieldV1��

������뱾�ĵ��ص㣬����������һ��IntsFieldV1.Values������Values������Ŀ���ǰѱ����ŷָ��string��ת��Ϊ[]int����Ϊ�����裺<br>
����1���ַ���SplitΪ[]string<br>
����2��forѭ��[]string��ȡ��Ƭ����ַ���<br>
����3���ַ���ת��Ϊint��ת������ȽϿ��ɣ�ʧ�ܺ��ԣ�ֻ�ռ�����ת���ɹ�������<br>

��ʮ�д��뿴����ûʲô���⣬��Ī�����е����ģ�����һ��������ĸо�������˵�������ﲻ�ԣ����ǽ������¿���

�ַ���"1,2,3,4,5"ת��Ϊ[]int{1,2,3,4,5}�����ֹ���<br>
1��ֻת����Ч������<br>
2����Ч���ݱ���ֹͣת��<br>
3��������Ϊ����Ч����ת��ʧ�ܣ�Ĭ����ֵ<br>
�ַ���"1,s,2,,3"ת����Ӧ�Ľ����<br>
1��[]int{1,2,3}<br>
2������<br>
3��[]int{1,0,2,0,3}

��Ӧʵ�ֵĴ��룺
```go
// ֻת����Ч������
func (req *IntsFieldV1) Values() []int {
    ss := strings.Split(string(*req), ",")
    var values []int
    for _, s := range ss {
        value, err := strconv.Atoi(s)
        if err != nil {
            continue
        }
        values = append(values, value)
    }
    return values
}

// ��Ч���ݱ���ֹͣת��
func (req *IntsFieldV1) MustValues() ([]int, error) {
    ss := strings.Split(string(*req), ",")
    var values []int
    for _, s := range ss {
        value, err := strconv.Atoi(s)
        if err != nil {
            return nil, err
        }
        values = append(values, value)
    }
    return values, nil
}

// ������Ϊ����Ч����ת��ʧ�ܣ�Ĭ����ֵ
func (req *IntsFieldV1) ShouldValues() []int {
    ss := strings.Split(string(*req), ",")
    var values []int
    for _, s := range ss {
        value, _ := strconv.Atoi(s)
        values = append(values, value)
    }
    return values
}
```
��ҿ��Ի����������һ��������������˼��һ����û���Ż��ռ䡣

��Values��MustValues��ShouldValues����һ��Աȣ��о��������ˣ��������һ�£�ֻ�в��ֵ�΢С���������������ֶ����������ر��Ѹ㣬�������־����ԭ����ʲô�أ�

���Ƿ���һ��Values�����е�forѭ�����̶̵��������õ���ss��s��values��value��err�����������ʽת���������жϡ�ѭ��ֹͣ�жϡ������ռ��ȹ���������һ�𣬾����ǽ���ʽ���롣

forѭ��������̫��Ĺ��ܣ���Ҫ�������Ḻ����Values�����ı�����ֻ�ռ���Ч���ݣ��Ż�˼·�����Ȱ���Ч���ݹ��˵���Ȼ��ֻ����Ч���ݽ��и�ʽת���������ͱ����˸�ʽת����ͬʱ��Ҫ����������Ч���жϡ�

ͬ��MustValues����������Ҳ�Ƚ���������Ч���жϣ�������Ч����ֱ�ӷ��ش���

����������V2�汾��
```
type IntsFieldV2 string
    
func (req *IntsFieldV2) Values() []int {
    ss := req.split()
    // ���˳���Ч����
    var ssFiltered []string
    for _, s := range ss {
        if numeral.IsInt(s) {
            ssFiltered = append(ssFiltered, s)
        }
    }
    
    var values []int
    for _, s := range ssFiltered {
        value, _ := strconv.Atoi(s)
        values = append(values, value)
    }
    return values
}
    
func (req *IntsFieldV2) MustValues() ([]int, error) {
    ss := req.split()
    // Ԥ�Ƚ�����Ч���ж�
    for _, s := range ss {
        if !numeral.IsInt(s) {
            return nil, errors.New(s + " is not an integer")
        }
    }
    
    var values []int
    for _, s := range ss {
        value, _ := strconv.Atoi(s)
        values = append(values, value)
    }
    return values, nil
}
```
numeral.IsInt������ͨ���������ж��ַ����Ƿ���int���͡�

����ԭ����forѭ��ֻ��ʵ�ּ򵥵����ݸ�ʽת�����ܣ�ͬʱValues��MustValues��ShouldValues����������������ͬ����飺
```go
var values []int
for _, s := range ss {
    value, _ := strconv.Atoi(s)
    values = append(values, value)
}
```
�����Ȱ�����ת�����ܴ����ȡ������������һ������������������V3�汾��
```go
func (req *IntsFieldV3) toInt(s string) int {
    value, _ := strconv.Atoi(s)
    return value
}
```
��Ȼforѭ����ȡ��toInt������������������Ȼ��������ͬ����飺
```go
var values []int
for _, s := range ss {
    values = append(values, req.toInt(s))
}
```
���forѭ��������[]stringת��[]int��ʵ�ֹ��̣�һ������ת��������һ�����ͣ��ڱ�����и�רҵ���mapӳ�䣬���ǰ�ӳ����̳�ȡ����������������V4�汾��
```go
func (req *IntsFieldV4) mapper(ss []string, fn func(s string) int) []int {
    var values []int
    for _, s := range ss {
        values = append(values, fn(s))
    }
    return values
}
```
�ܶ��������ж��Դ�mapӳ�䷽����python��js��map��java��stream.map��php��array_map������go����û��ʵ�֣�����ʹ�õ������⣺pie [Դ������](https://github.com/elliotchance/pie)������������V5�汾��
```go
pie.Map(ss, req.toInt)
```
toInt����pie����Ҳ�Դ�������������V6�汾��
```go
// ���հ汾
return pie.Ints(ssFiltered)

// �汾��
//return pie.Map(ssFiltered, numeral.ToInt)

// �汾һ
//return pie.Map(ssFiltered, func(s string) int {
//	return pie.Int(s)
//})
```
�汾һ��pie.Int�Ƿ��ͺ�����pie.Map����ֱ��ʹ�ã���Ҫ����������װһ�㣻<br>
�汾�����汾һ�������ţ��ɴ˷�װ��numeral.ToInt������<br>
���հ汾��pie���Դ�[]stringת��[]int��map����pie.Ints���������ࡣ<br>

Values��������ι��˹��ܵĴ�������ʹ��pie.Filter����
```go
var ssFiltered []string
for _, s := range ss {
    if numeral.IsInt(s) {
        ssFiltered = append(ssFiltered, s)
    }
}
```
ͬʱMustValues�е�����ж�������Ч�Դ������ʹ��pie.All����
```go
for _, s := range ss {
    if !numeral.IsInt(s) {
        return nil, errors.New(s + " is not an integer")
    }
}
```
����������V7�汾��
```go
// ����
pie.Filter(ss, numeral.IsInt)

// �ж���Ч��
if !pie.All(ss, numeral.IsInt) {
    return nil, errors.New("one of numbers is not an integer")
}
```
�ж���Ч�Ի�����ʹ��gin�Դ���validator������Ҫ��numeral.ToIntע�ᵽvalidator���У�
```go
if err := validate.Var(ss, "dive,int"); err != nil {
    return nil, err
}
// dive��ǩ���ܺ�ǿ��������ĵ�ѧϰ
```
����������һ�£�ʡȥһЩ�м���������������հ汾��
```go
type IntsField string

func (req *IntsField) Values() []int {
    return pie.Ints(pie.Filter(req.split(), numeral.IsInt))
}

func (req *IntsField) MustValues() ([]int, error) {
    ss := req.split()
    if !pie.All(ss, numeral.IsInt) {
        return nil, errors.New("one of numbers is not an integer")
    }
    return pie.Ints(ss), nil
}

func (req *IntsField) ShouldValues() []int {
    return pie.Ints(req.split())
}

func (req *IntsField) split() []string {
    return strings.Split(string(*req), ",")
}
```
���հ汾����һЩ����ʽ��̵�ζ�����Ա�V1�汾��ϣ���Դ������������

��V1�汾���������հ汾����Ȼ���벻�࣬���漰����֪ʶ��ͦ��ģ����һ��ǳ���Ҫ����������ϸ������

### ��ע�����
��ע����루Separation of Concerns�����SoC������������е�һ�����ԭ����������һ�����ӵ�����ֽ�ɶ����С�ġ������ڹ���Ĳ��֡�ÿ�����ֽ�������һ���ض����棬��һ������ע�㡱��ͨ�����ַ�ʽ����ע�����ּ���������Ŀ�ά���ԡ�����չ�ԺͿɸ����ԣ�ͬʱ���ٴ���ĸ��Ӷȡ�

�����ǲ��ϵ����Ĺ����У���ע�����������Ҫ��ָ��˼�룬�ᴩʼ�ա�

��V1�汾Values�����е�forѭ���ʹ��ڹ���Ĺ�ע�㣺
```go
for _, s := range ss {               // ��ע��1��forѭ��
    value, err := strconv.Atoi(s)    // ��ע��2����ʽת��
    if err != nil {                  // ��ע��3������ת������
        continue                     // ��ע��4��ѭ��ֹͣ
    }
    values = append(values, value)   // ��ע��5�������ռ�
}
```
ͨ�������ע�㣬�������forѭ������һ��forѭ���ǹ�����Ч���ݣ��ڶ���forѭ��������ת��������for��ְ��ͱȽϵ�һ������ĸ��ӶȽ��ͣ��ɶ�����ߣ�������V2�汾��
```go
for _, s := range ss {                      // ��ע��1��forѭ��   
    if numeral.IsInt(s) {                   // ��ע��2����������   
        ssFiltered = append(ssFiltered, s)  // ��ע��3�������ռ�
    }
}

for _, s := range ssFiltered {              // ��ע��1��forѭ��
    value, _ := strconv.Atoi(s)             // ��ע��2����ʽת�� 
    values = append(values, value)          // ��ע��3�������ռ�
}
```
### ��������߼���ҵ���߼�
�����߼���ҵ���߼���֪ʶ�㣬��ο��ѹʳ�����������£� [��̵ı���](https://time.geekbang.org/column/article/2751)

�����߼���forѭ��<br>
ҵ���߼�����ʽת��

ͨ�������ע�㣬��ҵ���߼���ʽת��������ȡ�ɶ����ķ����ﵽ�������õ�Ŀ�ģ�����������V3�汾��
```go
for _, s := range ssFiltered {              // ��ע��1��forѭ��
    values = append(values, req.toInt(s))   // ��ע��2�������ռ�
}
```
��ȡ�����߼���ʵ��mapper����������������V4�汾��
```go
func (req *IntsFieldV4) mapper(ss []string, fn func(s string) int) []int {
    var values []int
    for _, s := range ss {
        values = append(values, fn(s))
    }
    return values
}
```
�����߼��ǿ��Ա�׼���ģ�����pie���е�Map���������������Լ���mapper����������������V5�汾��
```go
pie.Map(ss, req.toInt)
```
�򵥵�ҵ���߼�Ҳ�ǿ��Ա�׼���ģ�����string to int��[]string to int[]�ȵȣ�������V6�汾��
```go
pie.Ints(ss)
```
�����߼��еĹ��˲����͵����ж�����Ҳ�ǿ��Ա�׼���ģ�����������V7�汾��
```go
pie.Filter(ss, numeral.IsInt)

if !pie.All(ss, numeral.IsInt) {
    return nil, errors.New("one of numbers is not an integer")
}
```
����һ�����������հ汾��
```go
func (req *IntsField) Values() []int {
    return pie.Ints(pie.Filter(req.split(), numeral.IsInt))
}

func (req *IntsField) MustValues() ([]int, error) {
    ss := req.split()
    if !pie.All(ss, numeral.IsInt) {
        return nil, errors.New("one of numbers is not an integer")
    }
    return pie.Ints(ss), nil
}

func (req *IntsField) ShouldValues() []int {
    return pie.Ints(req.split())
}
```
����Values�����������Ĺ�ע���ǹ����������ҵ���߼���ʣ�µĶ����Ա�׼����

����MustValues�����������Ĺ�ע����������Ч���ж����ҵ���߼���ʣ�µĶ����Ա�׼����

����ShouldValues��������ע���ֻ������α�׼����...

��׼������ζ�Ÿ�Ч���������Ǳ�̣����и�ҵ������ˡ�

pie�����Ǹ��ֿ����߼���׼��ʵ�֣�Ҳ��һЩ������ӳ��map������ͬ��go�Դ�slices��Ҳʵ����һЩ��׼���������Ұ�pie��slices���е�ÿ��������ѧϰһ�飬�����Լ�������ҵ����룬˼����ʲô�����¿���Ӧ�á�

���򿪷����������¶��ڵ������ݣ��������ĵ�������Ҳ���������������Ա���ж���׼��
<img src="../images/mfr.jpg" width="100%">


Sort:ID
Order:ASC
Offset:20
Limit:10
Amount:100,1000
Level:10,50
Status:normal,cancel,invalid
CreateTime:2024-01-01 00:00:00,2024-05-01 23:59:59
UpdateTime:2024-02-01,2024-06-01
StartTime:2024-03-01,2024-07-01 23:59:59
Filter:{"ID":1,"Nickname":"admin","Amount":"100,1000","Level":"10,50","Status":"normal,cancel,invalid","CreateTime":"2024-01-01 00:00:00,2024-05-01 23:59:59","UpdateTime":"2024-02-01,2024-06-01","StartTime":"2024-03-01,2024-07-01 23:59:59"}

