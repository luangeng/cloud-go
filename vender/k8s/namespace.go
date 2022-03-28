package k8s

import (
	"context"

	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func ListNs() []v1.Namespace {
	api := clientset.CoreV1()
	list, err := api.Namespaces().List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		panic(err)
	}
	return list.Items
}

func CreateNs(name string) {
	ns := &v1.Namespace{
		ObjectMeta: metav1.ObjectMeta{
			Name: name,
		},
	}
	api := clientset.CoreV1()
	_, err := api.Namespaces().Create(context.TODO(), ns, metav1.CreateOptions{})
	if err != nil {
		panic(err)
	}
}
