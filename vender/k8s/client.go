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
	configpath := filepath.Join(str, "config", "cluster.yaml")
	log.Println(configpath)
	var err error
	if IsPathExists(configpath) {
		config, err = clientcmd.BuildConfigFromFlags("", configpath)
	} else {
		config, err = clientcmd.BuildConfigFromFlags("", "")
	}
	if err != nil {
		log.Fatal(err)
	}
	clientset, err = kubernetes.NewForConfig(config)
	if err != nil {
		log.Fatal(err)
	}

}

func IsPathExists(path string) bool {
	_, err := os.Lstat(path)
	if err == nil {
		return true
	}
	if os.IsNotExist(err) {
		return false
	}
	// error other than not existing e.g. permission denied
	return false
}
