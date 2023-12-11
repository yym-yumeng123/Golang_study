package http

import (
	"net/http"
	"strings"
)

// HandlerBaseOnTree 根节点
type HandlerBaseOnTree struct {
	root *node
}

// 子节点
type node struct {
	path     string
	children []*node
	// 如果是叶子节点
	// 匹配上之后就可以调用该方法
	handler HandlerFunc
}

func (h *HandlerBaseOnTree) ServeHTTP(c *Context) {
	handler, found := h.findRouter(c.R.URL.Path)
	if !found {
		c.W.WriteHeader(http.StatusNotFound)
		return
	}
	handler(c)
}

func (h *HandlerBaseOnTree) findRouter(path string) (HandlerFunc, bool) {
	paths := strings.Split(strings.Trim(path, "/"), "/")
	cur := h.root
	for _, p := range paths {
		matchChild, found := h.findMatchChild(cur, p)
		if !found {
			return nil, false
		}
		cur = matchChild
	}

	// 到这里, 找完了
	if cur.handler == nil {
	}
	return nil, true
}

func (h *HandlerBaseOnTree) Route(
	method string,
	pattern string,
	handleFunc HandlerFunc) {
	// 例 /user/friends
	pattern = strings.Trim(pattern, "/")
	paths := strings.Split(pattern, "/")

	cur := h.root

	for index, path := range paths {
		mathChild, ok := h.findMatchChild(cur, path)
		if ok {
			cur = mathChild
		} else {
			// 没找到, 创建
			h.createSubTree(cur, paths[index:], handleFunc)
			return
		}
	}
}

func (h *HandlerBaseOnTree) createSubTree(root *node, paths []string, handlerFunc HandlerFunc) {
	cur := root
	for _, path := range paths {
		nn := newNode(path)
		cur.children = append(cur.children, nn)
		cur = nn
	}
	cur.handler = handlerFunc
}

func (h *HandlerBaseOnTree) findMatchChild(root *node, path string) (*node, bool) {
	var wildcardNode *node
	for _, child := range root.children {
		// 并不是 * 的节点命中了, 直接返回
		// != * 是为了防止用户乱输入
		if child.path == path && child.path != "*" {
			return child, true
		}
		// 命中了通配符的, 我们看看后面很有没有更详细的
		if child.path == "*" {
			wildcardNode = child
		}
	}

	return wildcardNode, wildcardNode != nil
}

func newNode(path string) *node {
	return &node{
		path:     path,
		children: make([]*node, 0, 2),
	}
}
