package route

import "errors"

var (
	// ErrRouterUninitialized 路由未初始化
	ErrRouterUninitialized = errors.New("router uninitialized")
)
