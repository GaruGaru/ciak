package worker

import (
	"github.com/sirupsen/logrus"
	"sync/atomic"
)

type WorkerPool struct {
	BaseSize       int
	Tasks          chan Task
	SignalsChannel chan interface{}
	WorkersCount   int32
}

func NewWorkerPool(size int, backPressure int) WorkerPool {
	return WorkerPool{
		BaseSize:       size,
		Tasks:          make(chan Task, backPressure),
		SignalsChannel: make(chan interface{}),
		WorkersCount:   0,
	}
}

func (wp WorkerPool) Start() {
	for i := 0; i < wp.BaseSize; i++ {
		go wp.startWorker()
	}

	<-wp.SignalsChannel

	logrus.Warn("Received stop signal")
}

func (wp WorkerPool) Stop() {
	wp.SignalsChannel <- true
}

func (wp WorkerPool) ScheduleWithWorker(arg Task) bool {
	go wp.startWorker()
	return wp.Schedule(arg)
}

func (wp WorkerPool) Schedule(arg Task) bool {
	select {
	case wp.Tasks <- arg:
		return true
	default:
		return false
	}
}

func (wp *WorkerPool) startWorker() {
	id := atomic.AddInt32(&wp.WorkersCount, 1)
	for {
		select {
		case <-wp.SignalsChannel:
			close(wp.Tasks)
			close(wp.SignalsChannel)
			return
		case task := <-wp.Tasks:
			err := task.Run()
			if err != nil {
				logrus.Error("Worker ", id, " error executing task: ", err)
			}

		}
	}
}
