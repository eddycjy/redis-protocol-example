# redis-protocol-example

用 Go 来了解一下 Redis 通讯协议，本项目支持 Redis 的五种通讯协议，如下：

- 状态回复（status reply）
- 错误回复（error reply）
- 整数回复（integer reply）
- 批量回复（bulk reply）
- 多条批量回复（multi bulk reply）

## 安装
```
$ go get github.com/EDDYCJY/redis-protocol-example
```

## 使用

### Status Reply

```
$ go run main.go SET test01 value01
2018/06/06 21:29:07 Reply: OK
2018/06/06 21:29:07 Command: +OK
```

### Error Reply

```
$ go run main.go error
2018/06/06 22:20:39 Reply: ERR unknown command 'error'
2018/06/06 22:20:39 Command: -ERR unknown command 'error'
```

### Integer Reply

```
$ go run main.go EXPIRE test01 3600
2018/06/06 22:18:00 Reply: 1
2018/06/06 22:18:00 Command: :1
```

### Bulk Reply

```
$ go run main.go GET test01
2018/06/06 22:13:36 Reply: value01
2018/06/06 22:13:36 Command: $7
value01
```

### Multi Bulk Reply

```
$ go run main.go LPUSH test-multi 01
2018/06/06 22:23:50 Reply: 1
2018/06/06 22:23:50 Command: :1

$ go run main.go LPUSH test-multi 02
2018/06/06 22:23:54 Reply: 2
2018/06/06 22:23:54 Command: :2

$ go run main.go LPUSH test-multi 03
2018/06/06 22:23:57 Reply: 3
2018/06/06 22:23:57 Command: :3

$ go run main.go LRANGE test-multi 0 10
2018/06/06 22:24:10 Reply: [03 02 01]
2018/06/06 22:24:10 Command: *3
$2
03
$2
02
$2
01
```