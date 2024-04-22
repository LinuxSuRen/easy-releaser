package image

import (
	"github.com/google/go-containerregistry/pkg/name"
	"github.com/google/go-containerregistry/pkg/v1/remote"
)

type defaultClient struct {
}

func NewDefaultClient() Client {
	return &defaultClient{}
}

// Auth returns an anonymous token
func (c *defaultClient) Auth() (string, error) {
	return "", nil
}

// ListTags lists the tags of an image
func (c *defaultClient) ListTags(image string) (imageTags []ImageTag, err error) {
	var ref name.Reference
	if ref, err = name.ParseReference(image, name.Insecure); err != nil {
		return
	}

	var tags []string
	if tags, err = remote.List(ref.Context()); err != nil {
		return
	}

	for _, tag := range tags {
		imageTags = append(imageTags, ImageTag{
			Name: tag,
		})
	}
	return
}
