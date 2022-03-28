package k8s

import (
	"context"
	"log"

	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func ListPv(ns string) []v1.PersistentVolumeClaim {
	api := clientset.CoreV1()
	// setup list options
	listOptions := metav1.ListOptions{}
	pvcs, err := api.PersistentVolumeClaims(ns).List(context.TODO(), listOptions)
	if err != nil {
		log.Fatal(err)
	}

	return pvcs.Items
}
