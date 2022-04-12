package deployment

// Deployment 定义全局结构体
type Deployment struct {
	Name   	     string         // Deployment 名称
	Replicas	 [3]int32       // Deployment 中 副本个数/成功的个数/失败的个数
	Image 		 string         // 使用的镜像名称
	NameSpace 	 string         // 命名空间
	CreateTime   string         // 创建时间
	Pods         []*Pod         // Pod
}

type Pod struct {
	Name string
}