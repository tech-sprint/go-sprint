# package websocket

> Package websocket implements the WebSocket protocol defined in [RFC 6455](https://tools.ietf.org/html/rfc6455).

```go
import "github.com/gorilla/websocket"
```

### [示例来源](https://github.com/gorilla/websocket/blob/master/examples/echo/README.md)

## 参考

* [failed: Error during WebSocket handshake: Unexpected response code: 403 OR upgrade:websocket: request origin not allowed by Upgrader.CheckOrigin](https://github.com/gorilla/websocket/issues/367)
* [跨域相关](https://godoc.org/github.com/gorilla/websocket#hdr-Origin_Considerations)

## 其它 websocket 实现

* [socket.io](https://github.com/socketio/socket.io): 也包含了服务端和客户端(包括浏览器端，封装了原生的 WebSocket API)。
