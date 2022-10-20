## HTTP接口

```Header
    MUST: jwt: XXX; device_id
```

### 1 加入游戏

REQUEST

POST /v1/snake/join

```javascript
{
    "player_cnt": 2
}
```

RESPONSE

```json
{
    "code": 0,
    "msg": "ok",
    "status": "OK",
    "data": null
}
```

### 2 throw

REQUEST

GET /v1/snake/throw

RESPONSE

```json
{
    "code": 0,
    "msg": "ok",
    "status": 1, // 骰子点数
    "data": null
}
```

### 3 info (断线重连)

REQUEST

GET /v1/snake/game_info

RESPONSE

```json
{
    "code": 0,
    "msg": "ok",
    "status": {
      "turn_user_id": 1, // 到谁掷骰子
      "countdown": 5, // 倒计时
      "start_at": 362178369, //开始时间
      "users": [
        {
          "user_id": 1,
          "thumb": "", //头像
          "name": "name1",
          "status": "online/offline/abandon",
          "position": [1,2],
        },
        {
          "user_id": 2,
          "thumb": "",//头像
          "name": "name2",
          "status": "online/offline/abandon",
          "position": [1,2],
        }]

        }
    },
    "data": null
}
```

## WS

### 1 开始游戏（下行消息）

```json
{
  "user_id": 1,
  "type": "start",
  "game_id": "gyusdaihdjsa",
  // 发生事件的时间戳
  "timestamp": 362178369,
  "users": [
    {
      "user_id": 1,
      "thumb": "", //头像
      "name": "name1",
      "status": "online/offline/abandon",
      "position": [1,2],
    },
    {
      "user_id": 2,
      "thumb": "",//头像
      "name": "name2",
      "status": "online/offline/abandon",
      "position": [1,2],
    }
  
}
```

### 2 掷骰子结果（下行消息）

```json
{
  "type": "throw",
  "user_id": 1,
  "point": 2,
  "from_position": [1,2],
  "to_position": [1,4],
  "timestamp": 362178369,
}
```

### 3 发生位移（上梯子、被吞掉）（下行消息）

```json
{
  "type": "move",
  "timestamp": 362178369,
  "reason": "up/down",
  "user_id": 1,
  "from_position": [1,2],
  "to_position": [1,4]
}
```

### 4 通知玩家掷骰子（下行消息）

```json
{
  "type": "turn",
  "timestamp": 362178369,
  "user_id": 1,
}
```

### 5 到达重点游戏结束（下行消息）

```json
{
  "type": "over",
  "timestamp": 362178369,
  "winer": 1,
  "users": [
    {
      "user_id": 1,
      "thumb": "", //头像
      "name": "name1",
      "status": "online/offline/abandon",
      "position": [1,2],
      "score": 1, //分数
    },
    {
      "user_id": 2,
      "thumb": "",//头像
      "name": "name2",
      "status": "online/offline/abandon",
      "position": [1,2],
      "score": 0, //分数
    }]
}
```
