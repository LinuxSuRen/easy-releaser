package image

import "time"

type Client interface {
	// Auth returns an anonymous token
	Auth() (string, error)
	// ListTags lists the tags of an image
	ListTags(string) ([]ImageTag, error)
}

type ImageTag struct {
	Name        string    `json:"name"`
	CreatedTime time.Time `json:"createdTime"`
}
