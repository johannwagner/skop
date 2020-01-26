package reconcile

import (
	"context"

	"github.com/ericchiang/k8s"
	appsv1 "github.com/ericchiang/k8s/apis/apps/v1"

	"github.com/thcyron/skop/skop"
)

func Deployment(ctx context.Context, client skop.Client, expected *appsv1.Deployment) error {
	return EnsureResource(ctx, client, expected, func(existingResource k8s.Resource) error {
		existingDeployment, _ := existingResource.(*appsv1.Deployment)

		existingDeployment.Metadata.Labels = expected.Metadata.Labels
		existingDeployment.Metadata.Annotations = expected.Metadata.Annotations
		existingDeployment.Spec = expected.Spec

		return nil
	})
}
