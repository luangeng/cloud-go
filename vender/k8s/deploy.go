package k8s

import (
	"context"
	"fmt"

	appsv1 "k8s.io/api/apps/v1"
	apiv1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/util/retry"
)

func CreateDeploy() {
	var envs []apiv1.EnvVar
	envs = append(envs, apiv1.EnvVar{Name: "test", Value: "123"})

	var ports []apiv1.ContainerPort
	ports = append(ports, apiv1.ContainerPort{
		Name:          "http",
		Protocol:      apiv1.ProtocolTCP,
		ContainerPort: 80,
	})

	var volumes []apiv1.Volume
	volumes = append(volumes, apiv1.Volume{
		Name: "mypv",
		VolumeSource: apiv1.VolumeSource{
			PersistentVolumeClaim: &apiv1.PersistentVolumeClaimVolumeSource{
				ClaimName: "myclaim",
			},
		},
	})

	var volumeMounts []apiv1.VolumeMount
	volumeMounts = append(volumeMounts, apiv1.VolumeMount{
		Name:      "mypv",
		MountPath: "/mypv",
	})

	deploymentsClient := GetClient().AppsV1().Deployments(apiv1.NamespaceDefault)

	deployment := &appsv1.Deployment{
		ObjectMeta: metav1.ObjectMeta{
			Name: "demo-deployment",
		},
		Spec: appsv1.DeploymentSpec{
			Replicas: int32Ptr(2),
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
					Volumes: volumes,
				},
			},
		},
	}

	// Create Deployment
	result, err := deploymentsClient.Create(context.TODO(), deployment, metav1.CreateOptions{})
	if err != nil {
		panic(err)
	}
	fmt.Printf("Created deployment %q.\n", result.GetObjectMeta().GetName())
}

func ListDeploy(ns string) []appsv1.Deployment {
	deploymentsClient := GetClient().AppsV1().Deployments(ns)
	list, err := deploymentsClient.List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		panic(err)
	}
	return list.Items
}

func DeleteDeploy(ns string, name string) {
	deploymentsClient := GetClient().AppsV1().Deployments(ns)
	deletePolicy := metav1.DeletePropagationForeground
	if err := deploymentsClient.Delete(context.TODO(), name, metav1.DeleteOptions{
		PropagationPolicy: &deletePolicy,
	}); err != nil {
		panic(err)
	}
}

func UpdateDeploy() {
	deploymentsClient := GetClient().AppsV1().Deployments(apiv1.NamespaceDefault)
	fmt.Println("Updating deployment...")
	retryErr := retry.RetryOnConflict(retry.DefaultRetry, func() error {
		result, getErr := deploymentsClient.Get(context.TODO(), "demo-deployment", metav1.GetOptions{})
		if getErr != nil {
			panic(fmt.Errorf("Failed to get latest version of Deployment: %v", getErr))
		}

		result.Spec.Replicas = int32Ptr(1)
		_, updateErr := deploymentsClient.Update(context.TODO(), result, metav1.UpdateOptions{})
		return updateErr
	})
	if retryErr != nil {
		panic(fmt.Errorf("Update failed: %v", retryErr))
	}
}

func int32Ptr(i int32) *int32 { return &i }
