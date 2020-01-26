package reconcile

import (
	"context"
	"github.com/ericchiang/k8s"
	corev1 "github.com/ericchiang/k8s/apis/core/v1"

	"github.com/thcyron/skop/skop"
)

func Service(ctx context.Context, client skop.Client, serviceDef *corev1.Service) error {
	return EnsureResource(ctx, client, serviceDef, func(existingResource k8s.Resource) error {
		existingService, _ := existingResource.(*corev1.Service)


		clusterIP := existingService.Spec.ClusterIP
		existingService.Metadata.Labels = serviceDef.Metadata.Labels
		existingService.Metadata.Annotations = serviceDef.Metadata.Annotations
		existingService.Spec = serviceDef.Spec
		existingService.Spec.ClusterIP = clusterIP

		return nil
	})
}
