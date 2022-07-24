package serverinterface

type IMessageHandle interface {
	DoMsgHandler(request IRequest)
	AddRouter(msgId uint32, router IRouter)
	StartWorkerPool()
	AddMessageToQueue(request IRequest)
}
