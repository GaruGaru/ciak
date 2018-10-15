package worker

type Task interface {
	Run() error
}
