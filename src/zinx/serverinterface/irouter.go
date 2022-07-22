package serverinterface

type IRouter interface {
	BeforeHandle(request IRequest)
	Handle(request IRequest)
	AfterHandle(request IRequest)
}
