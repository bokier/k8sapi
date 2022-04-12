package deployment

import (
	"fmt"
	v1 "k8s.io/api/apps/v1"
	"strings"
)

// GetImages 这里是用来处理获取的镜像的
/*
	目前这里是获取第一个镜像
	多个镜像不作处理
 */
func GetImages(dep v1.Deployment) string {
	var imageName string
	if ok := strings.Contains(dep.Spec.Template.Spec.Containers[0].Image,":"); ok {
		imageName = dep.Spec.Template.Spec.Containers[0].Image
	} else {
		imageName = dep.Spec.Template.Spec.Containers[0].Image + ":latest"
	}
	return imageName
}

// GetLabels 将 map 转化成 string， 输出格式:aa=xxx,bb=xxx,cc=xxx
func GetLabels(m map[string]string)string {
	labels := ""
	for k, v := range m {
		if labels != "" {
			labels += ","
		}
		labels += fmt.Sprintf("%s=%s",k,v)
	}
	return labels
}