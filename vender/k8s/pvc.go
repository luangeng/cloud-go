package k8s

import (
	"context"
	"log"

	model "cloud/model"

	v1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/resource"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func ListPv(ns string) []v1.PersistentVolumeClaim {
	api := clientset.CoreV1()
	listOptions := metav1.ListOptions{}
	pvcs, err := api.PersistentVolumeClaims(ns).List(context.TODO(), listOptions)
	if err != nil {
		log.Fatal(err)
	}

	return pvcs.Items
}

func CreatePvc(p model.Pvc) {
	className := "gluster2"
	pvc := &v1.PersistentVolumeClaim{
		ObjectMeta: metav1.ObjectMeta{
			Name: p.Name,
		},
		Spec: v1.PersistentVolumeClaimSpec{
			AccessModes: []v1.PersistentVolumeAccessMode{
				v1.ReadWriteMany,
			},
			StorageClassName: &className,
			Resources: v1.ResourceRequirements{
				Requests: v1.ResourceList{
					v1.ResourceName(v1.ResourceStorage): resource.MustParse("1Gi"),
				},
			},
		},
	}

	api := clientset.CoreV1()
	_, err := api.PersistentVolumeClaims(p.Namespace).Create(context.TODO(), pvc, metav1.CreateOptions{})
	if err != nil {
		panic(err)
	}
}

func DeletePvc(p model.Pvc) {
	api := clientset.CoreV1()
	err := api.PersistentVolumeClaims(p.Namespace).Delete(context.TODO(), p.Name, metav1.DeleteOptions{})
	if err != nil {
		panic(err)
	}
}
