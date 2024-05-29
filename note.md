工作区模式:
创建 workspace：
go work init

workspace 新增 Go Module：
go work use app console pkg
go work use -r github.com

运行：
go run ./app




go build -tags=jsoniter .
https://github.com/json-iterator/go