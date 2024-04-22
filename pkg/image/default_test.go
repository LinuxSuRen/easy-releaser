package image_test

import (
	"fmt"
	"testing"

	"github.com/kubesphere-sigs/ks-releaser/pkg/image"
	"github.com/linuxsuren/api-testing/pkg/mock"
	"github.com/stretchr/testify/assert"
)

func TestListTags(t *testing.T) {
	server := mock.NewInMemoryServer(0)

	err := server.Start(mock.NewLocalFileReader("testdata/api.yaml"), "/")
	assert.NoError(t, err)
	defer func() {
		server.Stop()
	}()

	tags, err := image.NewDefaultClient().ListTags(
		fmt.Sprintf("localhost:%s/linuxsuren/api-testing", server.GetPort()))
	assert.NoError(t, err)
	assert.Equal(t, 3, len(tags))

	t.Run("invalid image reference", func(t *testing.T) {
		_, err := image.NewDefaultClient().ListTags("http://fake/com:*^&")
		assert.Error(t, err, err)
	})

	t.Run("get error when requesting image", func(t *testing.T) {
		_, err := image.NewDefaultClient().ListTags(
			fmt.Sprintf("localhost:%s/linuxsuren/error", server.GetPort()))
		assert.Error(t, err, err)
	})
}
