package daemon

import "github.com/sirupsen/logrus"

type WorkerPool struct {
	Size           int
	Tasks          chan Task
	SignalsChannel chan interface{}
}

func NewWorkerPool(size int, backPressure int) WorkerPool {
	return WorkerPool{
		Size:           size,
		Tasks:          make(chan Task, backPressure),
		SignalsChannel: make(chan interface{}),
	}
}

func (wp WorkerPool) Start() {
	for i := 0; i < wp.Size; i++ {
		go wp.startWorker(i)
	}
}

func (wp WorkerPool) Stop(arg Task) {
	wp.SignalsChannel <- true
}

func (wp WorkerPool) Schedule(arg Task) bool {
	select {
	case wp.Tasks <- arg:
		return true
	default:
		return false
	}
}

func (wp WorkerPool) startWorker(id int) {
	for {
		select {
		case <-wp.SignalsChannel:
			return
		case task := <-wp.Tasks:
			err := task.Run()
			if err != nil {
				logrus.Error("Worker ", id, " error executing task: ", err)
			}

		}
	}
}
