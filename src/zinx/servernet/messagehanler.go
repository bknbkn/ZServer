package servernet

import (
	"Zserver/src/zinx/serverinterface"
	"Zserver/src/zinx/utils"
	"log"
)

type MessageHandle struct {
	MuxEntry map[uint32]serverinterface.IRouter

	// 消息队列
	TaskQueue      []chan serverinterface.IRequest
	WorkerPoolSize uint32
}

func NewMessageHandle() *MessageHandle {
	return &MessageHandle{
		MuxEntry:       make(map[uint32]serverinterface.IRouter),
		WorkerPoolSize: utils.GlobalConfig.WorkerPoolSize,
		TaskQueue:      make([]chan serverinterface.IRequest, utils.GlobalConfig.WorkerPoolSize),
	}
}

func (m *MessageHandle) DoMsgHandler(request serverinterface.IRequest) {
	msgId := request.GetMsgId()
	if handler, ok := m.MuxEntry[msgId]; !ok {
		log.Println("msgId = ", msgId, "is not found, need register first!")
	} else {
		handler.BeforeHandle(request)
		handler.Handle(request)
		handler.AfterHandle(request)
	}
}

func (m *MessageHandle) AddRouter(msgId uint32, router serverinterface.IRouter) {
	if _, ok := m.MuxEntry[msgId]; ok {
		log.Fatalln("Add router error, msgId = ", msgId, "exist")
	}
	m.MuxEntry[msgId] = router
	log.Println("Add router successfully, msgId = ", msgId)
}

func (m *MessageHandle) StartWorkerPool() {
	for i := 0; i < int(m.WorkerPoolSize); i++ {
		m.TaskQueue[i] = make(chan serverinterface.IRequest, utils.GlobalConfig.MaxWorkerTask)
		go m.StartOneWorker(i, m.TaskQueue[i])
	}
}

func (m *MessageHandle) StartOneWorker(workerId int, taskQueue chan serverinterface.IRequest) {
	log.Println("WorkId is ", workerId, "is started")
	for {
		select {
		case request := <-taskQueue:
			m.DoMsgHandler(request)
		}
	}
}

func (m *MessageHandle) AddMessageToQueue(request serverinterface.IRequest) {
	workerID := request.GetConnection().GetConnID() % m.WorkerPoolSize
	log.Println("Add ConnID = ", request.GetConnection().GetConnID(),
		"workerID = ", workerID)

	m.TaskQueue[workerID] <- request
}
