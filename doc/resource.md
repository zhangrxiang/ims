# 资源

## 定义
```go
package models
import (
	"errors"
	"time"
)
type ResourceModel struct {
	ID       int       `json:"id" gorm:"primary_key;AUTO_INCREMENT"`
	Name     string    `json:"name" gorm:"not null"`
	Type     int       `json:"type" gorm:"not null"`
	File     string    `json:"file"`
	Path     string    `json:"path"`
	Hash     string    `json:"hash"`
	Version  string    `json:"version" gorm:"not null"`
	Desc     string    `json:"desc" gorm:"not null"`
	CreateAt time.Time `json:"create_at"`
}
```

## 接口 

### 添加资源分类
`/api/v1/resource/add`

#### 方法
`post`

#### 参数:  
`name string`
`desc string`
`type string`
`version string`
`file file`

#### 返回:

##### 错误
```json
{
  "err_msg": "请输入资源名称,选择资源类型",
  "success": false,
  "data": []
 }
```
或
```json
{
  "err_msg": "保存资源失败",
  "success": false,
  "data": []
 }
```
##### 正确
```json
{
    "success": true,
    "err_msg": "保存文件成功:",
    "data": {
        "resource": {
            "id": 29,
            "name": "ttttttttttt",
            "type": 12,
            "file": "background.js",
            "path": "uploads/2019/08/aed1670572864e61618a071ea06dca1a.js",
            "hash": "aed1670572864e61618a071ea06dca1a",
            "version": "123",
            "desc": "ttttttttttt",
            "create_at": "2019-08-12T08:54:22.658871+08:00"
        }
    }
}

```

### 资源列表
`/api/v1/resource/lists`

#### 方法
`get`

#### 参数:  

#### 返回:

```json
{
    "success": true,
    "err_msg": "",
    "data": {
        "resources": [
            {
                "id": 29,
                "name": "ttttttttttt",
                "type": 12,
                "file": "background.js",
                "path": "uploads/2019/08/aed1670572864e61618a071ea06dca1a.js",
                "hash": "aed1670572864e61618a071ea06dca1a",
                "version": "123",
                "desc": "ttttttttttt",
                "create_at": "2019-08-12T08:54:22.658871+08:00"
            },
            {
                "id": 28,
                "name": "ttttttttttt",
                "type": 12,
                "file": "script.js",
                "path": "uploads/2019/08/e68fb5cce661c5f4ae809ec8074c3756.js",
                "hash": "e68fb5cce661c5f4ae809ec8074c3756",
                "version": "123",
                "desc": "ttttttttttt",
                "create_at": "2019-08-12T08:38:27.6876902+08:00"
            },
            {
                "id": 27,
                "name": "ttttttttttt",
                "type": 12,
                "file": "styles.css",
                "path": "uploads/2019/08/dcd3a32a255906bfc1201330031a70e7.css",
                "hash": "dcd3a32a255906bfc1201330031a70e7",
                "version": "123",
                "desc": "ttttttttttt",
                "create_at": "2019-08-12T08:38:11.3276706+08:00"
            }
        ],
        "timestamp": 1565572908
    }
}
```

### 资源删除
`/api/v1/resource/delete`

#### 方法
`get`

#### 参数:  
`id string`
id=1,2,3,4
 
#### 返回:

##### 错误
```json
{
  "success": false,
  "err_msg": "资源ID不能为空",
  "data": []
}
```
##### 正确

```json
{
  "success": true,
  "err_msg": "",
  "data": []
}
```

### 修改资源分类
`/api/v1/resource/update`

#### 方法
`post`

#### 参数:  
`name string`
`desc string`
`type string`
`version string`
`file file`

#### 返回:

##### 错误
```json
{
  "err_msg": "资源ID不能为空",
  "success": false,
  "data": []
 }
```
或
```json
{
  "err_msg": "更新资源失败",
  "success": false,
  "data": []
 }
```
##### 正确
```json
{
    "success": true,
    "err_msg": "保存文件成功:",
    "data": {
        "resource": {
            "id": 29,
            "name": "ttttttttttt",
            "type": 12,
            "file": "background.js",
            "path": "uploads/2019/08/aed1670572864e61618a071ea06dca1a.js",
            "hash": "aed1670572864e61618a071ea06dca1a",
            "version": "123",
            "desc": "ttttttttttt",
            "create_at": "2019-08-12T08:54:22.658871+08:00"
        }
    }
}

```

### 修改资源分类
`/api/v1/resource/download`

#### 方法
`get`

#### 参数:  
`path string`
`file string`

#### 返回:

##### 错误
```json
{
  "err_msg": "文件路径不能为空",
  "success": false,
  "data": []
 }
```
##### 
文件下载了
