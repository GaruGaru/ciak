package discovery

import (
	"fmt"
	"github.com/GaruGaru/ciak/internal/media/models"
	"hash/fnv"
)

type MediaDiscovery interface {
	Discover() ([]Media, error)
	Resolve(hash string) (Media, error)
}

type Media struct {
	Name     string
	Format   models.MediaFormat
	FilePath string
	Size     int64
}

func (m Media) Hash() string {
	h := fnv.New32a()
	h.Write([]byte(fmt.Sprintf("%s%s", m.FilePath, m.Name)))
	return fmt.Sprint(h.Sum32())
}
