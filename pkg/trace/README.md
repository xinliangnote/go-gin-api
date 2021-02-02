## trace

一个用于开发调试的辅助工具。

可以实时显示当前页面的操作的请求信息、运行情况、SQL执行、错误提示等。

- `trace.go` 主入口文件；
- `dialog.go` 处理 third_party_requests 记录；
- `debug.go` 处理 debug 记录；

#### 数据格式

##### trace_id

当前 trace 的 ID，例如：938ff86be98439c6c1a7，便于搜索使用。

##### request

请求信息，会包括：

- ttl 请求超时时间，例如：2s 或 un-limit
- method 请求方式，例如：GET 或 POST
- decoded_url 请求地址
- header 请求头信息
- body 请求体信息

##### response

- header 响应头信息
- body 响应提信息
- business_code 业务码，例如：10010
- business_code_msg 业务码信息，例如：签名错误 
- http_code HTTP 状态码，例如：200 
- http_code_msg HTTP 状态码信息，例如：OK
- cost_seconds 耗费时长：单位秒，比如 0.001105661

##### third_party_requests

每一个第三方 http 请求都会生成如下的一组数据，多个请求会生成多组数据。

- request，同上 request 结构一致
- response，同上 response 结构一致
- success，是否成功，true 或 false
- cost_seconds，耗费时长：单位秒

注意：response 中的 business_code、business_code_msg 为空，因为各个第三方返回结构不同，这两个字段为空。

##### sqls

执行的 SQL 信息，多个 SQL 会记录多组数据。

- timestamp，时间，格式：2006-01-02 15:04:05
- stack，文件地址和行号
- cost_seconds，执行时长，单位：秒
- sql，SQL 语句
- rows_affected，影响行数

##### debugs

- key 打印的标示
- value 打印的值

```cassandraql
// 调试时，使用这个方法：
p.Print("key", "value", p.WithTrace(c.Trace()))
```

只有参数中增加了 `p.WithTrace(c.Trace())`，才会记录到 `debugs` 中。

##### success

是否成功，true 或 false

```cassandraql
success = !ctx.IsAborted() && ctx.Writer.Status() == http.StatusOK
```

##### cost_seconds

耗费时长：单位秒，比如 0.001105661

