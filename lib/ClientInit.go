package lib

import (
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"log"
)

var K8sClient *kubernetes.Clientset

func init() {
	config := &rest.Config{
		Host:        "http://114.96.64.224:11111",
		BearerToken: "eyJhbGciOiJSUzI1NiIsImtpZCI6IkRQVVk4R3JTSUJBMWlSREVaM2stQVVKcnZsU1RRN3ZLdXF4LW8yVXpXOTQifQ.eyJpc3MiOiJrdWJlcm5ldGVzL3NlcnZpY2VhY2NvdW50Iiwia3ViZXJuZXRlcy5pby9zZXJ2aWNlYWNjb3VudC9uYW1lc3BhY2UiOiJrdWJlLXN5c3RlbSIsImt1YmVybmV0ZXMuaW8vc2VydmljZWFjY291bnQvc2VjcmV0Lm5hbWUiOiJrOHMtY3lrLWFkbWluLXRva2VuLXpiaDVrIiwia3ViZXJuZXRlcy5pby9zZXJ2aWNlYWNjb3VudC9zZXJ2aWNlLWFjY291bnQubmFtZSI6Ims4cy1jeWstYWRtaW4iLCJrdWJlcm5ldGVzLmlvL3NlcnZpY2VhY2NvdW50L3NlcnZpY2UtYWNjb3VudC51aWQiOiJlMjViODg3Ni1mNDA5LTQ5MDYtOTZhMC01MGJjNjE0OGM3OTIiLCJzdWIiOiJzeXN0ZW06c2VydmljZWFjY291bnQ6a3ViZS1zeXN0ZW06azhzLWN5ay1hZG1pbiJ9.aQfOpjBVaPnWf8It28y6jRGpoXooCnaKipsPLpZEzOA4i5ZBAvJKrWzuN1AMRzA98QHBkCP-mra4F-pvbVUjslAUuoEa7ZF_NHT86QeUxOsgNLC1mbIRng2sgrPrtP8Bgo-WLOSkqZqalU0c8MsSW1hIAGoBy9wlOSKJs1iFWzrfKkbfOAg7WOjfnHlEkeRpCvmdntnwc1HY42Mdt5mJKZV9-iu0EMusBMkLqOmcum7ScmQeBYdiBl_hAHAMB1crPxwGj0XhFaJggnQlAGeRuTomUI5hP20FdX6gYGEsswdwaQxAdTA43yWTV6ieQuO-_B-nlZVU9hn54bcdiwTC5g",
	}
	c, err := kubernetes.NewForConfig(config)
	if err != nil {
		log.Fatal(err)
	}
	K8sClient = c
}
