# Chess
 2023红岩后端考核项目：国际象棋

## API 参考

#### 注册

```http
  POST /user/register
```

| 参数         | 字段   | 类型     | 描述               |
|:-----------|:-----|:------- |:-----------------|
| `username` | body | `string` | **必选**. 用户名，不能重复 |
| `password` | body | `string` | **必选**. 用户密码     |
返回参数：

| 参数         | 类型     | 描述   |
|:-----------| :------- |:-----|
| `status` | `int` | 状态码  |
| `info` | `string` | 返回信息 |

成功返回示例：
```http
  {
    "status": 200,
    "info": "success"
  }
```

#### 登录

```http
  GET /user/login
```

| 参数         | 字段   | 类型     | 描述               |
|:-----------|:-----|:------- |:-----------------|
| `username` | body | `string` | **必选**. 用户名 |
| `password` | body | `string` | **必选**. 用户密码     |
返回参数：

| 参数              | 类型       | 描述             |
|:----------------|:---------|:---------------|
| `status`        | `int`    | 状态码            |
| `token`         | `string` | token，有效期12h   |
| `refresh_token` | `string` | 刷新token，有效期24h |

成功返回示例：
```http
{
  "refresh_token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOiIyMDIzLTA2LTEyIDE0OjU0OjE4LjEwNDE2NzYyMiArMDgwMCBDU1QgbT0rOTAzOTYuOTA3MDQ2OTc2IiwiaWQiOiIyIn0.UCzKCkrhnVOCY3eunSJFIHdjio3ZoB1sCkZLb8t3kbM",
  "status": 200,
  "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOiIyMDIzLTA2LTEyIDAyOjU0OjE4LjEwNDEzMTg2MyArMDgwMCBDU1QgbT0rNDcxOTYuOTA3MDExMjA3IiwiaWQiOiIyIn0.YPhnKSoEi33lezbaIZyBZjks44LDC9abOqcelDp_QHE"
}
```

#### 刷新token

```http
  GET /user/login/refresh
```

| 参数              | 字段   | 类型     | 描述              |
|:----------------|:-----|:------- |:----------------|
| `refresh_token` | Query | `string` | **必选**. 刷新token |

返回参数：

| 参数              | 类型       | 描述      |
|:----------------|:---------|:--------|
| `status`        | `int`    | 状态码     |
| `token`         | `string` | token   |
| `refresh_token` | `string` | 刷新token |

成功返回示例：
```http
{
    "refresh_token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOiIyMDIzLTA2LTEyIDE1OjA5OjEzLjcxMDU2ODgyMiArMDgwMCBDU1QgbT0rOTEyOTIuNTEzNDQ4MTc1IiwiaWQiOiIyIn0.qEfcN8RXIPik_hS-AR1mr1N-zJywysmXRsQnXCU2BMU",
    "status": 200,
    "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOiIyMDIzLTA2LTEyIDAzOjA5OjEzLjcxMDU0NDA3MSArMDgwMCBDU1QgbT0rNDgwOTIuNTEzNDIzNDE1IiwiaWQiOiIyIn0.RXTCSbIJ1vGlhG7RTGrscf7TKSfilObsDqaS75TTV_U"
}
```

#### 创建房间

```http
  POST /room/create
```

| 参数              | 字段     | 类型       | 描述            |
|:----------------|:-------|:---------|:--------------|
| `Authorization` | Header | `string` | **必选**. token |
| `room_name`     | body   | `string` | 房间名称          |

返回参数：

| 参数       | 类型    | 描述        |
|:---------|:------|:----------|
| `status` | `int` | 状态码       |
| `info`   | `int` | 随机八位的房间id |


成功返回示例：
```http
{
    "info": 24056115,
    "status": 200
}
```