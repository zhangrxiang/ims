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
    "timestamp": 1565923787,
    "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVc2VySUQiOjEsIlVzZXJuYW1lIjoiYWRtaW4iLCJleHAiOjE1NjYwMTAxODcsImlzcyI6ImlyaXMifQ.BEkfqhgvj8jqOgIkCQYHLY0cQI0anA5_DrM7ybRALlU",
    "token_type": "Bearer",
    "user": {
      "id": 1,
      "username": "admin",
      "password": "******",
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

### 用户注册
`/api/v1/user/register`

#### 方法
`post`

#### 参数:  
```text
username:{{$timestamp}}
password:{{$timestamp}}
role:admin
mail:{{$timestamp}}@qq.com
phone:18800000000
```
#### 返回:

```json
{
    "success": true,
    "err_msg": "注册用户成功",
    "data": {
        "user": {
            "id": 20,
            "username": "1565601079",
            "password": "1565601079",
            "role": "admin",
            "phone": "18800000000",
            "mail": "1565601079@qq.com"
        }
    }
}
```

### 用户更新
`/api/v1/user/update`

#### 方法
`post`

#### 参数:

##### header
```text
Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVc2VySUQiOjEsIlVzZXJuYW1lIjoiYWRtaW4iLCJleHAiOjE1NjYwMTAxODcsImlzcyI6ImlyaXMifQ.BEkfqhgvj8jqOgIkCQYHLY0cQI0anA5_DrM7ybRALlU
```
```text
id:16
username:admin1
password:O(∩_∩)O
role:test
mail:1@qq.com
phone:18800000000
```
#### 返回:

```json
{
    "success": true,
    "err_msg": "修改用户成功",
    "data": {
        "user": {
            "id": 16,
            "username": "admin1",
            "password": "O(∩_∩)O",
            "role": "test",
            "phone": "18800000000",
            "mail": "1@qq.com"
        }
    }
}
```


### 用户删除
`/api/v1/user/delete`

#### 方法
`get`

#### 参数:  

##### header
```text
Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVc2VySUQiOjEsIlVzZXJuYW1lIjoiYWRtaW4iLCJleHAiOjE1NjYwMTAxODcsImlzcyI6ImlyaXMifQ.BEkfqhgvj8jqOgIkCQYHLY0cQI0anA5_DrM7ybRALlU
```

```text
id:16
```
#### 返回:

```json
{
    "success": false,
    "err_msg": "删除用户成功",
    "data": []
}
```