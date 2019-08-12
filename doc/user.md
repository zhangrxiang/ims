# 用户

## 定义
```go
package models

type UserModel struct {
	ID       int    `json:"id",gorm:"primary_key;AUTO_INCREMENT"`
	Username string `json:"username",gorm:"not null;unique;type:varchar(30)"`
	Password string `json:"password",gorm:"not null;type:varchar(20)"`
	Role     string `json:"role"`
	Phone    string `json:"phone",gorm:"not null"`
	Mail     string `json:"mail",gorm:"not null"`
}
```

## 接口 

### 登陆
`/api/v1/user/login`

#### 方法
`get`

#### 参数:  
`username string`
`password string`

#### 返回:

##### 错误
```json
{
  "err_msg": "用户名或密码不能为空",
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
    "timestamp": 1565571887,
    "user": {
      "id": 1,
      "username": "zing",
      "password": "123456",
      "role": "",
      "phone": "",
      "mail": ""
    }
  }
}
```

### 用户列表
`/api/v1/user/lists`

#### 方法
`get`

#### 参数:  

#### 返回:

```json
{
  "success": true,
  "err_msg": "",
  "data": {
    "timestamp": 1565571998,
    "users": [
      {
        "id": 1,
        "username": "zing",
        "password": "123456",
        "role": "",
        "phone": "",
        "mail": ""
      },
      {
        "id": 2,
        "username": "admin",
        "password": "1234567",
        "role": "",
        "phone": "",
        "mail": ""
      }
    ]
  }
}
```