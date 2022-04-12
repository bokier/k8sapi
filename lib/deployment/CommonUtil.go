package deployment

import (
	"fmt"
	v1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	"strings"
)

// GetImagesByDeployment 这里通过Deployment获取镜像
/*
	目前这里是获取第一个镜像
	多个镜像不作处理
*/
func GetImagesByDeployment(dep v1.Deployment) string {
	return GetImagesByPod(dep.Spec.Template.Spec.Containers)
}

// GetImagesByPod 这里通过 Pod 获取镜像
func GetImagesByPod(containers []corev1.Container) string {
	var imageName string
	if ok := strings.Contains(containers[0].Image, ":"); ok {
		imageName = containers[0].Image
	} else {
		imageName = containers[0].Image + ":latest"
	}
	return imageName
}

// GetLabels 将 map 转化成 string， 输出格式:aa=xxx,bb=xxx,cc=xxx
func GetLabels(m map[string]string) string {
	labels := ""
	for k, v := range m {
		if labels != "" {
			labels += ","
		}
		labels += fmt.Sprintf("%s=%s", k, v)
	}
	return labels
}
