package k8s

import (
	"context"

	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func ListNs() ([]v1.Namespace, error) {
	api := clientset.CoreV1()
	list, err := api.Namespaces().List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		return nil, err
	}
	return list.Items, nil
}

func CreateNs(name string) error {
	ns := &v1.Namespace{
		ObjectMeta: metav1.ObjectMeta{
			Name: name,
		},
	}
	api := clientset.CoreV1()
	_, err := api.Namespaces().Create(context.TODO(), ns, metav1.CreateOptions{})
	return err
}
