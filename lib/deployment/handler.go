package deployment

import (
	"context"
	"github.com/gin-gonic/gin"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8sapi/lib"
)

func RegHandlers(r *gin.Engine) {
	r.POST("/update/deployment/scale", incrReplicas)
	r.POST("/core/deployments", ListAllDeployments)
}

func ListAllDeployments(c *gin.Context) {
	ns := c.DefaultQuery("namespace", "devops")
	c.JSON(200, gin.H{"message": "Ok", "result": ListAll(ns)})
}

// incrReplicas 只能 +1 或者 -1
func incrReplicas(c *gin.Context) {
	req := struct {
		Namespace  string `json:"ns" binding:"required,min=1"`
		Deployment string `json:"deployment" binding:"required,min=1"`
		Dec        bool   `json:"dec"` // 是否减少一个
	}{}

	if err := c.ShouldBindJSON(&req); err != nil {
		panic(err.Error())
	}
	ctx := context.Background()
	getOpt := v1.GetOptions{}
	scale, err := lib.K8sClient.AppsV1().Deployments(req.Namespace).GetScale(ctx, req.Deployment, getOpt)
	if err != nil {
		panic(err.Error())
	}
	if req.Dec { // req.Dec = true，表示减少
		scale.Spec.Replicas--
	} else {
		scale.Spec.Replicas++
	}
	updateOpt := v1.UpdateOptions{}
	_, err = lib.K8sClient.AppsV1().Deployments(req.Namespace).UpdateScale(ctx, req.Deployment, scale, updateOpt)
	if err != nil {
		panic(err.Error())
	}
	lib.Success("ok", c)
}
