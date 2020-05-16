package k8sconnect

import (
	"log"
	"os"
	"path/filepath"

	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
)

//GetClientSet returns kubernetes client pointer
func GetClientSet() (*kubernetes.Clientset, error) {
	log.Println("getClientSet: Trying to get in-cluster config")
	config, err := rest.InClusterConfig()
	if err != nil && err == rest.ErrNotInCluster {
		kubeConfigFile := filepath.Join(os.Getenv("HOME"), ".kube", "config")
		log.Println("getClientSet: Not in a kubernetes cluster, trying to load kubeconfig from", kubeConfigFile)
		config, err = clientcmd.BuildConfigFromFlags("", kubeConfigFile)
		if err != nil {
			log.Println("getClientSet: The kubeconfig ", kubeConfigFile, " can not be loaded: ", err)
			return nil, err
		}
	} else if err != nil {
		log.Println("getClientSet: error occured in executing rest.InClusterConfig(): ", err)
		return nil, err
	}
	clientSet, err := kubernetes.NewForConfig(config)
	if err != nil {
		log.Println("getClientSet: Could not create a clientset: ", err)
		return nil, err
	}
	log.Println("getClientSet: kubeconfig loaded")
	return clientSet, nil

}
