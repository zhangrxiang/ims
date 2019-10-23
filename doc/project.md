# 资源

## 定义
```go
package models

import "time"

type ProjectModel struct {
	ID        int       `json:"id" gorm:"primary_key;AUTO_INCREMENT"`
	PHId      int       `json:"ph_id"`
	Name      string    `json:"name" gorm:"not null"`
	Desc      string    `json:"desc" gorm:"not null"`
	UserId    int       `json:"user_id"`
	CreatedAt time.Time `json:"created_at"`
}
```


## 接口 

### 添加资源
`/api/v1/resource/add`

#### 方法
`post`

#### 参数:  

##### header
```text
Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVc2VySUQiOjEsIlVzZXJuYW1lIjoiYWRtaW4iLCJleHAiOjE1NjYwMTAxODcsImlzcyI6ImlyaXMifQ.BEkfqhgvj8jqOgIkCQYHLY0cQI0anA5_DrM7ybRALlU
```
```
name string
desc string
type string
version string
file file
```

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
