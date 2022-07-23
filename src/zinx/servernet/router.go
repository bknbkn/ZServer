package servernet

import "Zserver/src/zinx/serverinterface"

// BaseRouter 隔离接口, 自定义router嵌入BaseRouter之后只需要实现部分函数即可
// Router是一类handle策略的集合
type BaseRouter struct{}

func (b BaseRouter) BeforeHandle(request serverinterface.IRequest) {}

func (b BaseRouter) Handle(request serverinterface.IRequest) {}

func (b BaseRouter) AfterHandle(request serverinterface.IRequest) {}
