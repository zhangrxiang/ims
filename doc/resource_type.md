# 资源分类

## 定义
```go
package models

import (
	"errors"
	"strings"
	"time"
)

type ResourceTypeModel struct {
	ID       int       `gorm:"primary_key;AUTO_INCREMENT" json:"id"`
	Name     string    `gorm:"unique;not null" json:"name"`
	Desc     string    `json:"desc" gorm:"not null"`
	CreateAt time.Time `json:"create_at"`
}
```

## 接口 

### 添加资源分类
`/api/v1/resource-type/add`

#### 方法
`post`

#### 参数:  
##### header
```text
Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVc2VySUQiOjEsIlVzZXJuYW1lIjoiYWRtaW4iLCJleHAiOjE1NjYwMTAxODcsImlzcyI6ImlyaXMifQ.BEkfqhgvj8jqOgIkCQYHLY0cQI0anA5_DrM7ybRALlU
```
```text
name string
desc string
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
  "err_msg": "保存资源分类失败",
  "success": false,
  "data": []
 }
```
##### 正确
```json
{
  "success": true,
  "err_msg": "",
  "data": {
    "resource_type": {
      "id": 17,
      "name": "周界三",
      "desc": "周界三周界三",
      "create_at": "2019-08-12T09:11:14.9487648+08:00"
    },
    "timestamp": 1565572274
  }
}

```

### 资源分类列表
`/api/v1/resource-type/lists`

#### 方法
`get`

#### 参数:  
##### header
```text
Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVc2VySUQiOjEsIlVzZXJuYW1lIjoiYWRtaW4iLCJleHAiOjE1NjYwMTAxODcsImlzcyI6ImlyaXMifQ.BEkfqhgvj8jqOgIkCQYHLY0cQI0anA5_DrM7ybRALlU
```
#### 返回:

```json
{
  "success": true,
  "err_msg": "",
  "data": {
    "resource_types": [
      {
        "id": 17,
        "name": "周界三",
        "desc": "周界三周界三",
        "create_at": "2019-08-12T09:11:14.9487648+08:00"
      },
      {
        "id": 16,
        "name": "周界二",
        "desc": "周界一周界一周界一周界一",
        "create_at": "2019-08-09T14:04:51.0038471+08:00"
      },
      {
        "id": 15,
        "name": "周界一",
        "desc": "周界一周界一周界一周界一",
        "create_at": "2019-08-09T14:02:32.8969676+08:00"
      },
      {
        "id": 14,
        "name": "周界1二",
        "desc": "周界一周界一周界一周界一",
        "create_at": "2019-08-09T10:26:09.4834511+08:00"
      }
    ],
    "timestamp": 1565572335
  }
}

```

### 资源分类删除
`/api/v1/resource-type/delete`

#### 方法
`get`

#### 参数:  
##### header
```text
Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVc2VySUQiOjEsIlVzZXJuYW1lIjoiYWRtaW4iLCJleHAiOjE1NjYwMTAxODcsImlzcyI6ImlyaXMifQ.BEkfqhgvj8jqOgIkCQYHLY0cQI0anA5_DrM7ybRALlU
```
```text
id string
id=1,2,3,4
```
 
#### 返回:

##### 错误
```json
{
  "success": false,
  "err_msg": "资源分类ID不能为空",
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
`/api/v1/resource-type/update`

#### 方法
`post`

#### 参数:  
##### header
```text
Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVc2VySUQiOjEsIlVzZXJuYW1lIjoiYWRtaW4iLCJleHAiOjE1NjYwMTAxODcsImlzcyI6ImlyaXMifQ.BEkfqhgvj8jqOgIkCQYHLY0cQI0anA5_DrM7ybRALlU
```
```text
id string
name string
desc string
```

#### 返回:

##### 错误
```json
{
  "err_msg": "资源分类ID不能为空",
  "success": false,
  "data": []
 }
```
或
```json
{
  "err_msg": "更新资源分类失败",
  "success": false,
  "data": []
 }
```
##### 正确
```json
{
  "success": true,
  "err_msg": "",
  "data": {
    "resource_type": {
      "id": 17,
      "name": "周界三",
      "desc": "周界三周界三",
      "create_at": "2019-08-12T09:11:14.9487648+08:00"
    },
    "timestamp": 1565572274
  }
}

```