package daemon

type Task interface {
	Run() error
}
