package deployment

import (
	"k8sapi/core"
	"k8sapi/lib"
)

func ListAll(namespace string) (ret []*Deployment) {
	depList, err := core.DepMap.ListByNS(namespace)
	lib.CheckErr(err)
	for _, item := range depList {
		ret = append(ret, &Deployment{
			Name:     item.Name,
			Replicas: [3]int32{item.Status.Replicas, item.Status.AvailableReplicas, item.Status.UnavailableReplicas},
			Image:    GetImagesByDeployment(*item),
		})
	}
	return
}
