������ģʽ:
���� workspace��
go work init

workspace ���� Go Module��
go work use app console pkg
go work use -r github.com

���У�
go run ./app




go build -tags=jsoniter .
https://github.com/json-iterator/go






- ��ȥ�ո� ֧�����˺�ֻ�ܰ�һ���˺ţ�
- 4����Ŀ�ֲ��ģ��ת���ǵ��¶�
- 5��ʹ��pool���Ż�ģ��ת����ʵ���ԣ�
- 6��request��response�ֲ�ģ�����Ӧ�õ�ʵ�ʰ���  ҵ��ģ�������ݿ�ģ�͵����� � �Ƿ�ʼ ������ �Ƿ���� ����ʱ



- 11����Ŀ�㼶�Ļ������
- 12������ѩ������������͸��������Ⱦ
- 13����Ŀ�ṹ�ֲ���Զ�������
- 14����Ŀ�ṹʵ���е����ƣ�user_id,�ظ��ύ**


���������ѵ㣺������

���һ�� Ӧ�����µ�CodeReview

����ʼ������־

MySQL������뼶�� �ۿ��ĺ��ʼ��� ΪʲôMySQL Ĭ�ϸ��뼶����RR���ֱ���������ΪRC    RC������Ҫ����Ա�ֹ������
SELECT @@global.transaction_isolation


console.log(new Date("2024-07-03T08:00:10.325Z").toLocaleString())
console.log(new Date().toISOString())


1�����ORM�����ٶ�˹�ͨ�ɱ�������Ч�ʵ�����дSQL
2�����Ʒ�ת�Ż���ɾ�Ĳ�
3�����ò���У��ģ�ͣ������ظ�����
4���Զ�����ģ��ת����������ߴ��뿪��Ч�ʺ�ִ��Ч��
5��ģ�ͳػ�������GC���������ͷ���������
6����Ŀ�ֲ�����Զ���
7���޷��֪��cache




<img src="../images/postman-complex-parameters.jpg">

����Ƭ���ظ�
�������ظ�
1����װ�ɷ�����ȫ�ֱ�������
2����װ�ɶ�������
�ṹ���ڿ緽���ظ�
1����װ�ɶ�������
2����װ�ɰ�ȫ�ֱ�������
3����װ�ɶ����ṹ��
��ṹ���ظ�
1����װ�ɰ�ȫ�ֱ�������
2����װ�ɶ����ṹ��


TZ="Asia/Shanghai" go run main.go

https://wangbjun.site/2021/coding/golang/event-bus.html


ʵ���ϣ�����Ӧ��ϵͳ���ԣ�ֻ���������͵��쳣��
BizException��ҵ���쳣������ȷ��ҵ�����壬����Ҫ��¼Error��־������ҪRetry
SysException����֪��ϵͳ�쳣����Ҫ��¼Error��־������Retry
Exception��δ֪�������쳣����Ҫ������Error Stack��־������Retry


clear�ܹ�demo��
https://github.com/manuelkiessling/go-cleanarchitecture
https://github.com/bxcodec/go-clean-arch

https://github.com/xxjwxc/uber_go_guide_cn

https://darjun.github.io/2020/04/01/godailylib/govaluate/

���ߺ�����
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
Filter:{"ID":1,"Nickname":"�ǳ�","Amount":{"Start":10, "End":100},"Level":[10,20,30],"Status":["normal","cancel","invalid"],"CreateTime":{"Start":"2024-01-01T00:00:00Z", "End":"2024-10-01T00:00:00Z"}}

