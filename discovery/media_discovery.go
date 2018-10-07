package discovery

import (
	"fmt"
	"hash/fnv"
)

type Media struct {
	Name     string
	Extension string
	FilePath string
	Size     int64
}

func (m Media) Hash() string {
	h := fnv.New32a()
	h.Write([]byte(fmt.Sprintf("%s%s", m.FilePath, m.Name)))
	return fmt.Sprint(h.Sum32())
}

type MediaDiscovery interface {
	Discover() ([]Media, error)
	Resolve(hash string) (Media, error)
}
