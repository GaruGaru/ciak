package encoding

import (
	"github.com/GaruGaru/ciak/internal/media/models"
)

type MediaEncoder interface {
	Encode(input string, output string) error
	CanEncode(models.MediaFormat) bool
}
