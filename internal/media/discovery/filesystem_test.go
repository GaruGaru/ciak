package discovery

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestFileSystemDiscovery(t *testing.T) {
	discover := NewFileSystemDiscovery("testdata/")

	medias, err := discover.Discover()
	require.NoError(t, err)
	require.Len(t, medias, 3)

	for _, media := range medias {
		require.Contains(t, media.FilePath, "testdata/")
	}

	require.Equal(t, medias[0], Media{
		Name:      "movie0",
		Extension: "avi",
		FilePath:  "testdata/movie0.avi",
		Size:      0,
	})

	require.Equal(t, medias[1], Media{
		Name:      "movie1",
		Extension: "mkv",
		FilePath:  "testdata/movie1.mkv",
		Size:      0,
	})
}
