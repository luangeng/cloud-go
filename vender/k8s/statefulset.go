package k8s

import (
	"context"
	"fmt"

	appsv1 "k8s.io/api/apps/v1"
	v1 "k8s.io/api/apps/v1"
	apiv1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/resource"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func ListStateful(ns string) []v1.StatefulSet {
	list, err := GetClient().AppsV1().StatefulSets(ns).List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		panic(err)
	}
	return list.Items
}

func CreateStateful() {
	storageClassName := "gluster"
	var envs []apiv1.EnvVar
	envs = append(envs, apiv1.EnvVar{Name: "test", Value: "123"})

	var ports []apiv1.ContainerPort
	ports = append(ports, apiv1.ContainerPort{
		Name:          "http",
		Protocol:      apiv1.ProtocolTCP,
		ContainerPort: 80,
	})

	var volumeMounts []apiv1.VolumeMount
	volumeMounts = append(volumeMounts, apiv1.VolumeMount{
		Name:      "pv",
		MountPath: "/mypv",
	})

	client := GetClient().AppsV1().StatefulSets(apiv1.NamespaceDefault)

	statefulset := &appsv1.StatefulSet{
		ObjectMeta: metav1.ObjectMeta{
			Name: "demo",
		},
		Spec: appsv1.StatefulSetSpec{
			ServiceName: "stateful",
			Replicas:    int32Ptr(1),
			Selector: &metav1.LabelSelector{
				MatchLabels: map[string]string{
					"app": "demo",
				},
			},
			Template: apiv1.PodTemplateSpec{
				ObjectMeta: metav1.ObjectMeta{
					Labels: map[string]string{
						"app": "demo",
					},
				},
				Spec: apiv1.PodSpec{
					Containers: []apiv1.Container{
						{
							Name:            "web",
							Image:           "luangeng/tool",
							ImagePullPolicy: "IfNotPresent",
							Ports:           ports,
							Env:             envs,
							VolumeMounts:    volumeMounts,
						},
					},
				},
			},

			VolumeClaimTemplates: []apiv1.PersistentVolumeClaim{
				{
					ObjectMeta: metav1.ObjectMeta{
						Name: "pv",
					},
					Spec: apiv1.PersistentVolumeClaimSpec{
						StorageClassName: &storageClassName,
						AccessModes: []apiv1.PersistentVolumeAccessMode{
							apiv1.ReadWriteOnce,
						},
						Resources: apiv1.ResourceRequirements{
							Requests: apiv1.ResourceList{
								apiv1.ResourceName(apiv1.ResourceStorage): resource.MustParse("1Gi"),
							},
						},
					},
				},
			},
		},
	}

	result, err := client.Create(context.TODO(), statefulset, metav1.CreateOptions{})
	if err != nil {
		panic(err)
	}
	fmt.Printf("Created statefulset %q.\n", result.GetObjectMeta().GetName())
}

func DeleteStateful(ns string, name string) {
	client := GetClient().AppsV1().StatefulSets(ns)
	deletePolicy := metav1.DeletePropagationForeground
	if err := client.Delete(context.TODO(), name, metav1.DeleteOptions{
		PropagationPolicy: &deletePolicy,
	}); err != nil {
		panic(err)
	}
}
