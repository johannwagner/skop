package reconcile

import (
	"context"
	"errors"
	"fmt"

	"github.com/ericchiang/k8s"
	corev1 "github.com/ericchiang/k8s/apis/core/v1"

	"github.com/thcyron/skop/skop"
)

func ConfigMap(ctx context.Context, client skop.Client, configMap *corev1.ConfigMap) error {
	return EnsureResource(ctx, client, configMap, func(existingResource k8s.Resource) error {

		fmt.Printf("EnsureResource called, existingResource: %v\n", existingResource)

		existingConfigMap, isOk := existingResource.(*corev1.ConfigMap)

		if !isOk {
			fmt.Printf("Invalid cast from existingResource\n")
			return errors.New("Invalid cast from existing resource")
		}

		existingConfigMap.Metadata.Labels = configMap.Metadata.Labels
		existingConfigMap.Metadata.Annotations = configMap.Metadata.Annotations
		existingConfigMap.Data = configMap.Data
		existingConfigMap.BinaryData = configMap.BinaryData

		return nil
	})
}
