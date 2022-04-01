package k8s

import (
	"log"
	"os"
	"path/filepath"

	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
)

var clientset *kubernetes.Clientset
var config *rest.Config

func GetClient() *kubernetes.Clientset {
	return clientset
}

func GetConfig() *rest.Config {
	return config
}

func Init() {
	str, _ := os.Getwd()
	kubeconfig := filepath.Join(str, "config", "cluster.yaml")
	log.Println(kubeconfig)
	var err error
	config, err = clientcmd.BuildConfigFromFlags("", kubeconfig)
	if err != nil {
		log.Fatal(err)
	}
	clientset, err = kubernetes.NewForConfig(config)
	if err != nil {
		log.Fatal(err)
	}

}
