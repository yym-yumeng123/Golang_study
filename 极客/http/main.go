package http

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "this is home page")
}

// Body
func readBodyOnce(w http.ResponseWriter, r *http.Request) {
	// io.ReadAll 把 r.Body 所有数据读
	// r.Body 只能读一次, 意味着你读了别人不能读, 别人读了你就不能读了
	body, err := io.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "read body failed: %v", err)
		return
	}
	// 类型转换, 将 []byte 转换成 string
	fmt.Fprintf(w, "read the data: %s \n", string(body))

	// 尝试再次读取 啥也读不到, 但是也不会报错
	body, err = io.ReadAll(r.Body)
	if err != nil {
		// 不会进来这里
		fmt.Fprintf(w, "read the data one more time get error: %v", err)
		return
	}
	fmt.Fprintf(w, "read the data one more time: [%s] and read data", body)
}

func getBodyIsNil(w http.ResponseWriter, r *http.Request) {
	// 原则上 可以多次读取, 但是在原生的 http.Request 里面, 这个是 nil
	if r.GetBody == nil {
		fmt.Fprintf(w, "GetBody is nil \n")
	} else {
		fmt.Fprintf(w, "GetBody not nil")
	}
}

// Request Query
// /api?name=yym&age=18 所有的值被解释为字符串
func queryParams(w http.ResponseWriter, r *http.Request) {
	values := r.URL.Query()
	fmt.Fprintf(w, "query is %v\n", values)
}

/**
 * Request URL
 * URL 里面 HOST 不一定有值
 * r.HOST 一般都幼稚, 是 HOST 这个header的值
 * RawPath 也是不一定有
 * Path 肯定有
 */
func wholeUrl(w http.ResponseWriter, r *http.Request) {
	data, _ := json.Marshal(r.URL)
	fmt.Fprintf(w, string(data))
}

/**
 * Request Header
 * type Header map[string][]string
 * 一般分两类: 一类是 http 预定义的, 一类是自定义的
 */
func header(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "header is %v\n", r.Header)
}

/**
 * Content-Type: application/x-www-form-urlencoded
 */
func form(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "before parse form %v\n", r.Form)
	err := r.ParseForm()
	if err != nil {
		fmt.Fprintf(w, "parse from error %v\n", r.Form)
	}
	fmt.Fprintf(w, "after from error %v\n", r.Form)
}

func SignUp(c *Context) {
	fmt.Println("SignUp")
}

func main() {
	server := NewHttpServer("test-server")
	//server.Route("/", home)
	//server.Route("/body/once", readBodyOnce)
	//server.Route("body/multi", getBodyIsNil)
	//server.Route("/url/query", queryParams)
	//server.Route("/header", header)
	//server.Route("/wholeUrl", wholeUrl)
	//server.Route("/form", form)
	server.Route(http.MethodGet, "/start", SignUp)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}
}
