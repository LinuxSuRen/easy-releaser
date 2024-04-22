package v1alpha1

func (a Artifacts) AddUniq(artifact Artifact) Artifacts {
	if !a.HasItem(artifact) {
		a = append(a, artifact)
	}
	return a
}

func (a Artifacts) HasItem(artifact Artifact) (exist bool) {
	for _, item := range a {
		if item.Image.Image == artifact.Image.Image && item.Image.Tag == artifact.Image.Tag {
			exist = true
		}
	}
	return
}

func (a Artifacts) Len() int {
	return len(a)
}
