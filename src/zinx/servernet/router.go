package servernet

import "Zserver/src/zinx/serverinterface"

type BaseRouter struct{}

// 隔离接口, 自定义router嵌入BaseRouter之后只需要实现部分函数即可
func (b BaseRouter) BeforeHandle(request serverinterface.IRequest) {}

func (b BaseRouter) Handle(request serverinterface.IRequest) {}

func (b BaseRouter) AfterHandle(request serverinterface.IRequest) {}
