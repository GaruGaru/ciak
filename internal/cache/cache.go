package cache

type Cache interface {
	Get(interface{}) (interface{}, bool, error)
	Set(interface{}, interface{}) error
}
