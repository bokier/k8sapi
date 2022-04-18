package main

import (
	"github.com/gin-gonic/gin"
	"k8sapi/core"
	"k8sapi/lib"
	"k8sapi/lib/deployment"
	"net/http"
)

func main() {
	r := gin.New()
	deployment.RegHandlers(r)
	r.Static("/static", "./static")
	r.LoadHTMLGlob("html/**/*")
	r.GET("/deployments", func(c *gin.Context) {
		c.HTML(http.StatusOK, "deployment_list.html",
			lib.DataBuilder().
				SetTitle("deployment列表").
				SetData("DepList", deployment.ListAll("devops")))
	})
	r.GET("/deployments/:name", func(c *gin.Context) {
		c.HTML(http.StatusOK, "deployment_detail.html",
			lib.DataBuilder().
				SetTitle("deployment详细-"+c.Param("name")).
				SetData("DepDetail", deployment.GetDeployment("devops", c.Param("name"))))
	})
	core.InitDeployment() // 初始化 Deployment 列表
	r.Run(":18081")
}
