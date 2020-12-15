package details

import (
	"errors"
	"github.com/GaruGaru/ciak/pkg/cache"
	"github.com/GaruGaru/ciak/pkg/media/models"
	"github.com/sirupsen/logrus"
	"sync"
)

var (
	ErrDetailsNotFound = errors.New("media details not found")
)

type Request struct {
	Title string
}

type Retriever interface {
	Details(Request) (models.Details, error)
}

type Controller struct {
	Retrievers []Retriever
	Cache      cache.Cache
}

func NewController(cache cache.Cache, retrievers ...Retriever) *Controller {
	return &Controller{
		Retrievers: retrievers,
		Cache:      cache,
	}
}

func (c *Controller) Details(request Request) (models.Details, error) {
	cached, present, err := c.Cache.Get(request)
	if err != nil {
		logrus.Warnf("error reading from cache: %s", err)
	}

	if present {
		return cached.(models.Details), nil
	}

	for _, retriever := range c.Retrievers {
		details, err := retriever.Details(request)
		if err != nil {
			// todo we may want to log this
			continue
		}

		if err := c.Cache.Set(request, details); err != nil {
			logrus.Warnf("error writing from cache: %s", err)
		}
		return details, nil
	}

	return models.Details{}, ErrDetailsNotFound
}

func (c *Controller) DetailsByTitleBulk(requests ...Request) (map[string]models.Details, error) {
	var wg sync.WaitGroup
	wg.Add(len(requests))

	results := make(chan models.Details, len(requests))
	for _, request := range requests {
		go func(request Request) {
			defer wg.Done()
			details, err := c.Details(request)
			if err != nil {
				logrus.Debugf("unable to get title metadata for %s: %s", request.Title, err.Error())
			}
			results <- details
		}(request)
	}

	wg.Wait()

	close(results)

	out := make(map[string]models.Details)

	for res := range results {
		out[res.Name] = res
	}

	return out, nil
}
