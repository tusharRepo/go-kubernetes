package k8sapicall

import (
	"log"
	"webservice/k8sconnect"

	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

//CreateNamespace create a namespace in the cluster
func CreateNamespace(namespaceName string) (*v1.Namespace, error) {
	clientSet, _ := k8sconnect.GetClientSet()

	namespace := v1.Namespace{
		ObjectMeta: metav1.ObjectMeta{
			Name: namespaceName,
		},
	}
	resultNamespace, err := clientSet.CoreV1().Namespaces().Create(&namespace)
	if err != nil {
		log.Printf("CreateNamespace: Error occured while creating namespace: %v", err)
		return nil, err
	}
	log.Printf("CreateNamespace: successfully created namespace: %v", resultNamespace)
	return resultNamespace, err
}
