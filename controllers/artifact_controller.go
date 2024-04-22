package controllers

import (
	"context"
	"errors"
	"time"

	"github.com/go-logr/logr"
	devopsv1alpha1 "github.com/kubesphere-sigs/ks-releaser/api/v1alpha1"
	"github.com/kubesphere-sigs/ks-releaser/pkg/image"
	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/log"
)

// ArtifactController is responsible for reporting errors to git provider
type ArtifactController struct {
	logger logr.Logger
	client.Client
	Scheme *runtime.Scheme
}

func (c *ArtifactController) Reconcile(ctx context.Context, req ctrl.Request) (result ctrl.Result, err error) {
	c.logger = log.FromContext(ctx)

	releaser := &devopsv1alpha1.Releaser{}
	if err = c.Get(ctx, req.NamespacedName, releaser); err != nil {
		err = client.IgnoreNotFound(err)
		return
	}

	// start to fetch image tags only when the phase is done
	if releaser.Spec.Phase != devopsv1alpha1.PhaseDone {
		result = ctrl.Result{
			Requeue:      true,
			RequeueAfter: time.Minute,
		}
		return
	}

	// asume the workflow is successed
	// TODO it can only fetch the public access of image
	for _, artifact := range releaser.Spec.Artifacts {
		if releaser.Status.Artifacts.HasItem(artifact) {
			continue
		}

		imageRef := artifact.Image.Image
		c.logger.Info("start to watch image", "ref", imageRef)

		if tags, tErr := image.NewDefaultClient().ListTags(imageRef); tErr == nil {
			for _, tag := range tags {
				if tag.Name == artifact.Image.Tag {
					releaser.Status.Artifacts = append(releaser.Status.Artifacts, artifact)
					c.Status().Update(ctx, releaser)
					break
				}
			}
		} else {
			err = errors.Join(err)
		}
	}

	result = ctrl.Result{
		Requeue:      err != nil,
		RequeueAfter: time.Minute,
	}
	return
}

// SetupWithManager sets up the controller with the Manager.
func (c *ArtifactController) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&devopsv1alpha1.Releaser{}).
		Complete(c)
}
