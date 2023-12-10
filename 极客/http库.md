### HTTP Server - Request

- Request 请求

```go
func home(w http.ResponseWriter, r *http.Request) {
    // r.Body
    body, err := io.ReadAll(r.Body)
    // r.GetBody
    r.GetBody
    // r.URL
    // r.URL.Query()
    values := r.URL.Query()
    json.Marshal(r.URL)
    // r.Header
    r.Header
}
```

### HTTP Server - Server 和 Context

- Context 抽象与实现
  - 读取数据
  - 写入响应
- 创建 Context

```go
type Context struct {
	W http.ResponseWriter
	R *http.Request
}

func NewContext(writer http.ResponseWriter, request *http.Request) *Context {
    return &Context{
        W: writer,
        R: request,
    }
}

func (c Context) ReadJson(obj any) error {
	// 读取 body, 序列化
	body, err := io.ReadAll(c.R.Body)
	if err != nil {
		return err
	}
	err = json.Unmarshal(body, obj)
	if err != nil {
		return err
	}
	return nil
}

func (c Context) WriteJson(code int, resp interface{}) error {
	c.W.WriteHeader(code)
	respJson, err := json.Marshal(resp)
	if err != nil {
		return err
	}
	_, err = c.W.Write(respJson)
	return err
}

func (c Context) OkJson(resp any) error {
	return c.WriteJson(http.StatusOK, resp)
}

func (c Context) SystemErrorJson(resp any) error {
	return c.WriteJson(http.StatusInternalServerError, resp)
}
```