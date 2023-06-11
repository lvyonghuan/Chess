# Chess
 2023红岩后端考核项目：国际象棋

## API 参考

#### 注册

```
  POST /user/register
```

| 参数         | 字段   | 类型     | 描述               |
|:-----------|:-----|:------- |:-----------------|
| `username` | body | `string` | **必选**, 用户名，不能重复 |
| `password` | body | `string` | **必选**, 用户密码     |
返回参数：

| 参数         | 类型     | 描述   |
|:-----------| :------- |:-----|
| `status` | `int` | 状态码  |
| `info` | `string` | 返回信息 |

成功返回示例：
```json
  {
    "status": 200,
    "info": "success"
  }
```

#### 登录

```
  GET /user/login
```

| 参数         | 字段   | 类型     | 描述           |
|:-----------|:-----|:------- |:-------------|
| `username` | body | `string` | **必选**, 用户名  |
| `password` | body | `string` | **必选**, 用户密码 |
返回参数：

| 参数              | 类型       | 描述             |
|:----------------|:---------|:---------------|
| `status`        | `int`    | 状态码            |
| `token`         | `string` | token，有效期12h   |
| `refresh_token` | `string` | 刷新token，有效期24h |

成功返回示例：
```json
{
  "refresh_token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOiIyMDIzLTA2LTEyIDE0OjU0OjE4LjEwNDE2NzYyMiArMDgwMCBDU1QgbT0rOTAzOTYuOTA3MDQ2OTc2IiwiaWQiOiIyIn0.UCzKCkrhnVOCY3eunSJFIHdjio3ZoB1sCkZLb8t3kbM",
  "status": 200,
  "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOiIyMDIzLTA2LTEyIDAyOjU0OjE4LjEwNDEzMTg2MyArMDgwMCBDU1QgbT0rNDcxOTYuOTA3MDExMjA3IiwiaWQiOiIyIn0.YPhnKSoEi33lezbaIZyBZjks44LDC9abOqcelDp_QHE"
}
```

#### 刷新token

```
  GET /user/login/refresh
```

| 参数              | 字段   | 类型     | 描述              |
|:----------------|:-----|:------- |:----------------|
| `refresh_token` | Query | `string` | **必选**, 刷新token |

返回参数：

| 参数              | 类型       | 描述      |
|:----------------|:---------|:--------|
| `status`        | `int`    | 状态码     |
| `token`         | `string` | token   |
| `refresh_token` | `string` | 刷新token |

成功返回示例：
```json
{
    "refresh_token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOiIyMDIzLTA2LTEyIDE1OjA5OjEzLjcxMDU2ODgyMiArMDgwMCBDU1QgbT0rOTEyOTIuNTEzNDQ4MTc1IiwiaWQiOiIyIn0.qEfcN8RXIPik_hS-AR1mr1N-zJywysmXRsQnXCU2BMU",
    "status": 200,
    "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOiIyMDIzLTA2LTEyIDAzOjA5OjEzLjcxMDU0NDA3MSArMDgwMCBDU1QgbT0rNDgwOTIuNTEzNDIzNDE1IiwiaWQiOiIyIn0.RXTCSbIJ1vGlhG7RTGrscf7TKSfilObsDqaS75TTV_U"
}
```

#### 创建房间

```
  POST /room/create
```

| 参数              | 字段     | 类型       | 描述            |
|:----------------|:-------|:---------|:--------------|
| `Authorization` | Header | `string` | **必选**, token |
| `room_name`     | body   | `string` | 房间名称          |

返回参数：

| 参数       | 类型    | 描述        |
|:---------|:------|:----------|
| `status` | `int` | 状态码       |
| `info`   | `int` | 随机八位的房间id |


成功返回示例：
```json
{
    "info": 24056115,
    "status": 200
}
```
### 游戏相关操作
#### 加入房间
```
ws://your address/room/connect
```
| 参数              | 字段     | 类型       | 描述            |
|:----------------|:-------|:---------|:--------------|
| `Authorization` | Header | `string` | **必选**, token |
| `room_id`       | query  | `int`    | 房间ID          |
利用房间ID加入已经创建的房间。
#### 准备/取消准备
发送以下json结构体
```json
{
  "type": 1
}
```
第一次发送，玩家进入准备状态；第二次发送，玩家退出准备状态。第一个进入准备状态者执白棋。而后状态互斥。

当双方玩家同时进入准备状态时，游戏方可开始。
#### 移动棋子
发送以下json结构体
```json
{
  "type": 1,
  "move_before": [3, 4],
  "move_after": [5, 6]
}
```
游戏开始时，执白棋者先移动。每次移动后，返回如下结构：
```json
[[[3,1,0,0,0],[5,1,1,0,0],[4,1,1,0,0],[2,1,1,0,0],[1,1,1,0,0],[4,1,1,0,0],[5,1,1,0,0],[3,1,0,0,0]],[[6,1,1,0,0],[0,0,1,0,0],[6,1,1,0,0],[6,1,1,0,0],[6,1,1,0,0],[6,1,1,0,0],[6,1,1,0,0],[6,1,1,0,0]],[[0,0,1,0,0],[6,1,1,0,0],[0,0,1,0,0],[0,0,1,0,0],[0,0,1,0,0],[0,0,1,0,0],[0,0,1,0,0],[0,0,1,0,0]],[[0,0,1,0,0],[0,0,0,0,0],[0,0,1,0,0],[0,0,0,0,0],[0,0,0,0,0],[0,0,0,0,0],[0,0,0,0,0],[0,0,0,0,0]],[[0,0,0,0,0],[0,0,0,0,0],[0,0,0,0,0],[0,0,0,0,0],[0,0,0,0,0],[0,0,0,0,0],[0,0,0,0,0],[0,0,0,0,0]],[[0,0,0,1,0],[0,0,0,1,0],[0,0,0,1,0],[0,0,0,1,0],[0,0,0,1,0],[0,0,0,1,0],[0,0,0,1,0],[0,0,0,1,0]],[[6,2,0,1,0],[6,2,0,1,0],[6,2,0,1,0],[6,2,0,1,0],[6,2,0,1,0],[6,2,0,1,0],[6,2,0,1,0],[6,2,0,1,0]],[[3,2,0,0,0],[5,2,0,1,0],[4,2,0,1,0],[2,2,0,1,0],[1,2,0,1,0],[4,2,0,1,0],[5,2,0,1,0],[3,2,0,0,0]]]
```
这是一个8*8*5数组的第三维度。索引0代表了棋子类型，索引1代表了棋色，索引2与索引3分别代表了白黑方棋子能威胁到的格子，索引4是针对王和车的检测，用于判断是否可以进行王车易位。

关于棋子：
```
King    王 1
Queen   后 2
Rook    车 3
Bishop  象 4
Knight  马 5
Pawn    兵 6
```
关于棋色：
```
White 白棋 1
Black 黑棋 2
```
#### 兵的晋升
发送以下json结构体：
```json
{
  "type": 1,
  "upgrade": 1
}
```
upgrade字段代表了兵晋升的形态。参考上文的棋子编号。

## 目前效果
![img.png](img.png)