package deployment

import (
	"context"
	v1 "k8s.io/api/apps/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8sapi/lib"
	"log"
)

//
func GetPodsByDep(ns string, dep *v1.Deployment) []*Pod{
	ctx := context.Background()
	ListOpt := metav1.ListOptions{
		LabelSelector: GetLabels(dep.Spec.Selector.MatchLabels),
	}
	list, err := lib.K8sClient.CoreV1().Pods(ns).List(ctx, ListOpt)
	if err != nil {
		panic(err.Error())
	}
	pods := make([]*Pod,len(list.Items))
	for i, item := range list.Items {
		pods[i] = &Pod{
			Name: item.Name,
		}
	}
	return pods
}


// GetDeployment 使用 .AppsV1 获取命名空间详情
func GetDeployment(ns string, name string) *Deployment {
	ctx := context.Background()
	getOpt := metav1.GetOptions{}
	dep, err := lib.K8sClient.AppsV1().Deployments(ns).Get(ctx,name,getOpt)
	if err != nil {
		log.Fatal(err)
	}
	return &Deployment{
		Name: dep.Name,
		NameSpace: dep.Namespace,
		Image: GetImages(*dep),
		CreateTime: dep.CreationTimestamp.Format("2006-01-02 15:04:05"),
		Pods: GetPodsByDep(ns,dep),
	}
}
