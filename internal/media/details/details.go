package details

import (
	"github.com/GaruGaru/ciak/internal/media/models"
)

type Request struct {
	Name string
}

type Retriever interface {
	Details(Request) (models.Details, error)
}
