package k8s

import (
	model "cloud/model"
	"context"

	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	intstr "k8s.io/apimachinery/pkg/util/intstr"
)

func ListService(ns string) []v1.Service {
	list, err := GetClient().CoreV1().Services(ns).List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		panic(err)
	}
	return list.Items
}

func CreateService(param model.ServiceParam) {
	var ports []v1.ServicePort
	ports = append(ports, v1.ServicePort{
		Port:       80,
		TargetPort: intstr.IntOrString{IntVal: 80},
		NodePort:   30081,
	})

	service := &v1.Service{
		ObjectMeta: metav1.ObjectMeta{
			Name: param.Name,
		},
		Spec: v1.ServiceSpec{
			Selector: map[string]string{
				"app": "demo",
			},
			SessionAffinity: v1.ServiceAffinityClientIP,
			Type:            v1.ServiceTypeNodePort,
			Ports:           ports,
		},
	}
	_, err := GetClient().CoreV1().Services(param.Namespace).Create(context.TODO(), service, metav1.CreateOptions{})
	if err != nil {
		panic(err)
	}
}

func DeleteService(namespace string, name string) {
	err := GetClient().CoreV1().Services(namespace).Delete(context.TODO(), name, metav1.DeleteOptions{})
	if err != nil {
		panic(err)
	}
}
