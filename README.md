# CloudDisk
Slight cloud disk system , based on go-zero, xorm

## 1.Init program environment

```bash
go install github.com/zeromicro/go-zero/tools/goctl@latest

goctl --version

goctl api new core

go get github.com/zeromicro/go-zero/rest
 
go get github.com/zeromicro/go-zero/core/logx
 
go get github.com/zeromicro/go-zero/rest/httpx
 
go get github.com/zeromicro/go-zero/core/conf
 # run program
go run core.go -f etc/core-api.yaml
```
