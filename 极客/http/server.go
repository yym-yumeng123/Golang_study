package http

import (
	"fmt"
	"net/http"
)

/**
 * interface 接口
 * 里面只能是方法
 * 接口是一组行为的抽象
 * 尽量用接口, 以实面向接口编程
 */
type Server interface {
	Routable
	// Start 启动我们的服务器
	Start(address string) error
}

// 结构体
type sdkHttpServer struct {
	Name    string
	handler Handler
}

// Route 注册路由
func (s *sdkHttpServer) Route(
	method string,
	pattern string,
	handleFunc func(ctx *Context)) {
	s.handler.Route(method, pattern, handleFunc)
}

func (s *sdkHttpServer) Start(address string) error {
	http.Handle("/", s.handler)
	return http.ListenAndServe(address, nil)
}

func NewHttpServer(name string) Server {
	return &sdkHttpServer{
		Name:    name,
		handler: NewHandlerBasedOnMap(),
	}
}

func SignUpWithoutContext(w http.ResponseWriter, r *http.Request) {
	req := &signUpReq{}
	ctx := Context{
		W: w,
		R: r,
	}
	err := ctx.ReadJson(req)
	if err != nil {
		fmt.Fprintf(w, "err %v", err)
	}

	resp := &commonResponse{}

	err = ctx.OkJson(resp)
	if err != nil {
		fmt.Printf("写入响应失败: %v", err)
	}

	fmt.Fprintf(w, "%d", 123)
}

type signUpReq struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type commonResponse struct {
	Data    any    `json:"data"`
	BizCode int    `json:"biz_code"`
	Msg     string `json:"msg"`
}
