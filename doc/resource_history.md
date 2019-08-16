# 资源历史记录

## 定义
```go
package models

type ResourceHistoryModel struct {
	ID         int    `json:"id" gorm:"primary_key;AUTO_INCREMENT"`
	ResourceID int    `json:"resource_id" gorm:"not null"`
	File       string `json:"file"`
	Path       string `json:"path"`
	Hash       string `json:"hash"`
	Version    string `json:"version" gorm:"not null"`
}
```

## 接口

### 添加资源分类
`/api/v1/resource-history/lists`

#### 方法
`get`

#### 参数:
##### header
```text
Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVc2VySUQiOjEsIlVzZXJuYW1lIjoiYWRtaW4iLCJleHAiOjE1NjYwMTAxODcsImlzcyI6ImlyaXMifQ.BEkfqhgvj8jqOgIkCQYHLY0cQI0anA5_DrM7ybRALlU
```
```text
resource_id string
```

#### 返回:

```json
{
    "success": true,
    "err_msg": "",
    "data": {
        "resources": [
            {
                "id": 6,
                "resource_id": 64,
                "file": "debug-raw.html",
                "path": "uploads/2019/08/debug-raw.htmla4d2e12a763f446f27e44bfef5a0680c.html",
                "hash": "a4d2e12a763f446f27e44bfef5a0680c",
                "version": "123"
            },
            {
                "id": 5,
                "resource_id": 64,
                "file": "debug-raw.html",
                "path": "uploads/2019/08/debug-raw.htmla4d2e12a763f446f27e44bfef5a0680c.html",
                "hash": "a4d2e12a763f446f27e44bfef5a0680c",
                "version": "123"
            },
            {
                "id": 4,
                "resource_id": 64,
                "file": "debug-raw.html",
                "path": "uploads/2019/08/debug-raw.htmla4d2e12a763f446f27e44bfef5a0680c.html",
                "hash": "a4d2e12a763f446f27e44bfef5a0680c",
                "version": "123"
            },
            {
                "id": 3,
                "resource_id": 64,
                "file": "debug-raw.html",
                "path": "uploads/2019/08/debug-raw.htmla4d2e12a763f446f27e44bfef5a0680c.html",
                "hash": "a4d2e12a763f446f27e44bfef5a0680c",
                "version": "123"
            },
            {
                "id": 1,
                "resource_id": 64,
                "file": "debug-raw.html",
                "path": "uploads/2019/08/debug-raw.htmla4d2e12a763f446f27e44bfef5a0680c.html",
                "hash": "a4d2e12a763f446f27e44bfef5a0680c",
                "version": "123"
            }
        ]
    }
}
```