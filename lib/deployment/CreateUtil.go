package deployment

import (
	"context"
	"fmt"
	"io/ioutil"
	"k8s.io/api/apps/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/json"
	"k8s.io/apimachinery/pkg/util/yaml"
	"k8sapi/lib"
	"log"
)

func CreateDeployment() {
	rdsDep := &v1.Deployment{}
	b, _ := ioutil.ReadFile("yamls/redis.yaml")
	rdsJson, _ := yaml.ToJSON(b)
	_ = json.Unmarshal(rdsJson,rdsDep)

	ctx := context.Background()
	createOpt := metav1.CreateOptions{}
	_, err := lib.K8sClient.AppsV1().Deployments("devops").
		Create(ctx,rdsDep,createOpt)

	if err != nil {
		log.Fatal("[error] find a err in Create redis Deployment")
	}
	fmt.Println("[info] create redis seccess")
}

