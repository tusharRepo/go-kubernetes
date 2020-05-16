package k8sapicall

import (
	"encoding/json"
	"log"
	"webservice/k8sconnect"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type data struct {
	Name      string `json:"name"`
	Namespace string `json:"namespace,omitempty"`
}

//GetPods returns pod list from all namespaces in json format
func GetPods() ([]byte, error) {
	var podData []data

	clientSet, _ := k8sconnect.GetClientSet()
	pods, err := clientSet.CoreV1().Pods("").List(metav1.ListOptions{})
	if err != nil {
		log.Fatalln("GetPods: failed to get pods:", err)
	}
	for _, pod := range pods.Items {
		temp := data{}
		temp.Name = pod.GetName()
		temp.Namespace = pod.GetNamespace()
		podData = append(podData, temp)
	}
	podJSON, err := json.Marshal(podData)
	return podJSON, err
}

//GetNamespace return list of namespaces in json format
func GetNamespace() ([]byte, error) {
	var namespaceData []data

	clientSet, _ := k8sconnect.GetClientSet()
	namespace, err := clientSet.CoreV1().Namespaces().List(metav1.ListOptions{})
	if err != nil {
		log.Fatalln("GetNamespace: failed to get namespace:", err)
	}
	for _, namespace := range namespace.Items {
		temp := data{}
		temp.Name = namespace.GetName()
		namespaceData = append(namespaceData, temp)
	}
	namespaceJSON, err := json.Marshal(namespaceData)
	return namespaceJSON, err
}
