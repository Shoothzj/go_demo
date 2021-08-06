# go module 开发
## 初始化go module
```bash
go mod init {module名称}
```
## go module 添加依赖**
在go module路径下
```bash
go get github.com/prometheus/client_golang/prometheus
```
# release
## release new version
```bash
git tag $version
git push --tags
```
