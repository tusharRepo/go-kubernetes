package k8sapicall

import (
	"log"
	"webservice/k8sconnect"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

//DeleteNamespace function removes namesace from cluster
func DeleteNamespace(namespaceName string) error {
	clientSet, _ := k8sconnect.GetClientSet()

	err := clientSet.CoreV1().Namespaces().Delete(namespaceName, &metav1.DeleteOptions{})

	if err != nil {
		log.Printf("DeleteNamespace: error occured while deleting namespace %s: %v", namespaceName, err)
		return err
	}
	return nil
}
