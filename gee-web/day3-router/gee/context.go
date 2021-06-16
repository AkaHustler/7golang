package gee

import (
	"encoding/json"
	"fmt"
	"net/http"
)

/**
1.开头给map[string]interface{}起了个H的别名，构建JSON数据时，显得简洁
2.Context目前只包含 http.ResponseWriter 和 *http.Request，另外提供了对Method和Path的快捷访问
3.提供了访问Query和PostForm参数的方法
4.提供了快速构造String/Data/Json/Html响应的方法
 */

type H map[string]interface{}

// context struct
type Context struct {
	//origin objects
	Writer http.ResponseWriter
	Req *http.Request
	//request info
	Path string
	Method string
	//response info
	StatusCode int
}

func newContext(w http.ResponseWriter, req *http.Request) *Context {
	return &Context{
		Writer: w,
		Req: req,
		Path: req.URL.Path,
		Method: req.Method,
	}
}

//format form request
func (c *Context) PostForm(key string) string {
	return c.Req.FormValue(key)
}

//format query request
func (c *Context) Query(key string) string {
	return c.Req.URL.Query().Get(key)
}

//format status code
func (c *Context) Status(code int)  {
	c.StatusCode = code
	c.Writer.WriteHeader(code)
}

//set response header
func (c *Context) SetHeader(key string, value string)  {
	c.Writer.Header().Set(key, value)
}

//format text response
func (c *Context) String(code int, format string, values ...interface{})  {
	c.SetHeader("Content-Type", "text/plain")
	c.Status(code)
	c.Writer.Write([]byte(fmt.Sprintf(format, values...)))
}

//format json response
func (c *Context) Json(code int, obj interface{})  {
	c.SetHeader("Content-Type", "application/json")
	c.Status(code)
	encoder := json.NewEncoder(c.Writer)
	if err := encoder.Encode(obj); err != nil {
		http.Error(c.Writer, err.Error(), 500)
	}
}

//format data response
func (c *Context) Data(code int, data []byte)  {
	c.Status(code)
	c.Writer.Write(data)
}

//format html response
func (c *Context) Html(code int, html string)  {
	c.Status(code)
	c.Writer.Write([]byte(html))
}