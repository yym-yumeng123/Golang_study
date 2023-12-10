package http

import (
	"encoding/json"
	"io"
	"net/http"
)

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

func (c Context) BadRequestJson(resp any) error {
	return c.WriteJson(http.StatusBadRequest, resp)
}
