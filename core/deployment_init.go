package core

import (
	"fmt"
	v1 "k8s.io/api/apps/v1"
	"k8s.io/apimachinery/pkg/util/wait"
	"k8s.io/client-go/informers"
	"k8sapi/lib"
	"sync"
)

type DeploymentMap struct {
	data sync.Map // 多线程操作，不使用原生的Map，使用sync.Map [key string][]*v1.Deployment, key = namespace
}

func (d *DeploymentMap) Add(dep *v1.Deployment) {
	if list, ok := d.data.Load(dep.Namespace); ok {
		list = append(list.([]*v1.Deployment), dep)
		d.data.Store(dep.Namespace, list)
	} else {
		d.data.Store(dep.Namespace, []*v1.Deployment{dep})
	}
}

func (d *DeploymentMap) ListByNS(ns string) ([]*v1.Deployment, error) {
	if list, ok := d.data.Load(ns); ok {
		return list.([]*v1.Deployment), nil
	}
	return nil, fmt.Errorf("record not found")
}

// DepMap 全局对象
var DepMap *DeploymentMap

func init() {
	DepMap = &DeploymentMap{}
}

type DepHandler struct{}

func (d *DepHandler) OnAdd(obj interface{}) {
	DepMap.Add(obj.(*v1.Deployment))
}

func (d *DepHandler) OnUpdate(oldObj, newObj interface{}) {
	if dep, ok := newObj.(*v1.Deployment); ok {
		fmt.Println(dep.Name)
	}
}
func (d *DepHandler) OnDelete(obj interface{}) {}

func InitDeployment() {
	fact := informers.NewSharedInformerFactory(lib.K8sClient, 0)

	depInformer := fact.Apps().V1().Deployments()
	depInformer.Informer().AddEventHandler(&DepHandler{})

	fact.Start(wait.NeverStop)

	//c,_ := context.WithTimeout(context.Background(), time.Second * 3)
}
