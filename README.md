# gct

golang common tools

### 安装使用

安装

```shell
go get -u github.com/viodo/gct
``` 

使用

```
import (
    . "github.com/viodo/gct"
)
```
### 功能

#### 文件操作

判断文件是否存在

```golang
FileUtils.Exist("E:/test.txt")
```