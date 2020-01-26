package reconcile

import (
	"context"
	"fmt"
	"net/http"
	"reflect"

	"github.com/ericchiang/k8s"
	"github.com/thcyron/skop/skop"
)

// EnsureResource is a function, which ensures a Resource in the cluster. This can be used to implement different functions
// to ensure other resources, like Deployments.
func EnsureResource(ctx context.Context, client skop.Client, serviceSpec k8s.Resource, customizationCallback func(existing k8s.Resource) error) error {
	clusterResourceTest := reflect.New(reflect.TypeOf(serviceSpec)).Elem()
	clusterResourceInterface := clusterResourceTest.Interface()
	clusterResource, _ := clusterResourceInterface.(k8s.Resource)

	fmt.Printf("typeOf clusterResourceInterface %v\n", reflect.TypeOf(clusterResourceInterface))
	fmt.Printf("typeOf clusterResource %v\n", reflect.TypeOf(clusterResource))
	fmt.Printf("typeOf serviceSpec %v\n", reflect.TypeOf(serviceSpec))
	fmt.Printf("serviceSpec %v\n", serviceSpec)

	namespace := *serviceSpec.GetMetadata().Namespace
	name := *serviceSpec.GetMetadata().Name

	fmt.Printf("namespace: %v, name: %v \n", namespace, name)

	err := client.Get(
		ctx,
		namespace,
		name,
		clusterResource,
	)

	fmt.Printf("Getting existing resource err: %v, clusterResource: %v \n", err, clusterResource)

	if err != nil {
		if apiErr, ok := err.(*k8s.APIError); ok {
			if apiErr.Code == http.StatusNotFound {
				return client.Create(ctx, serviceSpec)
			}
		}
		return err
	}

	fmt.Printf("clusterResource: %v \n", clusterResource)
	if clusterResource == nil {
		fmt.Printf("Creating clusterResource with serviceSpec: %v \n", clusterResource)
		return client.Create(ctx, serviceSpec)
	}

	if err := customizationCallback(clusterResource); err != nil {
		fmt.Printf("customizationCallback error: %v\n", err)
		return err
	}
	return client.Update(ctx, clusterResource)
}
