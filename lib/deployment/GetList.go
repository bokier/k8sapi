package deployment

import (
	"context"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8sapi/lib"
	"log"
)

func ListAll(namespace string) (ret []*Deployment) {
	ctx := context.Background()
	listOpt := v1.ListOptions{}
	depList, err := lib.K8sClient.AppsV1().Deployments(namespace).List(ctx, listOpt)
	if err != nil {
		log.Fatal(err)
	}
	for _, item := range depList.Items {
		ret = append(ret, &Deployment{
			Name:     item.Name,
			Replicas: [3]int32{item.Status.Replicas, item.Status.AvailableReplicas, item.Status.UnavailableReplicas},
			Image:    GetImagesByDeployment(item),
		})
	}
	return
}
