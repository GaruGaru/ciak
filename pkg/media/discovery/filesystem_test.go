package discovery

import (
	"github.com/GaruGaru/ciak/pkg/media/models"
	"github.com/stretchr/testify/require"
	"path"
	"testing"
)

func TestFileSystemDiscovery(t *testing.T) {
	discover := NewFileSystemDiscovery("testdata")

	medias, err := discover.Discover()
	require.NoError(t, err)
	require.Len(t, medias, 3)

	for _, media := range medias {
		require.Contains(t, media.FilePath, "testdata")
	}

	require.Equal(t, medias[0], models.Media{
		Name:     "movie0",
		Format:   models.MediaFormatAvi,
		FilePath: path.Join("testdata", "movie0.avi"),
		Size:     0,
	})

	require.Equal(t, medias[1], models.Media{
		Name:     "movie1",
		Format:   models.MediaFormatMkv,
		FilePath: path.Join("testdata", "movie1.mkv"),
		Size:     0,
	})
}
