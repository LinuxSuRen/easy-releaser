package v1alpha1_test

import (
	"testing"

	"github.com/kubesphere-sigs/ks-releaser/api/v1alpha1"
	"github.com/stretchr/testify/assert"
)

func TestArtifacts(t *testing.T) {
	defaultArtifact := v1alpha1.Artifact{
		Image: v1alpha1.ArtifactImage{
			Image: "test",
		},
	}

	t.Run("cannot add duplicated items", func(t *testing.T) {
		artifacts := v1alpha1.Artifacts{}

		artifacts = artifacts.AddUniq(defaultArtifact)

		assert.Equal(t, 1, artifacts.Len())
		assert.True(t, artifacts.HasItem(defaultArtifact))
	})
}
