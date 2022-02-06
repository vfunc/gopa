# 网络编程作业

1. 总结几种 socket 粘包的解包方式：fix length/delimiter based/length field based frame decoder。尝试举例其应用。

* FixedLengthFrameDecoder 按固定长度的字节数解码
* DelimiterBasedFrameDecoder 按指定的分隔符解码
* LengthFieldBasedFrameDecoder 根据消息中长度字段的值动态解码

> 参考 netty Class ByteToMessageDecoder

2. 实现一个从 socket connection 中解码出 goim 协议的解码器。

> <https://goim.io/docs/protocol.html>  
> <https://www.tony.wiki/development/2016/09/04/goim-protocol.html>  
> <https://github.com/Terry-Mao/goim/blob/master/api/protocol/protocol.go>  
