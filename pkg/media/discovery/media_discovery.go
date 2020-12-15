package discovery

import (
	"github.com/GaruGaru/ciak/pkg/media/models"
)

type MediaDiscovery interface {
	Discover() ([]models.Media, error)
	Resolve(hash string) (models.Media, error)
}
