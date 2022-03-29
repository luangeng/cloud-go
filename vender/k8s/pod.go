package k8s

import (
	"context"
	"fmt"

	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func ListPod() []v1.Pod {
	pods, err := GetClient().CoreV1().Pods("default").List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		panic(err.Error())
	}
	fmt.Printf("There are %d pods in the cluster\n", len(pods.Items))
	return pods.Items

	// Examples for error handling:
	// - Use helper functions like e.g. errors.IsNotFound()
	// - And/or cast to StatusError and use its properties like e.g. ErrStatus.Message
	// namespace := "default"
	// pod := "me"
	// _, err = GetClient().CoreV1().Pods(namespace).Get(context.TODO(), pod, metav1.GetOptions{})
	// if errors.IsNotFound(err) {
	// 	fmt.Printf("Pod %s in namespace %s not found\n", pod, namespace)
	// } else if statusError, isStatus := err.(*errors.StatusError); isStatus {
	// 	fmt.Printf("Error getting pod %s in namespace %s: %v\n",
	// 		pod, namespace, statusError.ErrStatus.Message)
	// } else if err != nil {
	// 	panic(err.Error())
	// } else {
	// 	fmt.Printf("Found pod %s in namespace %s\n", pod, namespace)
	// }
}

func CretaePod() {
	pod := &v1.Pod{
		ObjectMeta: metav1.ObjectMeta{
			Name: "demo",
			Labels: map[string]string{
				"app": "demo",
			},
		},
		Spec: v1.PodSpec{
			Containers: []v1.Container{
				{
					Image:           "",
					Name:            "",
					ImagePullPolicy: v1.PullIfNotPresent,
				},
			},
		},
	}
	_, err := GetClient().CoreV1().Pods("default").Create(context.TODO(), pod, metav1.CreateOptions{})
	if err != nil {
		panic(err.Error())
	}
}

func DeletePod(ns string, name string) {
	err := GetClient().CoreV1().Pods(ns).Delete(context.TODO(), name, metav1.DeleteOptions{})
	if err != nil {
		panic(err.Error())
	}
}
