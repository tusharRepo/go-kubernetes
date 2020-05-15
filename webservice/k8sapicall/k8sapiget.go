package k8sapicall

import (
	"encoding/json"
	"log"
	"webservice/k8sconnect"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type Data struct {
	Name      string `json:"name"`
	Namespace string `json:"namespace"`
}

func GetPods() ([]byte, error) {
	var podData []Data

	clientSet, _ := k8sconnect.GetClientSet()
	pods, err := clientSet.CoreV1().Pods("").List(metav1.ListOptions{})
	if err != nil {
		log.Fatalln("GetPods: failed to get pods:", err)
	}
	for _, pod := range pods.Items {
		temp := Data{}
		temp.Name = pod.GetName()
		temp.Namespace = pod.GetNamespace()
		podData = append(podData, temp)
	}
	podJson, err := json.Marshal(podData)
	return podJson, err
}

func GetNamespace() ([]byte, error) {
	var namespaceData []Data

	clientSet, _ := k8sconnect.GetClientSet()
	namespace, err := clientSet.CoreV1().Namespaces().List(metav1.ListOptions{})
	if err != nil {
		log.Fatalln("GetNamespace: failed to get namespace:", err)
	}
	for _, namespace := range namespace.Items {
		temp := Data{}
		temp.Name = namespace.GetName()
		namespaceData = append(namespaceData, temp)
	}
	namespaceJson, err := json.Marshal(namespaceData)
	return namespaceJson, err
}
